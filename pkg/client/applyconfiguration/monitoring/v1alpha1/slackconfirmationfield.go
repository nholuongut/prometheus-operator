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

// SlackConfirmationFieldApplyConfiguration represents a declarative configuration of the SlackConfirmationField type for use
// with apply.
type SlackConfirmationFieldApplyConfiguration struct {
	Text        *string `json:"text,omitempty"`
	Title       *string `json:"title,omitempty"`
	OkText      *string `json:"okText,omitempty"`
	DismissText *string `json:"dismissText,omitempty"`
}

// SlackConfirmationFieldApplyConfiguration constructs a declarative configuration of the SlackConfirmationField type for use with
// apply.
func SlackConfirmationField() *SlackConfirmationFieldApplyConfiguration {
	return &SlackConfirmationFieldApplyConfiguration{}
}

// WithText sets the Text field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Text field is set to the value of the last call.
func (b *SlackConfirmationFieldApplyConfiguration) WithText(value string) *SlackConfirmationFieldApplyConfiguration {
	b.Text = &value
	return b
}

// WithTitle sets the Title field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Title field is set to the value of the last call.
func (b *SlackConfirmationFieldApplyConfiguration) WithTitle(value string) *SlackConfirmationFieldApplyConfiguration {
	b.Title = &value
	return b
}

// WithOkText sets the OkText field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OkText field is set to the value of the last call.
func (b *SlackConfirmationFieldApplyConfiguration) WithOkText(value string) *SlackConfirmationFieldApplyConfiguration {
	b.OkText = &value
	return b
}

// WithDismissText sets the DismissText field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DismissText field is set to the value of the last call.
func (b *SlackConfirmationFieldApplyConfiguration) WithDismissText(value string) *SlackConfirmationFieldApplyConfiguration {
	b.DismissText = &value
	return b
}
