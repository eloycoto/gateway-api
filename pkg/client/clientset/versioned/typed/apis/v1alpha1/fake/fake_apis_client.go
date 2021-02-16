/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/service-apis/pkg/client/clientset/versioned/typed/apis/v1alpha1"
)

type FakeNetworkingV1alpha1 struct {
	*testing.Fake
}

func (c *FakeNetworkingV1alpha1) BackendPolicies(namespace string) v1alpha1.BackendPolicyInterface {
	return &FakeBackendPolicies{c, namespace}
}

func (c *FakeNetworkingV1alpha1) Gateways(namespace string) v1alpha1.GatewayInterface {
	return &FakeGateways{c, namespace}
}

func (c *FakeNetworkingV1alpha1) GatewayClasses() v1alpha1.GatewayClassInterface {
	return &FakeGatewayClasses{c}
}

func (c *FakeNetworkingV1alpha1) HTTPRoutes(namespace string) v1alpha1.HTTPRouteInterface {
	return &FakeHTTPRoutes{c, namespace}
}

func (c *FakeNetworkingV1alpha1) TCPRoutes(namespace string) v1alpha1.TCPRouteInterface {
	return &FakeTCPRoutes{c, namespace}
}

func (c *FakeNetworkingV1alpha1) TLSRoutes(namespace string) v1alpha1.TLSRouteInterface {
	return &FakeTLSRoutes{c, namespace}
}

func (c *FakeNetworkingV1alpha1) UDPRoutes(namespace string) v1alpha1.UDPRouteInterface {
	return &FakeUDPRoutes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkingV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
