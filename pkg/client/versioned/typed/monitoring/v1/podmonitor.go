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

package v1

import (
	"context"

	v1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1 "github.com/nholuongut/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	scheme "github.com/nholuongut/prometheus-operator/pkg/client/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// PodMonitorsGetter has a method to return a PodMonitorInterface.
// A group's client should implement this interface.
type PodMonitorsGetter interface {
	PodMonitors(namespace string) PodMonitorInterface
}

// PodMonitorInterface has methods to work with PodMonitor resources.
type PodMonitorInterface interface {
	Create(ctx context.Context, podMonitor *v1.PodMonitor, opts metav1.CreateOptions) (*v1.PodMonitor, error)
	Update(ctx context.Context, podMonitor *v1.PodMonitor, opts metav1.UpdateOptions) (*v1.PodMonitor, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PodMonitor, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PodMonitorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PodMonitor, err error)
	Apply(ctx context.Context, podMonitor *monitoringv1.PodMonitorApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PodMonitor, err error)
	PodMonitorExpansion
}

// podMonitors implements PodMonitorInterface
type podMonitors struct {
	*gentype.ClientWithListAndApply[*v1.PodMonitor, *v1.PodMonitorList, *monitoringv1.PodMonitorApplyConfiguration]
}

// newPodMonitors returns a PodMonitors
func newPodMonitors(c *MonitoringV1Client, namespace string) *podMonitors {
	return &podMonitors{
		gentype.NewClientWithListAndApply[*v1.PodMonitor, *v1.PodMonitorList, *monitoringv1.PodMonitorApplyConfiguration](
			"podmonitors",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.PodMonitor { return &v1.PodMonitor{} },
			func() *v1.PodMonitorList { return &v1.PodMonitorList{} }),
	}
}
