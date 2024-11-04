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
	v1 "k8s.io/api/core/v1"
)

// SlackConfigApplyConfiguration represents a declarative configuration of the SlackConfig type for use
// with apply.
type SlackConfigApplyConfiguration struct {
	SendResolved *bool                           `json:"sendResolved,omitempty"`
	APIURL       *v1.SecretKeySelector           `json:"apiURL,omitempty"`
	Channel      *string                         `json:"channel,omitempty"`
	Username     *string                         `json:"username,omitempty"`
	Color        *string                         `json:"color,omitempty"`
	Title        *string                         `json:"title,omitempty"`
	TitleLink    *string                         `json:"titleLink,omitempty"`
	Pretext      *string                         `json:"pretext,omitempty"`
	Text         *string                         `json:"text,omitempty"`
	Fields       []SlackFieldApplyConfiguration  `json:"fields,omitempty"`
	ShortFields  *bool                           `json:"shortFields,omitempty"`
	Footer       *string                         `json:"footer,omitempty"`
	Fallback     *string                         `json:"fallback,omitempty"`
	CallbackID   *string                         `json:"callbackId,omitempty"`
	IconEmoji    *string                         `json:"iconEmoji,omitempty"`
	IconURL      *string                         `json:"iconURL,omitempty"`
	ImageURL     *string                         `json:"imageURL,omitempty"`
	ThumbURL     *string                         `json:"thumbURL,omitempty"`
	LinkNames    *bool                           `json:"linkNames,omitempty"`
	MrkdwnIn     []string                        `json:"mrkdwnIn,omitempty"`
	Actions      []SlackActionApplyConfiguration `json:"actions,omitempty"`
	HTTPConfig   *HTTPConfigApplyConfiguration   `json:"httpConfig,omitempty"`
}

// SlackConfigApplyConfiguration constructs a declarative configuration of the SlackConfig type for use with
// apply.
func SlackConfig() *SlackConfigApplyConfiguration {
	return &SlackConfigApplyConfiguration{}
}

// WithSendResolved sets the SendResolved field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SendResolved field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithSendResolved(value bool) *SlackConfigApplyConfiguration {
	b.SendResolved = &value
	return b
}

// WithAPIURL sets the APIURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIURL field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithAPIURL(value v1.SecretKeySelector) *SlackConfigApplyConfiguration {
	b.APIURL = &value
	return b
}

// WithChannel sets the Channel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Channel field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithChannel(value string) *SlackConfigApplyConfiguration {
	b.Channel = &value
	return b
}

// WithUsername sets the Username field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Username field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithUsername(value string) *SlackConfigApplyConfiguration {
	b.Username = &value
	return b
}

// WithColor sets the Color field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Color field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithColor(value string) *SlackConfigApplyConfiguration {
	b.Color = &value
	return b
}

// WithTitle sets the Title field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Title field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithTitle(value string) *SlackConfigApplyConfiguration {
	b.Title = &value
	return b
}

// WithTitleLink sets the TitleLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TitleLink field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithTitleLink(value string) *SlackConfigApplyConfiguration {
	b.TitleLink = &value
	return b
}

// WithPretext sets the Pretext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pretext field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithPretext(value string) *SlackConfigApplyConfiguration {
	b.Pretext = &value
	return b
}

// WithText sets the Text field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Text field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithText(value string) *SlackConfigApplyConfiguration {
	b.Text = &value
	return b
}

// WithFields adds the given value to the Fields field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Fields field.
func (b *SlackConfigApplyConfiguration) WithFields(values ...*SlackFieldApplyConfiguration) *SlackConfigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFields")
		}
		b.Fields = append(b.Fields, *values[i])
	}
	return b
}

// WithShortFields sets the ShortFields field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ShortFields field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithShortFields(value bool) *SlackConfigApplyConfiguration {
	b.ShortFields = &value
	return b
}

// WithFooter sets the Footer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Footer field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithFooter(value string) *SlackConfigApplyConfiguration {
	b.Footer = &value
	return b
}

// WithFallback sets the Fallback field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Fallback field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithFallback(value string) *SlackConfigApplyConfiguration {
	b.Fallback = &value
	return b
}

// WithCallbackID sets the CallbackID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CallbackID field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithCallbackID(value string) *SlackConfigApplyConfiguration {
	b.CallbackID = &value
	return b
}

// WithIconEmoji sets the IconEmoji field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IconEmoji field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithIconEmoji(value string) *SlackConfigApplyConfiguration {
	b.IconEmoji = &value
	return b
}

// WithIconURL sets the IconURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IconURL field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithIconURL(value string) *SlackConfigApplyConfiguration {
	b.IconURL = &value
	return b
}

// WithImageURL sets the ImageURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageURL field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithImageURL(value string) *SlackConfigApplyConfiguration {
	b.ImageURL = &value
	return b
}

// WithThumbURL sets the ThumbURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ThumbURL field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithThumbURL(value string) *SlackConfigApplyConfiguration {
	b.ThumbURL = &value
	return b
}

// WithLinkNames sets the LinkNames field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LinkNames field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithLinkNames(value bool) *SlackConfigApplyConfiguration {
	b.LinkNames = &value
	return b
}

// WithMrkdwnIn adds the given value to the MrkdwnIn field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MrkdwnIn field.
func (b *SlackConfigApplyConfiguration) WithMrkdwnIn(values ...string) *SlackConfigApplyConfiguration {
	for i := range values {
		b.MrkdwnIn = append(b.MrkdwnIn, values[i])
	}
	return b
}

// WithActions adds the given value to the Actions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Actions field.
func (b *SlackConfigApplyConfiguration) WithActions(values ...*SlackActionApplyConfiguration) *SlackConfigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithActions")
		}
		b.Actions = append(b.Actions, *values[i])
	}
	return b
}

// WithHTTPConfig sets the HTTPConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPConfig field is set to the value of the last call.
func (b *SlackConfigApplyConfiguration) WithHTTPConfig(value *HTTPConfigApplyConfiguration) *SlackConfigApplyConfiguration {
	b.HTTPConfig = value
	return b
}
