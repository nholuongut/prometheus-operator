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

// ThanosRulersGetter has a method to return a ThanosRulerInterface.
// A group's client should implement this interface.
type ThanosRulersGetter interface {
	ThanosRulers(namespace string) ThanosRulerInterface
}

// ThanosRulerInterface has methods to work with ThanosRuler resources.
type ThanosRulerInterface interface {
	Create(ctx context.Context, thanosRuler *v1.ThanosRuler, opts metav1.CreateOptions) (*v1.ThanosRuler, error)
	Update(ctx context.Context, thanosRuler *v1.ThanosRuler, opts metav1.UpdateOptions) (*v1.ThanosRuler, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, thanosRuler *v1.ThanosRuler, opts metav1.UpdateOptions) (*v1.ThanosRuler, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ThanosRuler, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ThanosRulerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ThanosRuler, err error)
	Apply(ctx context.Context, thanosRuler *monitoringv1.ThanosRulerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ThanosRuler, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, thanosRuler *monitoringv1.ThanosRulerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ThanosRuler, err error)
	ThanosRulerExpansion
}

// thanosRulers implements ThanosRulerInterface
type thanosRulers struct {
	*gentype.ClientWithListAndApply[*v1.ThanosRuler, *v1.ThanosRulerList, *monitoringv1.ThanosRulerApplyConfiguration]
}

// newThanosRulers returns a ThanosRulers
func newThanosRulers(c *MonitoringV1Client, namespace string) *thanosRulers {
	return &thanosRulers{
		gentype.NewClientWithListAndApply[*v1.ThanosRuler, *v1.ThanosRulerList, *monitoringv1.ThanosRulerApplyConfiguration](
			"thanosrulers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.ThanosRuler { return &v1.ThanosRuler{} },
			func() *v1.ThanosRulerList { return &v1.ThanosRulerList{} }),
	}
}
