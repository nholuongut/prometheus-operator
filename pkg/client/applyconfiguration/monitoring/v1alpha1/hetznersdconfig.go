// Copyright The Nho Luong DevOps
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	monitoringv1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "github.com/nholuongut/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
)

// HetznerSDConfigApplyConfiguration represents a declarative configuration of the HetznerSDConfig type for use
// with apply.
type HetznerSDConfigApplyConfiguration struct {
	Role                             *string                                 `json:"role,omitempty"`
	BasicAuth                        *v1.BasicAuthApplyConfiguration         `json:"basicAuth,omitempty"`
	Authorization                    *v1.SafeAuthorizationApplyConfiguration `json:"authorization,omitempty"`
	OAuth2                           *v1.OAuth2ApplyConfiguration            `json:"oauth2,omitempty"`
	v1.ProxyConfigApplyConfiguration `json:",inline"`
	FollowRedirects                  *bool                               `json:"followRedirects,omitempty"`
	EnableHTTP2                      *bool                               `json:"enableHTTP2,omitempty"`
	TLSConfig                        *v1.SafeTLSConfigApplyConfiguration `json:"tlsConfig,omitempty"`
	Port                             *int                                `json:"port,omitempty"`
	RefreshInterval                  *monitoringv1.Duration              `json:"refreshInterval,omitempty"`
}

// HetznerSDConfigApplyConfiguration constructs a declarative configuration of the HetznerSDConfig type for use with
// apply.
func HetznerSDConfig() *HetznerSDConfigApplyConfiguration {
	return &HetznerSDConfigApplyConfiguration{}
}

// WithRole sets the Role field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Role field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithRole(value string) *HetznerSDConfigApplyConfiguration {
	b.Role = &value
	return b
}

// WithBasicAuth sets the BasicAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BasicAuth field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithBasicAuth(value *v1.BasicAuthApplyConfiguration) *HetznerSDConfigApplyConfiguration {
	b.BasicAuth = value
	return b
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithAuthorization(value *v1.SafeAuthorizationApplyConfiguration) *HetznerSDConfigApplyConfiguration {
	b.Authorization = value
	return b
}

// WithOAuth2 sets the OAuth2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OAuth2 field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithOAuth2(value *v1.OAuth2ApplyConfiguration) *HetznerSDConfigApplyConfiguration {
	b.OAuth2 = value
	return b
}

// WithProxyURL sets the ProxyURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyURL field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithProxyURL(value string) *HetznerSDConfigApplyConfiguration {
	b.ProxyURL = &value
	return b
}

// WithNoProxy sets the NoProxy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NoProxy field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithNoProxy(value string) *HetznerSDConfigApplyConfiguration {
	b.NoProxy = &value
	return b
}

// WithProxyFromEnvironment sets the ProxyFromEnvironment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyFromEnvironment field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithProxyFromEnvironment(value bool) *HetznerSDConfigApplyConfiguration {
	b.ProxyFromEnvironment = &value
	return b
}

// WithProxyConnectHeader puts the entries into the ProxyConnectHeader field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ProxyConnectHeader field,
// overwriting an existing map entries in ProxyConnectHeader field with the same key.
func (b *HetznerSDConfigApplyConfiguration) WithProxyConnectHeader(entries map[string][]corev1.SecretKeySelector) *HetznerSDConfigApplyConfiguration {
	if b.ProxyConnectHeader == nil && len(entries) > 0 {
		b.ProxyConnectHeader = make(map[string][]corev1.SecretKeySelector, len(entries))
	}
	for k, v := range entries {
		b.ProxyConnectHeader[k] = v
	}
	return b
}

// WithFollowRedirects sets the FollowRedirects field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FollowRedirects field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithFollowRedirects(value bool) *HetznerSDConfigApplyConfiguration {
	b.FollowRedirects = &value
	return b
}

// WithEnableHTTP2 sets the EnableHTTP2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableHTTP2 field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithEnableHTTP2(value bool) *HetznerSDConfigApplyConfiguration {
	b.EnableHTTP2 = &value
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithTLSConfig(value *v1.SafeTLSConfigApplyConfiguration) *HetznerSDConfigApplyConfiguration {
	b.TLSConfig = value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithPort(value int) *HetznerSDConfigApplyConfiguration {
	b.Port = &value
	return b
}

// WithRefreshInterval sets the RefreshInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RefreshInterval field is set to the value of the last call.
func (b *HetznerSDConfigApplyConfiguration) WithRefreshInterval(value monitoringv1.Duration) *HetznerSDConfigApplyConfiguration {
	b.RefreshInterval = &value
	return b
}
