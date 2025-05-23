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

// VictorOpsConfigApplyConfiguration represents a declarative configuration of the VictorOpsConfig type for use
// with apply.
type VictorOpsConfigApplyConfiguration struct {
	SendResolved      *bool                                `json:"sendResolved,omitempty"`
	APIKey            *SecretKeySelectorApplyConfiguration `json:"apiKey,omitempty"`
	APIURL            *string                              `json:"apiUrl,omitempty"`
	RoutingKey        *string                              `json:"routingKey,omitempty"`
	MessageType       *string                              `json:"messageType,omitempty"`
	EntityDisplayName *string                              `json:"entityDisplayName,omitempty"`
	StateMessage      *string                              `json:"stateMessage,omitempty"`
	MonitoringTool    *string                              `json:"monitoringTool,omitempty"`
	CustomFields      []KeyValueApplyConfiguration         `json:"customFields,omitempty"`
	HTTPConfig        *HTTPConfigApplyConfiguration        `json:"httpConfig,omitempty"`
}

// VictorOpsConfigApplyConfiguration constructs a declarative configuration of the VictorOpsConfig type for use with
// apply.
func VictorOpsConfig() *VictorOpsConfigApplyConfiguration {
	return &VictorOpsConfigApplyConfiguration{}
}

// WithSendResolved sets the SendResolved field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SendResolved field is set to the value of the last call.
func (b *VictorOpsConfigApplyConfiguration) WithSendResolved(value bool) *VictorOpsConfigApplyConfiguration {
	b.SendResolved = &value
	return b
}

// WithAPIKey sets the APIKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIKey field is set to the value of the last call.
func (b *VictorOpsConfigApplyConfiguration) WithAPIKey(value *SecretKeySelectorApplyConfiguration) *VictorOpsConfigApplyConfiguration {
	b.APIKey = value
	return b
}

// WithAPIURL sets the APIURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIURL field is set to the value of the last call.
func (b *VictorOpsConfigApplyConfiguration) WithAPIURL(value string) *VictorOpsConfigApplyConfiguration {
	b.APIURL = &value
	return b
}

// WithRoutingKey sets the RoutingKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RoutingKey field is set to the value of the last call.
func (b *VictorOpsConfigApplyConfiguration) WithRoutingKey(value string) *VictorOpsConfigApplyConfiguration {
	b.RoutingKey = &value
	return b
}

// WithMessageType sets the MessageType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MessageType field is set to the value of the last call.
func (b *VictorOpsConfigApplyConfiguration) WithMessageType(value string) *VictorOpsConfigApplyConfiguration {
	b.MessageType = &value
	return b
}

// WithEntityDisplayName sets the EntityDisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EntityDisplayName field is set to the value of the last call.
func (b *VictorOpsConfigApplyConfiguration) WithEntityDisplayName(value string) *VictorOpsConfigApplyConfiguration {
	b.EntityDisplayName = &value
	return b
}

// WithStateMessage sets the StateMessage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StateMessage field is set to the value of the last call.
func (b *VictorOpsConfigApplyConfiguration) WithStateMessage(value string) *VictorOpsConfigApplyConfiguration {
	b.StateMessage = &value
	return b
}

// WithMonitoringTool sets the MonitoringTool field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MonitoringTool field is set to the value of the last call.
func (b *VictorOpsConfigApplyConfiguration) WithMonitoringTool(value string) *VictorOpsConfigApplyConfiguration {
	b.MonitoringTool = &value
	return b
}

// WithCustomFields adds the given value to the CustomFields field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CustomFields field.
func (b *VictorOpsConfigApplyConfiguration) WithCustomFields(values ...*KeyValueApplyConfiguration) *VictorOpsConfigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCustomFields")
		}
		b.CustomFields = append(b.CustomFields, *values[i])
	}
	return b
}

// WithHTTPConfig sets the HTTPConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPConfig field is set to the value of the last call.
func (b *VictorOpsConfigApplyConfiguration) WithHTTPConfig(value *HTTPConfigApplyConfiguration) *VictorOpsConfigApplyConfiguration {
	b.HTTPConfig = value
	return b
}
