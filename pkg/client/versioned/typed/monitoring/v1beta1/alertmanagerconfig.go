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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1beta1"
	monitoringv1beta1 "github.com/nholuongut/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1beta1"
	scheme "github.com/nholuongut/prometheus-operator/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AlertmanagerConfigsGetter has a method to return a AlertmanagerConfigInterface.
// A group's client should implement this interface.
type AlertmanagerConfigsGetter interface {
	AlertmanagerConfigs(namespace string) AlertmanagerConfigInterface
}

// AlertmanagerConfigInterface has methods to work with AlertmanagerConfig resources.
type AlertmanagerConfigInterface interface {
	Create(ctx context.Context, alertmanagerConfig *v1beta1.AlertmanagerConfig, opts v1.CreateOptions) (*v1beta1.AlertmanagerConfig, error)
	Update(ctx context.Context, alertmanagerConfig *v1beta1.AlertmanagerConfig, opts v1.UpdateOptions) (*v1beta1.AlertmanagerConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.AlertmanagerConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.AlertmanagerConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AlertmanagerConfig, err error)
	Apply(ctx context.Context, alertmanagerConfig *monitoringv1beta1.AlertmanagerConfigApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.AlertmanagerConfig, err error)
	AlertmanagerConfigExpansion
}

// alertmanagerConfigs implements AlertmanagerConfigInterface
type alertmanagerConfigs struct {
	*gentype.ClientWithListAndApply[*v1beta1.AlertmanagerConfig, *v1beta1.AlertmanagerConfigList, *monitoringv1beta1.AlertmanagerConfigApplyConfiguration]
}

// newAlertmanagerConfigs returns a AlertmanagerConfigs
func newAlertmanagerConfigs(c *MonitoringV1beta1Client, namespace string) *alertmanagerConfigs {
	return &alertmanagerConfigs{
		gentype.NewClientWithListAndApply[*v1beta1.AlertmanagerConfig, *v1beta1.AlertmanagerConfigList, *monitoringv1beta1.AlertmanagerConfigApplyConfiguration](
			"alertmanagerconfigs",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.AlertmanagerConfig { return &v1beta1.AlertmanagerConfig{} },
			func() *v1beta1.AlertmanagerConfigList { return &v1beta1.AlertmanagerConfigList{} }),
	}
}
