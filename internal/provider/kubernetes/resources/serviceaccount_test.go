// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package resources_test

import (
	"testing"

	"github.com/juju/errors"
	"github.com/juju/tc"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/juju/juju/internal/provider/kubernetes/resources"
)

type serviceAccountSuite struct {
	resourceSuite
}

func TestServiceAccountSuite(t *testing.T) {
	tc.Run(t, &serviceAccountSuite{})
}

func (s *serviceAccountSuite) TestApply(c *tc.C) {
	sa := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sa1",
			Namespace: "test",
		},
	}
	// Create.
	saResource := resources.NewServiceAccount("sa1", "test", sa)
	c.Assert(saResource.Apply(c.Context(), s.client), tc.ErrorIsNil)
	result, err := s.client.CoreV1().ServiceAccounts("test").Get(c.Context(), "sa1", metav1.GetOptions{})
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(len(result.GetAnnotations()), tc.Equals, 0)

	// Update.
	sa.SetAnnotations(map[string]string{"a": "b"})
	saResource = resources.NewServiceAccount("sa1", "test", sa)
	c.Assert(saResource.Apply(c.Context(), s.client), tc.ErrorIsNil)

	result, err = s.client.CoreV1().ServiceAccounts("test").Get(c.Context(), "sa1", metav1.GetOptions{})
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(result.GetName(), tc.Equals, `sa1`)
	c.Assert(result.GetNamespace(), tc.Equals, `test`)
	c.Assert(result.GetAnnotations(), tc.DeepEquals, map[string]string{"a": "b"})
}

func (s *serviceAccountSuite) TestGet(c *tc.C) {
	template := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sa1",
			Namespace: "test",
		},
	}
	sa1 := template
	sa1.SetAnnotations(map[string]string{"a": "b"})
	_, err := s.client.CoreV1().ServiceAccounts("test").Create(c.Context(), &sa1, metav1.CreateOptions{})
	c.Assert(err, tc.ErrorIsNil)

	saResource := resources.NewServiceAccount("sa1", "test", &template)
	c.Assert(len(saResource.GetAnnotations()), tc.Equals, 0)
	err = saResource.Get(c.Context(), s.client)
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(saResource.GetName(), tc.Equals, `sa1`)
	c.Assert(saResource.GetNamespace(), tc.Equals, `test`)
	c.Assert(saResource.GetAnnotations(), tc.DeepEquals, map[string]string{"a": "b"})
}

func (s *serviceAccountSuite) TestDelete(c *tc.C) {
	sa := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sa1",
			Namespace: "test",
		},
	}
	_, err := s.client.CoreV1().ServiceAccounts("test").Create(c.Context(), &sa, metav1.CreateOptions{})
	c.Assert(err, tc.ErrorIsNil)

	result, err := s.client.CoreV1().ServiceAccounts("test").Get(c.Context(), "sa1", metav1.GetOptions{})
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(result.GetName(), tc.Equals, `sa1`)

	saResource := resources.NewServiceAccount("sa1", "test", &sa)
	err = saResource.Delete(c.Context(), s.client)
	c.Assert(err, tc.ErrorIsNil)

	err = saResource.Get(c.Context(), s.client)
	c.Assert(err, tc.ErrorIs, errors.NotFound)

	_, err = s.client.CoreV1().ServiceAccounts("test").Get(c.Context(), "sa1", metav1.GetOptions{})
	c.Assert(err, tc.Satisfies, k8serrors.IsNotFound)
}

func (s *serviceAccountSuite) TestUpdate(c *tc.C) {
	sa := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sa1",
			Namespace: "test",
		},
	}
	_, err := s.client.CoreV1().ServiceAccounts("test").Create(
		c.Context(),
		&sa,
		metav1.CreateOptions{},
	)
	c.Assert(err, tc.ErrorIsNil)

	sa.ObjectMeta.Labels = map[string]string{
		"test": "label",
	}

	saResource := resources.NewServiceAccount("sa1", "test", &sa)
	err = saResource.Update(c.Context(), s.client)
	c.Assert(err, tc.ErrorIsNil)

	rsa, err := s.client.CoreV1().ServiceAccounts("test").Get(
		c.Context(),
		"sa1",
		metav1.GetOptions{},
	)
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(rsa, tc.DeepEquals, &saResource.ServiceAccount)
}

func (s *serviceAccountSuite) TestEnsureCreatesNew(c *tc.C) {
	sa := resources.NewServiceAccount("sa1", "test", &corev1.ServiceAccount{})
	cleanups, err := sa.Ensure(c.Context(), s.client)
	c.Assert(err, tc.ErrorIsNil)

	obj, err := s.client.CoreV1().ServiceAccounts("test").Get(
		c.Context(), "sa1", metav1.GetOptions{})
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(&sa.ServiceAccount, tc.DeepEquals, obj)

	for _, v := range cleanups {
		v()
	}

	// Test cleanup removes service account
	_, err = s.client.CoreV1().ServiceAccounts("test").Get(
		c.Context(), "sa1", metav1.GetOptions{})
	c.Assert(k8serrors.IsNotFound(err), tc.IsTrue)
}

func (s *serviceAccountSuite) TestEnsureUpdates(c *tc.C) {
	sa := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sa2",
			Namespace: "testing",
		},
	}

	_, err := s.client.CoreV1().ServiceAccounts("testing").Create(
		c.Context(), sa, metav1.CreateOptions{})
	c.Assert(err, tc.ErrorIsNil)

	sa.ObjectMeta.Labels = map[string]string{
		"test": "case",
	}

	resource := resources.NewServiceAccount("sa2", "testing", sa)
	_, err = resource.Ensure(c.Context(), s.client)
	c.Assert(err, tc.ErrorIsNil)

	obj, err := s.client.CoreV1().ServiceAccounts("testing").Get(
		c.Context(), sa.Name, metav1.GetOptions{})
	c.Assert(err, tc.ErrorIsNil)

	c.Assert(obj, tc.DeepEquals, &resource.ServiceAccount)
}
