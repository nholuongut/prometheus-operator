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

// AlertmanagerConfigSpecApplyConfiguration represents a declarative configuration of the AlertmanagerConfigSpec type for use
// with apply.
type AlertmanagerConfigSpecApplyConfiguration struct {
	Route         *RouteApplyConfiguration         `json:"route,omitempty"`
	Receivers     []ReceiverApplyConfiguration     `json:"receivers,omitempty"`
	InhibitRules  []InhibitRuleApplyConfiguration  `json:"inhibitRules,omitempty"`
	TimeIntervals []TimeIntervalApplyConfiguration `json:"timeIntervals,omitempty"`
}

// AlertmanagerConfigSpecApplyConfiguration constructs a declarative configuration of the AlertmanagerConfigSpec type for use with
// apply.
func AlertmanagerConfigSpec() *AlertmanagerConfigSpecApplyConfiguration {
	return &AlertmanagerConfigSpecApplyConfiguration{}
}

// WithRoute sets the Route field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Route field is set to the value of the last call.
func (b *AlertmanagerConfigSpecApplyConfiguration) WithRoute(value *RouteApplyConfiguration) *AlertmanagerConfigSpecApplyConfiguration {
	b.Route = value
	return b
}

// WithReceivers adds the given value to the Receivers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Receivers field.
func (b *AlertmanagerConfigSpecApplyConfiguration) WithReceivers(values ...*ReceiverApplyConfiguration) *AlertmanagerConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithReceivers")
		}
		b.Receivers = append(b.Receivers, *values[i])
	}
	return b
}

// WithInhibitRules adds the given value to the InhibitRules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InhibitRules field.
func (b *AlertmanagerConfigSpecApplyConfiguration) WithInhibitRules(values ...*InhibitRuleApplyConfiguration) *AlertmanagerConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithInhibitRules")
		}
		b.InhibitRules = append(b.InhibitRules, *values[i])
	}
	return b
}

// WithTimeIntervals adds the given value to the TimeIntervals field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TimeIntervals field.
func (b *AlertmanagerConfigSpecApplyConfiguration) WithTimeIntervals(values ...*TimeIntervalApplyConfiguration) *AlertmanagerConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTimeIntervals")
		}
		b.TimeIntervals = append(b.TimeIntervals, *values[i])
	}
	return b
}
