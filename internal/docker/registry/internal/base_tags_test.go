// Copyright 2021 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package internal_test

import (
	"io"
	"net/http"
	"strings"

	"github.com/juju/tc"
	"go.uber.org/mock/gomock"

	"github.com/juju/juju/core/semversion"
	"github.com/juju/juju/internal/docker/registry/image"
	"github.com/juju/juju/internal/tools"
)

func (s *baseSuite) TestTagsPublicRegistry(c *tc.C) {
	// Use anonymous login for public repository.
	s.isPrivate = false
	reg, ctrl := s.getRegistry(c)
	defer ctrl.Finish()

	data := `
{"name":"jujuqa/jujud-operator","tags":["2.9.10.1","2.9.10.2","2.9.10"]}
`[1:]

	gomock.InOrder(
		s.mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
			c.Assert(req.Header, tc.DeepEquals, http.Header{})
			c.Assert(req.Method, tc.Equals, `GET`)
			c.Assert(req.URL.String(), tc.Equals, `https://example.com/v2/jujuqa/jujud-operator/tags/list`)
			return &http.Response{
				Request:    req,
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(nil),
				Header: http.Header{
					http.CanonicalHeaderKey("WWW-Authenticate"): []string{
						`Bearer realm="https://auth.example.com/token",service="registry.example.com",scope="repository:jujuqa/jujud-operator:pull"`,
					},
				},
			}, nil
		}),
		// Refresh OAuth Token without credential.
		s.mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(
			func(req *http.Request) (*http.Response, error) {
				c.Assert(req.Header, tc.DeepEquals, http.Header{})
				c.Assert(req.Method, tc.Equals, `GET`)
				c.Assert(req.URL.String(), tc.Equals, `https://auth.example.com/token?scope=repository%3Ajujuqa%2Fjujud-operator%3Apull&service=registry.example.com`)
				return &http.Response{
					Request:    req,
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"token": "jwt-token", "access_token": "jwt-token","expires_in": 300}`)),
				}, nil
			},
		),
		s.mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
			c.Assert(req.Header, tc.DeepEquals, http.Header{"Authorization": []string{"Bearer jwt-token"}})
			c.Assert(req.Method, tc.Equals, `GET`)
			c.Assert(req.URL.String(), tc.Equals, `https://example.com/v2/jujuqa/jujud-operator/tags/list`)
			resps := &http.Response{
				Request:    req,
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(data)),
			}
			return resps, nil
		}),
	)
	vers, err := reg.Tags("jujud-operator")
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(vers, tc.DeepEquals, tools.Versions{
		image.NewImageInfo(semversion.MustParse("2.9.10.1")),
		image.NewImageInfo(semversion.MustParse("2.9.10.2")),
		image.NewImageInfo(semversion.MustParse("2.9.10")),
	})
}

func (s *baseSuite) TestTagsPrivateRegistry(c *tc.C) {
	// Use v2 for private repository.
	s.isPrivate = true
	reg, ctrl := s.getRegistry(c)
	defer ctrl.Finish()

	data := `
{"name":"jujuqa/jujud-operator","tags":["2.9.10.1","2.9.10.2","2.9.10"]}
`[1:]

	gomock.InOrder(
		s.mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
			c.Assert(req.Header, tc.DeepEquals, http.Header{})
			c.Assert(req.Method, tc.Equals, `GET`)
			c.Assert(req.URL.String(), tc.Equals, `https://example.com/v2/jujuqa/jujud-operator/tags/list`)
			return &http.Response{
				Request:    req,
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(nil),
				Header: http.Header{
					http.CanonicalHeaderKey("WWW-Authenticate"): []string{
						`Bearer realm="https://auth.example.com/token",service="registry.example.com",scope="repository:jujuqa/jujud-operator:pull"`,
					},
				},
			}, nil
		}),
		// Refresh OAuth Token.
		s.mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(
			func(req *http.Request) (*http.Response, error) {
				c.Assert(req.Header, tc.DeepEquals, http.Header{"Authorization": []string{"Basic " + s.getAuthToken("username", "pwd")}})
				c.Assert(req.Method, tc.Equals, `GET`)
				c.Assert(req.URL.String(), tc.Equals, `https://auth.example.com/token?scope=repository%3Ajujuqa%2Fjujud-operator%3Apull&service=registry.example.com`)
				return &http.Response{
					Request:    req,
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"token": "jwt-token", "access_token": "jwt-token","expires_in": 300}`)),
				}, nil
			},
		),
		s.mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
			c.Assert(req.Header, tc.DeepEquals, http.Header{"Authorization": []string{"Bearer jwt-token"}})
			c.Assert(req.Method, tc.Equals, `GET`)
			c.Assert(req.URL.String(), tc.Equals, `https://example.com/v2/jujuqa/jujud-operator/tags/list`)
			resps := &http.Response{
				Request:    req,
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(data)),
			}
			return resps, nil
		}),
	)
	vers, err := reg.Tags("jujud-operator")
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(vers, tc.DeepEquals, tools.Versions{
		image.NewImageInfo(semversion.MustParse("2.9.10.1")),
		image.NewImageInfo(semversion.MustParse("2.9.10.2")),
		image.NewImageInfo(semversion.MustParse("2.9.10")),
	})
}

func (s *baseSuite) TestTagsErrorResponse(c *tc.C) {
	s.isPrivate = true
	reg, ctrl := s.getRegistry(c)
	defer ctrl.Finish()

	data := `
{"errors":[{"code":"UNAUTHORIZED","message":"authentication required"}]}
`[1:]

	gomock.InOrder(
		s.mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
			c.Assert(req.Header, tc.DeepEquals, http.Header{})
			c.Assert(req.Method, tc.Equals, `GET`)
			c.Assert(req.URL.String(), tc.Equals, `https://example.com/v2/jujuqa/jujud-operator/tags/list`)
			return &http.Response{
				Request:    req,
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(nil),
				Header: http.Header{
					http.CanonicalHeaderKey("WWW-Authenticate"): []string{
						`Bearer realm="https://auth.example.com/token",service="registry.example.com",scope="repository:jujuqa/jujud-operator:pull"`,
					},
				},
			}, nil
		}),
		// Refresh OAuth Token
		s.mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(
			func(req *http.Request) (*http.Response, error) {
				c.Assert(req.Header, tc.DeepEquals, http.Header{"Authorization": []string{"Basic " + s.getAuthToken("username", "pwd")}})
				c.Assert(req.Method, tc.Equals, `GET`)
				c.Assert(req.URL.String(), tc.Equals, `https://auth.example.com/token?scope=repository%3Ajujuqa%2Fjujud-operator%3Apull&service=registry.example.com`)
				return &http.Response{
					Request:    req,
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"token": "jwt-token", "access_token": "jwt-token","expires_in": 300}`)),
				}, nil
			},
		),
		s.mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
			c.Assert(req.Header, tc.DeepEquals, http.Header{"Authorization": []string{"Bearer jwt-token"}})
			c.Assert(req.Method, tc.Equals, `GET`)
			c.Assert(req.URL.String(), tc.Equals, `https://example.com/v2/jujuqa/jujud-operator/tags/list`)
			resps := &http.Response{
				Request:    req,
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader(data)),
			}
			return resps, nil
		}),
	)
	_, err := reg.Tags("jujud-operator")
	c.Assert(err, tc.ErrorMatches, `Get "https://example.com/v2/jujuqa/jujud-operator/tags/list": non-successful response status=403`)
}
