// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package caasrbacmapper_test

import (
	"sync"
	"testing"

	"github.com/juju/errors"
	"github.com/juju/tc"
	"go.uber.org/mock/gomock"
	core "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"

	loggertesting "github.com/juju/juju/internal/logger/testing"
	"github.com/juju/juju/internal/provider/kubernetes"
	"github.com/juju/juju/internal/provider/kubernetes/constants"
	"github.com/juju/juju/internal/provider/kubernetes/mocks"
	coretesting "github.com/juju/juju/internal/testing"
	"github.com/juju/juju/internal/worker/caasrbacmapper"
)

type MapperSuite struct {
	ctrl                    *gomock.Controller
	mockSAInformer          *mocks.MockServiceAccountInformer
	mockSALister            *mocks.MockServiceAccountLister
	mockSANamespaceLister   *mocks.MockServiceAccountNamespaceLister
	mockSharedIndexInformer *mocks.MockSharedIndexInformer
}

func TestMapperSuite(t *testing.T) {
	tc.Run(t, &MapperSuite{})
}

func (m *MapperSuite) SetUpTest(c *tc.C) {
	m.ctrl = gomock.NewController(c)
	m.mockSAInformer = mocks.NewMockServiceAccountInformer(m.ctrl)
	m.mockSharedIndexInformer = mocks.NewMockSharedIndexInformer(m.ctrl)
	m.mockSharedIndexInformer.EXPECT().HasSynced().AnyTimes().Return(true)
	m.mockSALister = mocks.NewMockServiceAccountLister(m.ctrl)
	m.mockSANamespaceLister = mocks.NewMockServiceAccountNamespaceLister(m.ctrl)
	m.mockSharedIndexInformer.EXPECT().Run(gomock.Any()).AnyTimes()
	m.mockSAInformer.EXPECT().Informer().AnyTimes().Return(m.mockSharedIndexInformer)
	m.mockSAInformer.EXPECT().Lister().AnyTimes().Return(m.mockSALister)
}

func (m *MapperSuite) TestMapperAdditionSync(c *tc.C) {
	defer m.ctrl.Finish()
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	var eventHandlers cache.ResourceEventHandler
	m.mockSharedIndexInformer.EXPECT().AddEventHandler(gomock.Any()).
		DoAndReturn(func(h cache.ResourceEventHandler) (cache.ResourceEventHandlerRegistration, error) {
			eventHandlers = h
			waitGroup.Done()
			return m.mockSharedIndexInformer, nil
		})

	mapper, err := caasrbacmapper.NewMapper(loggertesting.WrapCheckLog(c), m.mockSAInformer)
	c.Assert(err, tc.ErrorIsNil)
	waitGroup.Wait()

	appName := "test"
	name := "test-sa"
	namespace := "test-model"
	uid := types.UID("123")
	sa := &core.ServiceAccount{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    kubernetes.RBACLabels(appName, "test-model", "badf00d", "d0gf00d", false, constants.LabelVersion1),
			UID:       uid,
		},
	}

	waitGroup.Add(1)
	m.mockSALister.EXPECT().ServiceAccounts(gomock.Eq(namespace)).
		Return(m.mockSANamespaceLister)
	m.mockSANamespaceLister.EXPECT().Get(gomock.Eq(name)).
		DoAndReturn(func(_ string) (*core.ServiceAccount, error) {
			waitGroup.Done()
			return sa, nil
		})

	eventHandlers.OnAdd(sa, false)
	waitGroup.Wait()

	mapper.Kill()
	mapper.Wait()

	for a := coretesting.LongAttempt.Start(); a.Next(); {
		rAppName, err := mapper.AppNameForServiceAccount(uid)
		if err == nil {
			c.Assert(rAppName, tc.Equals, appName)
			break
		}
		if !a.HasNext() {
			c.Fatal("timed out waiting for service account app")
		}
	}
}

