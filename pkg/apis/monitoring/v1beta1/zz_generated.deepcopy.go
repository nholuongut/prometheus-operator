//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	monitoringv1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfig) DeepCopyInto(out *AlertmanagerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfig.
func (in *AlertmanagerConfig) DeepCopy() *AlertmanagerConfig {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfigList) DeepCopyInto(out *AlertmanagerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*AlertmanagerConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AlertmanagerConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfigList.
func (in *AlertmanagerConfigList) DeepCopy() *AlertmanagerConfigList {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfigSpec) DeepCopyInto(out *AlertmanagerConfigSpec) {
	*out = *in
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(Route)
		(*in).DeepCopyInto(*out)
	}
	if in.Receivers != nil {
		in, out := &in.Receivers, &out.Receivers
		*out = make([]Receiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InhibitRules != nil {
		in, out := &in.InhibitRules, &out.InhibitRules
		*out = make([]InhibitRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeIntervals != nil {
		in, out := &in.TimeIntervals, &out.TimeIntervals
		*out = make([]TimeInterval, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfigSpec.
func (in *AlertmanagerConfigSpec) DeepCopy() *AlertmanagerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DayOfMonthRange) DeepCopyInto(out *DayOfMonthRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DayOfMonthRange.
func (in *DayOfMonthRange) DeepCopy() *DayOfMonthRange {
	if in == nil {
		return nil
	}
	out := new(DayOfMonthRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscordConfig) DeepCopyInto(out *DiscordConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	in.APIURL.DeepCopyInto(&out.APIURL)
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscordConfig.
func (in *DiscordConfig) DeepCopy() *DiscordConfig {
	if in == nil {
		return nil
	}
	out := new(DiscordConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailConfig) DeepCopyInto(out *EmailConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.AuthPassword != nil {
		in, out := &in.AuthPassword, &out.AuthPassword
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.AuthSecret != nil {
		in, out := &in.AuthSecret, &out.AuthSecret
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]KeyValue, len(*in))
		copy(*out, *in)
	}
	if in.HTML != nil {
		in, out := &in.HTML, &out.HTML
		*out = new(string)
		**out = **in
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = new(string)
		**out = **in
	}
	if in.RequireTLS != nil {
		in, out := &in.RequireTLS, &out.RequireTLS
		*out = new(bool)
		**out = **in
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(monitoringv1.SafeTLSConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailConfig.
func (in *EmailConfig) DeepCopy() *EmailConfig {
	if in == nil {
		return nil
	}
	out := new(EmailConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPConfig) DeepCopyInto(out *HTTPConfig) {
	*out = *in
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(monitoringv1.SafeAuthorization)
		(*in).DeepCopyInto(*out)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(monitoringv1.BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.OAuth2 != nil {
		in, out := &in.OAuth2, &out.OAuth2
		*out = new(monitoringv1.OAuth2)
		(*in).DeepCopyInto(*out)
	}
	if in.BearerTokenSecret != nil {
		in, out := &in.BearerTokenSecret, &out.BearerTokenSecret
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(monitoringv1.SafeTLSConfig)
		(*in).DeepCopyInto(*out)
	}
	in.ProxyConfig.DeepCopyInto(&out.ProxyConfig)
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPConfig.
func (in *HTTPConfig) DeepCopy() *HTTPConfig {
	if in == nil {
		return nil
	}
	out := new(HTTPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InhibitRule) DeepCopyInto(out *InhibitRule) {
	*out = *in
	if in.TargetMatch != nil {
		in, out := &in.TargetMatch, &out.TargetMatch
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.SourceMatch != nil {
		in, out := &in.SourceMatch, &out.SourceMatch
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.Equal != nil {
		in, out := &in.Equal, &out.Equal
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InhibitRule.
func (in *InhibitRule) DeepCopy() *InhibitRule {
	if in == nil {
		return nil
	}
	out := new(InhibitRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyValue) DeepCopyInto(out *KeyValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyValue.
func (in *KeyValue) DeepCopy() *KeyValue {
	if in == nil {
		return nil
	}
	out := new(KeyValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSTeamsConfig) DeepCopyInto(out *MSTeamsConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	in.WebhookURL.DeepCopyInto(&out.WebhookURL)
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(string)
		**out = **in
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = new(string)
		**out = **in
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSTeamsConfig.
func (in *MSTeamsConfig) DeepCopy() *MSTeamsConfig {
	if in == nil {
		return nil
	}
	out := new(MSTeamsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Matcher) DeepCopyInto(out *Matcher) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Matcher.
func (in *Matcher) DeepCopy() *Matcher {
	if in == nil {
		return nil
	}
	out := new(Matcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsGenieConfig) DeepCopyInto(out *OpsGenieConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]KeyValue, len(*in))
		copy(*out, *in)
	}
	if in.Responders != nil {
		in, out := &in.Responders, &out.Responders
		*out = make([]OpsGenieConfigResponder, len(*in))
		copy(*out, *in)
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsGenieConfig.
func (in *OpsGenieConfig) DeepCopy() *OpsGenieConfig {
	if in == nil {
		return nil
	}
	out := new(OpsGenieConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsGenieConfigResponder) DeepCopyInto(out *OpsGenieConfigResponder) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsGenieConfigResponder.
func (in *OpsGenieConfigResponder) DeepCopy() *OpsGenieConfigResponder {
	if in == nil {
		return nil
	}
	out := new(OpsGenieConfigResponder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyConfig) DeepCopyInto(out *PagerDutyConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.RoutingKey != nil {
		in, out := &in.RoutingKey, &out.RoutingKey
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.ServiceKey != nil {
		in, out := &in.ServiceKey, &out.ServiceKey
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]KeyValue, len(*in))
		copy(*out, *in)
	}
	if in.PagerDutyImageConfigs != nil {
		in, out := &in.PagerDutyImageConfigs, &out.PagerDutyImageConfigs
		*out = make([]PagerDutyImageConfig, len(*in))
		copy(*out, *in)
	}
	if in.PagerDutyLinkConfigs != nil {
		in, out := &in.PagerDutyLinkConfigs, &out.PagerDutyLinkConfigs
		*out = make([]PagerDutyLinkConfig, len(*in))
		copy(*out, *in)
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyConfig.
func (in *PagerDutyConfig) DeepCopy() *PagerDutyConfig {
	if in == nil {
		return nil
	}
	out := new(PagerDutyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyImageConfig) DeepCopyInto(out *PagerDutyImageConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyImageConfig.
func (in *PagerDutyImageConfig) DeepCopy() *PagerDutyImageConfig {
	if in == nil {
		return nil
	}
	out := new(PagerDutyImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyLinkConfig) DeepCopyInto(out *PagerDutyLinkConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyLinkConfig.
func (in *PagerDutyLinkConfig) DeepCopy() *PagerDutyLinkConfig {
	if in == nil {
		return nil
	}
	out := new(PagerDutyLinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushoverConfig) DeepCopyInto(out *PushoverConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.UserKey != nil {
		in, out := &in.UserKey, &out.UserKey
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.UserKeyFile != nil {
		in, out := &in.UserKeyFile, &out.UserKeyFile
		*out = new(string)
		**out = **in
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.TokenFile != nil {
		in, out := &in.TokenFile, &out.TokenFile
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(monitoringv1.Duration)
		**out = **in
	}
	if in.Device != nil {
		in, out := &in.Device, &out.Device
		*out = new(string)
		**out = **in
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushoverConfig.
func (in *PushoverConfig) DeepCopy() *PushoverConfig {
	if in == nil {
		return nil
	}
	out := new(PushoverConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Receiver) DeepCopyInto(out *Receiver) {
	*out = *in
	if in.OpsGenieConfigs != nil {
		in, out := &in.OpsGenieConfigs, &out.OpsGenieConfigs
		*out = make([]OpsGenieConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PagerDutyConfigs != nil {
		in, out := &in.PagerDutyConfigs, &out.PagerDutyConfigs
		*out = make([]PagerDutyConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DiscordConfigs != nil {
		in, out := &in.DiscordConfigs, &out.DiscordConfigs
		*out = make([]DiscordConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SlackConfigs != nil {
		in, out := &in.SlackConfigs, &out.SlackConfigs
		*out = make([]SlackConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebhookConfigs != nil {
		in, out := &in.WebhookConfigs, &out.WebhookConfigs
		*out = make([]WebhookConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WeChatConfigs != nil {
		in, out := &in.WeChatConfigs, &out.WeChatConfigs
		*out = make([]WeChatConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EmailConfigs != nil {
		in, out := &in.EmailConfigs, &out.EmailConfigs
		*out = make([]EmailConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VictorOpsConfigs != nil {
		in, out := &in.VictorOpsConfigs, &out.VictorOpsConfigs
		*out = make([]VictorOpsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PushoverConfigs != nil {
		in, out := &in.PushoverConfigs, &out.PushoverConfigs
		*out = make([]PushoverConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SNSConfigs != nil {
		in, out := &in.SNSConfigs, &out.SNSConfigs
		*out = make([]SNSConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TelegramConfigs != nil {
		in, out := &in.TelegramConfigs, &out.TelegramConfigs
		*out = make([]TelegramConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebexConfigs != nil {
		in, out := &in.WebexConfigs, &out.WebexConfigs
		*out = make([]WebexConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MSTeamsConfigs != nil {
		in, out := &in.MSTeamsConfigs, &out.MSTeamsConfigs
		*out = make([]MSTeamsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Receiver.
func (in *Receiver) DeepCopy() *Receiver {
	if in == nil {
		return nil
	}
	out := new(Receiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	if in.GroupBy != nil {
		in, out := &in.GroupBy, &out.GroupBy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Matchers != nil {
		in, out := &in.Matchers, &out.Matchers
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]v1.JSON, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MuteTimeIntervals != nil {
		in, out := &in.MuteTimeIntervals, &out.MuteTimeIntervals
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActiveTimeIntervals != nil {
		in, out := &in.ActiveTimeIntervals, &out.ActiveTimeIntervals
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSConfig) DeepCopyInto(out *SNSConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.Sigv4 != nil {
		in, out := &in.Sigv4, &out.Sigv4
		*out = new(monitoringv1.Sigv4)
		(*in).DeepCopyInto(*out)
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSConfig.
func (in *SNSConfig) DeepCopy() *SNSConfig {
	if in == nil {
		return nil
	}
	out := new(SNSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackAction) DeepCopyInto(out *SlackAction) {
	*out = *in
	if in.ConfirmField != nil {
		in, out := &in.ConfirmField, &out.ConfirmField
		*out = new(SlackConfirmationField)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackAction.
func (in *SlackAction) DeepCopy() *SlackAction {
	if in == nil {
		return nil
	}
	out := new(SlackAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackConfig) DeepCopyInto(out *SlackConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.APIURL != nil {
		in, out := &in.APIURL, &out.APIURL
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]SlackField, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MrkdwnIn != nil {
		in, out := &in.MrkdwnIn, &out.MrkdwnIn
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]SlackAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackConfig.
func (in *SlackConfig) DeepCopy() *SlackConfig {
	if in == nil {
		return nil
	}
	out := new(SlackConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackConfirmationField) DeepCopyInto(out *SlackConfirmationField) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackConfirmationField.
func (in *SlackConfirmationField) DeepCopy() *SlackConfirmationField {
	if in == nil {
		return nil
	}
	out := new(SlackConfirmationField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackField) DeepCopyInto(out *SlackField) {
	*out = *in
	if in.Short != nil {
		in, out := &in.Short, &out.Short
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackField.
func (in *SlackField) DeepCopy() *SlackField {
	if in == nil {
		return nil
	}
	out := new(SlackField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelegramConfig) DeepCopyInto(out *TelegramConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.BotToken != nil {
		in, out := &in.BotToken, &out.BotToken
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.BotTokenFile != nil {
		in, out := &in.BotTokenFile, &out.BotTokenFile
		*out = new(string)
		**out = **in
	}
	if in.DisableNotifications != nil {
		in, out := &in.DisableNotifications, &out.DisableNotifications
		*out = new(bool)
		**out = **in
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelegramConfig.
func (in *TelegramConfig) DeepCopy() *TelegramConfig {
	if in == nil {
		return nil
	}
	out := new(TelegramConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeInterval) DeepCopyInto(out *TimeInterval) {
	*out = *in
	if in.TimeIntervals != nil {
		in, out := &in.TimeIntervals, &out.TimeIntervals
		*out = make([]TimePeriod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeInterval.
func (in *TimeInterval) DeepCopy() *TimeInterval {
	if in == nil {
		return nil
	}
	out := new(TimeInterval)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimePeriod) DeepCopyInto(out *TimePeriod) {
	*out = *in
	if in.Times != nil {
		in, out := &in.Times, &out.Times
		*out = make([]TimeRange, len(*in))
		copy(*out, *in)
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]WeekdayRange, len(*in))
		copy(*out, *in)
	}
	if in.DaysOfMonth != nil {
		in, out := &in.DaysOfMonth, &out.DaysOfMonth
		*out = make([]DayOfMonthRange, len(*in))
		copy(*out, *in)
	}
	if in.Months != nil {
		in, out := &in.Months, &out.Months
		*out = make([]MonthRange, len(*in))
		copy(*out, *in)
	}
	if in.Years != nil {
		in, out := &in.Years, &out.Years
		*out = make([]YearRange, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimePeriod.
func (in *TimePeriod) DeepCopy() *TimePeriod {
	if in == nil {
		return nil
	}
	out := new(TimePeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRange) DeepCopyInto(out *TimeRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRange.
func (in *TimeRange) DeepCopy() *TimeRange {
	if in == nil {
		return nil
	}
	out := new(TimeRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VictorOpsConfig) DeepCopyInto(out *VictorOpsConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.CustomFields != nil {
		in, out := &in.CustomFields, &out.CustomFields
		*out = make([]KeyValue, len(*in))
		copy(*out, *in)
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VictorOpsConfig.
func (in *VictorOpsConfig) DeepCopy() *VictorOpsConfig {
	if in == nil {
		return nil
	}
	out := new(VictorOpsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeChatConfig) DeepCopyInto(out *WeChatConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.APISecret != nil {
		in, out := &in.APISecret, &out.APISecret
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeChatConfig.
func (in *WeChatConfig) DeepCopy() *WeChatConfig {
	if in == nil {
		return nil
	}
	out := new(WeChatConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebexConfig) DeepCopyInto(out *WebexConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.APIURL != nil {
		in, out := &in.APIURL, &out.APIURL
		*out = new(URL)
		**out = **in
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebexConfig.
func (in *WebexConfig) DeepCopy() *WebexConfig {
	if in == nil {
		return nil
	}
	out := new(WebexConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfig) DeepCopyInto(out *WebhookConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.URLSecret != nil {
		in, out := &in.URLSecret, &out.URLSecret
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfig.
func (in *WebhookConfig) DeepCopy() *WebhookConfig {
	if in == nil {
		return nil
	}
	out := new(WebhookConfig)
	in.DeepCopyInto(out)
	return out
}
