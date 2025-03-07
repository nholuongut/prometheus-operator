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
	v1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1"
)

// RuleGroupApplyConfiguration represents a declarative configuration of the RuleGroup type for use
// with apply.
type RuleGroupApplyConfiguration struct {
	Name                    *string                  `json:"name,omitempty"`
	Interval                *v1.Duration             `json:"interval,omitempty"`
	QueryOffset             *v1.Duration             `json:"query_offset,omitempty"`
	Rules                   []RuleApplyConfiguration `json:"rules,omitempty"`
	PartialResponseStrategy *string                  `json:"partial_response_strategy,omitempty"`
	Limit                   *int                     `json:"limit,omitempty"`
}

// RuleGroupApplyConfiguration constructs a declarative configuration of the RuleGroup type for use with
// apply.
func RuleGroup() *RuleGroupApplyConfiguration {
	return &RuleGroupApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *RuleGroupApplyConfiguration) WithName(value string) *RuleGroupApplyConfiguration {
	b.Name = &value
	return b
}

// WithInterval sets the Interval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Interval field is set to the value of the last call.
func (b *RuleGroupApplyConfiguration) WithInterval(value v1.Duration) *RuleGroupApplyConfiguration {
	b.Interval = &value
	return b
}

// WithQueryOffset sets the QueryOffset field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the QueryOffset field is set to the value of the last call.
func (b *RuleGroupApplyConfiguration) WithQueryOffset(value v1.Duration) *RuleGroupApplyConfiguration {
	b.QueryOffset = &value
	return b
}

// WithRules adds the given value to the Rules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Rules field.
func (b *RuleGroupApplyConfiguration) WithRules(values ...*RuleApplyConfiguration) *RuleGroupApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRules")
		}
		b.Rules = append(b.Rules, *values[i])
	}
	return b
}

// WithPartialResponseStrategy sets the PartialResponseStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PartialResponseStrategy field is set to the value of the last call.
func (b *RuleGroupApplyConfiguration) WithPartialResponseStrategy(value string) *RuleGroupApplyConfiguration {
	b.PartialResponseStrategy = &value
	return b
}

// WithLimit sets the Limit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Limit field is set to the value of the last call.
func (b *RuleGroupApplyConfiguration) WithLimit(value int) *RuleGroupApplyConfiguration {
	b.Limit = &value
	return b
}
