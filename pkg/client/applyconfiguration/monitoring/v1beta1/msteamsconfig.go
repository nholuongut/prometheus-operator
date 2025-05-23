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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
)

// MSTeamsConfigApplyConfiguration represents a declarative configuration of the MSTeamsConfig type for use
// with apply.
type MSTeamsConfigApplyConfiguration struct {
	SendResolved *bool                         `json:"sendResolved,omitempty"`
	WebhookURL   *v1.SecretKeySelector         `json:"webhookUrl,omitempty"`
	Title        *string                       `json:"title,omitempty"`
	Summary      *string                       `json:"summary,omitempty"`
	Text         *string                       `json:"text,omitempty"`
	HTTPConfig   *HTTPConfigApplyConfiguration `json:"httpConfig,omitempty"`
}

// MSTeamsConfigApplyConfiguration constructs a declarative configuration of the MSTeamsConfig type for use with
// apply.
func MSTeamsConfig() *MSTeamsConfigApplyConfiguration {
	return &MSTeamsConfigApplyConfiguration{}
}

// WithSendResolved sets the SendResolved field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SendResolved field is set to the value of the last call.
func (b *MSTeamsConfigApplyConfiguration) WithSendResolved(value bool) *MSTeamsConfigApplyConfiguration {
	b.SendResolved = &value
	return b
}

// WithWebhookURL sets the WebhookURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WebhookURL field is set to the value of the last call.
func (b *MSTeamsConfigApplyConfiguration) WithWebhookURL(value v1.SecretKeySelector) *MSTeamsConfigApplyConfiguration {
	b.WebhookURL = &value
	return b
}

// WithTitle sets the Title field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Title field is set to the value of the last call.
func (b *MSTeamsConfigApplyConfiguration) WithTitle(value string) *MSTeamsConfigApplyConfiguration {
	b.Title = &value
	return b
}

// WithSummary sets the Summary field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Summary field is set to the value of the last call.
func (b *MSTeamsConfigApplyConfiguration) WithSummary(value string) *MSTeamsConfigApplyConfiguration {
	b.Summary = &value
	return b
}

// WithText sets the Text field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Text field is set to the value of the last call.
func (b *MSTeamsConfigApplyConfiguration) WithText(value string) *MSTeamsConfigApplyConfiguration {
	b.Text = &value
	return b
}

// WithHTTPConfig sets the HTTPConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPConfig field is set to the value of the last call.
func (b *MSTeamsConfigApplyConfiguration) WithHTTPConfig(value *HTTPConfigApplyConfiguration) *MSTeamsConfigApplyConfiguration {
	b.HTTPConfig = value
	return b
}