func (m *MapperSuite) TestRBACMapperUpdateSync(c *tc.C) {
	defer m.ctrl.Finish()
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	var eventHandlers cache.ResourceEventHandler
	m.mockSharedIndexInformer.EXPECT().AddEventHandler(gomock.Any()).
		DoAndReturn(func(h cache.ResourceEventHandler) (cache.ResourceEventHandlerRegistration, error) {
			eventHandlers = h
			waitGroup.Done()
			return m.mockSharedIndexInformer, nil
		})

	mapper, err := caasrbacmapper.NewMapper(loggertesting.WrapCheckLog(c), m.mockSAInformer)
	c.Assert(err, tc.ErrorIsNil)
	waitGroup.Wait()

	appName := "test"
	name := "test-sa"
	namespace := "test-model"
	uid := types.UID("123")
	sa := &core.ServiceAccount{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    kubernetes.RBACLabels(appName, "test-model", "deadbeef", "badf00d", false, constants.LabelVersion1),
			UID:       uid,
		},
	}

	waitGroup.Add(1)
	m.mockSALister.EXPECT().ServiceAccounts(gomock.Eq(namespace)).
		Return(m.mockSANamespaceLister).AnyTimes()
	m.mockSANamespaceLister.EXPECT().Get(gomock.Eq(name)).
		DoAndReturn(func(_ string) (*core.ServiceAccount, error) {
			waitGroup.Done()
			return sa, nil
		})

	eventHandlers.OnAdd(sa, false)
	waitGroup.Wait()

	for a := coretesting.LongAttempt.Start(); a.Next(); {
		rAppName, err := mapper.AppNameForServiceAccount(uid)
		if err == nil {
			c.Assert(rAppName, tc.Equals, appName)
			break
		}
		if !a.HasNext() {
			c.Fatal("timed out waiting for service account app")
		}
	}

	// Update SA with a new app name to check propagation
	appName = "test-2"
	sa2 := sa.DeepCopy()
	sa2.ObjectMeta.Labels = kubernetes.RBACLabels(appName, "test-model", "deadbeef", "badf00d", false, constants.LabelVersion1)
	waitGroup.Add(1)
	m.mockSANamespaceLister.EXPECT().Get(gomock.Eq(name)).
		DoAndReturn(func(_ string) (*core.ServiceAccount, error) {
			waitGroup.Done()
			return sa2, nil
		})

	//Send the update event, oldObj, newObj
	eventHandlers.OnUpdate(sa, sa2)
	waitGroup.Wait()

	mapper.Kill()
	mapper.Wait()

	for a := coretesting.LongAttempt.Start(); a.Next(); {
		rAppName, err := mapper.AppNameForServiceAccount(uid)
		if err == nil {
			c.Assert(rAppName, tc.Equals, appName)
			break
		}
		if !a.HasNext() {
			c.Fatal("timed out waiting for service account app")
		}
	}
}

func (m *MapperSuite) TestRBACMapperDeleteSync(c *tc.C) {
	defer m.ctrl.Finish()
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	var eventHandlers cache.ResourceEventHandler
	m.mockSharedIndexInformer.EXPECT().AddEventHandler(gomock.Any()).
		DoAndReturn(func(h cache.ResourceEventHandler) (cache.ResourceEventHandlerRegistration, error) {
			eventHandlers = h
			waitGroup.Done()
			return m.mockSharedIndexInformer, nil
		})

	mapper, err := caasrbacmapper.NewMapper(loggertesting.WrapCheckLog(c), m.mockSAInformer)
	c.Assert(err, tc.ErrorIsNil)
	waitGroup.Wait()

	appName := "test"
	name := "test-sa"
	namespace := "test-model"
	uid := types.UID("123")
	sa := &core.ServiceAccount{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    kubernetes.RBACLabels(appName, "test-model", "deadbeef", "badf00d", false, constants.LabelVersion1),
			UID:       uid,
		},
	}

	m.mockSALister.EXPECT().ServiceAccounts(gomock.Eq(namespace)).
		Return(m.mockSANamespaceLister).AnyTimes()
	m.mockSANamespaceLister.EXPECT().Get(gomock.Eq(name)).Return(sa, nil)
	eventHandlers.OnAdd(sa, false)

	for a := coretesting.LongAttempt.Start(); a.Next(); {
		rAppName, err := mapper.AppNameForServiceAccount(uid)
		if err == nil {
			c.Assert(rAppName, tc.Equals, appName)
			break
		}
		if !a.HasNext() {
			c.Fatal("timed out waiting for service account app")
		}
	}

	// Update SA with a new app name to check propagation
	m.mockSANamespaceLister.EXPECT().Get(gomock.Eq(name)).Return(
		nil, k8serrors.NewNotFound(core.Resource("serviceaccount"), name))

	//Send the delete event for the service account
	eventHandlers.OnDelete(sa)

	mapper.Kill()
	mapper.Wait()

	for a := coretesting.LongAttempt.Start(); a.Next(); {
		_, err = mapper.AppNameForServiceAccount(uid)
		if err != nil {
			c.Assert(err, tc.ErrorIs, errors.NotFound)
			break
		}
		if !a.HasNext() {
			c.Fatal("timed out waiting for service account app to be removed")
		}
	}
}

func (m *MapperSuite) TestRBACMapperNotFound(c *tc.C) {
	defer m.ctrl.Finish()
	m.mockSharedIndexInformer.EXPECT().AddEventHandler(gomock.Any())

	mapper, err := caasrbacmapper.NewMapper(loggertesting.WrapCheckLog(c), m.mockSAInformer)
	c.Assert(err, tc.ErrorIsNil)

	_, err = mapper.AppNameForServiceAccount("testing")
	c.Assert(err, tc.ErrorIs, errors.NotFound)
}
