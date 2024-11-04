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

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// AuthorizationApplyConfiguration represents a declarative configuration of the Authorization type for use
// with apply.
type AuthorizationApplyConfiguration struct {
	SafeAuthorizationApplyConfiguration `json:",inline"`
	CredentialsFile                     *string `json:"credentialsFile,omitempty"`
}

// AuthorizationApplyConfiguration constructs a declarative configuration of the Authorization type for use with
// apply.
func Authorization() *AuthorizationApplyConfiguration {
	return &AuthorizationApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *AuthorizationApplyConfiguration) WithType(value string) *AuthorizationApplyConfiguration {
	b.Type = &value
	return b
}

// WithCredentials sets the Credentials field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Credentials field is set to the value of the last call.
func (b *AuthorizationApplyConfiguration) WithCredentials(value corev1.SecretKeySelector) *AuthorizationApplyConfiguration {
	b.Credentials = &value
	return b
}

// WithCredentialsFile sets the CredentialsFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CredentialsFile field is set to the value of the last call.
func (b *AuthorizationApplyConfiguration) WithCredentialsFile(value string) *AuthorizationApplyConfiguration {
	b.CredentialsFile = &value
	return b
}
