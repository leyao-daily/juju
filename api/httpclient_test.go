// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"

	"github.com/go-macaroon-bakery/macaroon-bakery/v3/httpbakery"
	"github.com/juju/errors"
	"github.com/juju/names/v6"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	"gopkg.in/httprequest.v1"
	"gopkg.in/macaroon.v2"

	"github.com/juju/juju/api"
	apiservertesting "github.com/juju/juju/apiserver/testing"
	"github.com/juju/juju/core/version"
	"github.com/juju/juju/internal/testing"
	jujutesting "github.com/juju/juju/juju/testing"
	"github.com/juju/juju/rpc/params"
)

type httpSuite struct {
	testing.BaseSuite

	client *httprequest.Client
	conn   api.Connection
}

var _ = gc.Suite(&httpSuite{})

func (s *httpSuite) SetUpTest(c *gc.C) {
	s.BaseSuite.SetUpTest(c)

	srv := apiservertesting.NewAPIServer(func(modelUUID string) (interface{}, error) {
		return &testRootAPI{}, nil
	})
	s.AddCleanup(func(_ *gc.C) { srv.Close() })
	info := &api.Info{
		Addrs:          srv.Addrs,
		CACert:         testing.CACert,
		ControllerUUID: testing.ControllerTag.Id(),
		ModelTag:       testing.ModelTag,
	}
	var err error
	s.conn, err = api.Open(context.Background(), info, api.DialOpts{})
	c.Assert(err, jc.ErrorIsNil)
	s.AddCleanup(func(c *gc.C) { c.Assert(s.conn.Close(), jc.ErrorIsNil) })
	client, err := s.conn.HTTPClient()
	c.Assert(err, gc.IsNil)
	s.client = client
}

var httpClientTests = []struct {
	about           string
	handler         http.HandlerFunc
	expectResponse  interface{}
	expectError     string
	expectErrorIs   errors.ConstError
	expectErrorCode string
	expectErrorInfo map[string]interface{}
}{{
	about: "success",
	handler: func(w http.ResponseWriter, req *http.Request) {
		httprequest.WriteJSON(w, http.StatusOK, "hello, world")
	},
	expectResponse: newString("hello, world"),
}, {
	about: "unauthorized status without discharge-required error",
	handler: func(w http.ResponseWriter, req *http.Request) {
		httprequest.WriteJSON(w, http.StatusUnauthorized, params.Error{
			Message: "something",
		})
	},
	expectError: `Get http://.*/: something`,
}, {
	about:         "non-JSON NotFound error response",
	handler:       http.NotFound,
	expectError:   `(?m)Get http://.*/: 404 page not found.*`,
	expectErrorIs: errors.NotFound,
}, {
	about: "bad error response",
	handler: func(w http.ResponseWriter, req *http.Request) {
		type badResponse struct {
			Message map[string]int
		}
		httprequest.WriteJSON(w, http.StatusUnauthorized, badResponse{
			Message: make(map[string]int),
		})
	},
	expectError: `Get http://.*/: incompatible error response: json: cannot unmarshal object into Go .+`,
}, {
	about: "bad charms error response",
	handler: func(w http.ResponseWriter, req *http.Request) {
		type badResponse struct {
			Error    string         `json:"error"`
			CharmURL map[string]int `json:"charm-url"`
		}
		httprequest.WriteJSON(w, http.StatusUnauthorized, badResponse{
			Error:    "something",
			CharmURL: make(map[string]int),
		})
	},
	expectError: `Get http://.*/: incompatible error response: json: cannot unmarshal object into Go .+`,
}, {
	about: "no message in ErrorResponse",
	handler: func(w http.ResponseWriter, req *http.Request) {
		httprequest.WriteJSON(w, http.StatusUnauthorized, params.ErrorResult{
			Error: &params.Error{},
		})
	},
	expectError: `Get http://.*/: error response with no message`,
}, {
	about: "no message in Error",
	handler: func(w http.ResponseWriter, req *http.Request) {
		httprequest.WriteJSON(w, http.StatusUnauthorized, params.Error{})
	},
	expectError: `Get http://.*/: error response with no message`,
}, {
	about: "charms error response",
	handler: func(w http.ResponseWriter, req *http.Request) {
		httprequest.WriteJSON(w, http.StatusBadRequest, params.CharmsResponse{
			Error:     "some error",
			ErrorCode: params.CodeBadRequest,
			ErrorInfo: params.DischargeRequiredErrorInfo{
				MacaroonPath: "foo",
			}.AsMap(),
		})
	},
	expectError:     `.*some error$`,
	expectErrorCode: params.CodeBadRequest,
	expectErrorInfo: params.DischargeRequiredErrorInfo{
		MacaroonPath: "foo",
	}.AsMap(),
}, {
	about: "discharge-required response with no attached info",
	handler: func(w http.ResponseWriter, req *http.Request) {
		httprequest.WriteJSON(w, http.StatusUnauthorized, params.Error{
			Message: "some error",
			Code:    params.CodeDischargeRequired,
		})
	},
	expectError:     `Get http://.*/: no error info found in discharge-required response error: some error`,
	expectErrorCode: params.CodeDischargeRequired,
}, {
	about: "discharge-required response with no macaroon",
	handler: func(w http.ResponseWriter, req *http.Request) {
		httprequest.WriteJSON(w, http.StatusUnauthorized, params.Error{
			Message: "some error",
			Code:    params.CodeDischargeRequired,
			Info: params.DischargeRequiredErrorInfo{
				MacaroonPath: "/",
			}.AsMap(),
		})
	},
	expectError: `Get http://.*/: no macaroon found in discharge-required response`,
}}

