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

// AzureADApplyConfiguration represents a declarative configuration of the AzureAD type for use
// with apply.
type AzureADApplyConfiguration struct {
	Cloud           *string                            `json:"cloud,omitempty"`
	ManagedIdentity *ManagedIdentityApplyConfiguration `json:"managedIdentity,omitempty"`
	OAuth           *AzureOAuthApplyConfiguration      `json:"oauth,omitempty"`
	SDK             *AzureSDKApplyConfiguration        `json:"sdk,omitempty"`
}

// AzureADApplyConfiguration constructs a declarative configuration of the AzureAD type for use with
// apply.
func AzureAD() *AzureADApplyConfiguration {
	return &AzureADApplyConfiguration{}
}

// WithCloud sets the Cloud field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cloud field is set to the value of the last call.
func (b *AzureADApplyConfiguration) WithCloud(value string) *AzureADApplyConfiguration {
	b.Cloud = &value
	return b
}

// WithManagedIdentity sets the ManagedIdentity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagedIdentity field is set to the value of the last call.
func (b *AzureADApplyConfiguration) WithManagedIdentity(value *ManagedIdentityApplyConfiguration) *AzureADApplyConfiguration {
	b.ManagedIdentity = value
	return b
}

// WithOAuth sets the OAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OAuth field is set to the value of the last call.
func (b *AzureADApplyConfiguration) WithOAuth(value *AzureOAuthApplyConfiguration) *AzureADApplyConfiguration {
	b.OAuth = value
	return b
}

// WithSDK sets the SDK field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SDK field is set to the value of the last call.
func (b *AzureADApplyConfiguration) WithSDK(value *AzureSDKApplyConfiguration) *AzureADApplyConfiguration {
	b.SDK = value
	return b
}