func (s *httpSuite) TestHTTPClient(c *gc.C) {
	var handler http.HandlerFunc
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		handler(w, req)
	}))
	defer srv.Close()
	s.client.BaseURL = srv.URL
	for i, test := range httpClientTests {
		c.Logf("test %d: %s", i, test.about)
		handler = test.handler
		var resp interface{}
		if test.expectResponse != nil {
			resp = reflect.New(reflect.TypeOf(test.expectResponse).Elem()).Interface()
		}
		err := s.client.Get(context.Background(), "/", resp)
		if test.expectError != "" {
			c.Check(err, gc.ErrorMatches, test.expectError)
			c.Check(params.ErrCode(err), gc.Equals, test.expectErrorCode)
			if test.expectErrorIs != "" {
				c.Check(errors.Cause(err), jc.ErrorIs, test.expectErrorIs)
			}
			if err, ok := errors.Cause(err).(*params.Error); ok {
				c.Check(err.Info, jc.DeepEquals, test.expectErrorInfo)
			} else if test.expectErrorInfo != nil {
				c.Fatalf("no error info found in error")
			}
			continue
		}
		c.Check(err, gc.IsNil)
		c.Check(resp, jc.DeepEquals, test.expectResponse)
	}
}

func (s *httpSuite) TestControllerMachineAuthForHostedModel(c *gc.C) {
	const nonce = "gary"

	srv := apiservertesting.NewAPIServer(func(modelUUID string) (interface{}, error) {
		return &testRootAPI{}, nil
	})
	s.AddCleanup(func(_ *gc.C) { srv.Close() })
	info := &api.Info{
		Addrs:          srv.Addrs,
		CACert:         testing.CACert,
		ControllerUUID: testing.ControllerTag.Id(),
		ModelTag:       testing.ModelTag,
		Tag:            names.NewMachineTag("1"),
		Password:       "password",
		Nonce:          nonce,
	}

	conn, err := api.Open(context.Background(), info, api.DialOpts{})
	c.Assert(err, jc.ErrorIsNil)
	httpClient, err := conn.HTTPClient()
	c.Assert(err, jc.ErrorIsNil)

	// Test with a dummy HTTP server returns the auth related headers used.
	httpSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()
		if ok {
			httprequest.WriteJSON(w, http.StatusOK, map[string]string{
				"username": username,
				"password": password,
				"nonce":    req.Header.Get(params.MachineNonceHeader),
			})
		} else {
			httprequest.WriteJSON(w, http.StatusUnauthorized, params.Error{
				Message: "no auth header",
			})
		}
	}))
	defer httpSrv.Close()
	httpClient.BaseURL = httpSrv.URL
	var out map[string]string
	c.Assert(httpClient.Get(context.Background(), "/", &out), jc.ErrorIsNil)
	c.Assert(out, gc.DeepEquals, map[string]string{
		"username": "machine-1",
		"password": "password",
		"nonce":    nonce,
	})
}

func (s *httpSuite) TestAuthHTTPRequest(c *gc.C) {
	apiInfo := &api.Info{}

	req := s.authHTTPRequest(c, apiInfo)
	_, _, ok := req.BasicAuth()
	c.Assert(ok, jc.IsFalse)
	c.Assert(req.Header, gc.HasLen, 2)
	c.Assert(req.Header.Get(httpbakery.BakeryProtocolHeader), gc.Equals, "3")
	c.Assert(req.Header.Get(params.JujuClientVersion), gc.Equals, version.Current.String())

	apiInfo.Nonce = "foo"
	req = s.authHTTPRequest(c, apiInfo)
	_, _, ok = req.BasicAuth()
	c.Assert(ok, jc.IsFalse)
	c.Assert(req.Header.Get(params.MachineNonceHeader), gc.Equals, "foo")
	c.Assert(req.Header.Get(httpbakery.BakeryProtocolHeader), gc.Equals, "3")

	apiInfo.Tag = names.NewMachineTag("123")
	apiInfo.Password = "password"
	req = s.authHTTPRequest(c, apiInfo)
	user, pass, ok := req.BasicAuth()
	c.Assert(ok, jc.IsTrue)
	c.Assert(user, gc.Equals, "machine-123")
	c.Assert(pass, gc.Equals, "password")
	c.Assert(req.Header.Get(params.MachineNonceHeader), gc.Equals, "foo")
	c.Assert(req.Header.Get(httpbakery.BakeryProtocolHeader), gc.Equals, "3")

	mac, err := jujutesting.NewMacaroon("id")
	c.Assert(err, jc.ErrorIsNil)
	apiInfo.Macaroons = []macaroon.Slice{{mac}}
	req = s.authHTTPRequest(c, apiInfo)
	c.Assert(req.Header.Get(params.MachineNonceHeader), gc.Equals, "foo")
	c.Assert(req.Header.Get(httpbakery.BakeryProtocolHeader), gc.Equals, "3")
	macaroons := httpbakery.RequestMacaroons(req)
	jujutesting.MacaroonsEqual(c, macaroons, apiInfo.Macaroons)
}

func (s *httpSuite) authHTTPRequest(c *gc.C, info *api.Info) *http.Request {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	c.Assert(err, jc.ErrorIsNil)
	err = api.AuthHTTPRequest(req, info)
	c.Assert(err, jc.ErrorIsNil)
	return req
}

// Note: the fact that the code works against the actual API server is
// well tested by some of the other API tests.
// This suite focuses on less reachable paths by changing
// the BaseURL of the httprequest.Client so that
// we can use our own custom servers.

func newString(s string) *string {
	return &s
}
