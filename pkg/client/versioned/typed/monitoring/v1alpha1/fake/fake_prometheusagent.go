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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	monitoringv1alpha1 "github.com/nholuongut/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1alpha1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrometheusAgents implements PrometheusAgentInterface
type FakePrometheusAgents struct {
	Fake *FakeMonitoringV1alpha1
	ns   string
}

var prometheusagentsResource = v1alpha1.SchemeGroupVersion.WithResource("prometheusagents")

var prometheusagentsKind = v1alpha1.SchemeGroupVersion.WithKind("PrometheusAgent")

// Get takes name of the prometheusAgent, and returns the corresponding prometheusAgent object, and an error if there is any.
func (c *FakePrometheusAgents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PrometheusAgent, err error) {
	emptyResult := &v1alpha1.PrometheusAgent{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(prometheusagentsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrometheusAgent), err
}

// List takes label and field selectors, and returns the list of PrometheusAgents that match those selectors.
func (c *FakePrometheusAgents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PrometheusAgentList, err error) {
	emptyResult := &v1alpha1.PrometheusAgentList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(prometheusagentsResource, prometheusagentsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrometheusAgentList{ListMeta: obj.(*v1alpha1.PrometheusAgentList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrometheusAgentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prometheusAgents.
func (c *FakePrometheusAgents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(prometheusagentsResource, c.ns, opts))

}

// Create takes the representation of a prometheusAgent and creates it.  Returns the server's representation of the prometheusAgent, and an error, if there is any.
func (c *FakePrometheusAgents) Create(ctx context.Context, prometheusAgent *v1alpha1.PrometheusAgent, opts v1.CreateOptions) (result *v1alpha1.PrometheusAgent, err error) {
	emptyResult := &v1alpha1.PrometheusAgent{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(prometheusagentsResource, c.ns, prometheusAgent, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrometheusAgent), err
}

// Update takes the representation of a prometheusAgent and updates it. Returns the server's representation of the prometheusAgent, and an error, if there is any.
func (c *FakePrometheusAgents) Update(ctx context.Context, prometheusAgent *v1alpha1.PrometheusAgent, opts v1.UpdateOptions) (result *v1alpha1.PrometheusAgent, err error) {
	emptyResult := &v1alpha1.PrometheusAgent{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(prometheusagentsResource, c.ns, prometheusAgent, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrometheusAgent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrometheusAgents) UpdateStatus(ctx context.Context, prometheusAgent *v1alpha1.PrometheusAgent, opts v1.UpdateOptions) (result *v1alpha1.PrometheusAgent, err error) {
	emptyResult := &v1alpha1.PrometheusAgent{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(prometheusagentsResource, "status", c.ns, prometheusAgent, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrometheusAgent), err
}

// Delete takes name of the prometheusAgent and deletes it. Returns an error if one occurs.
func (c *FakePrometheusAgents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(prometheusagentsResource, c.ns, name, opts), &v1alpha1.PrometheusAgent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrometheusAgents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(prometheusagentsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrometheusAgentList{})
	return err
}

// Patch applies the patch and returns the patched prometheusAgent.
func (c *FakePrometheusAgents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PrometheusAgent, err error) {
	emptyResult := &v1alpha1.PrometheusAgent{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(prometheusagentsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrometheusAgent), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied prometheusAgent.
func (c *FakePrometheusAgents) Apply(ctx context.Context, prometheusAgent *monitoringv1alpha1.PrometheusAgentApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PrometheusAgent, err error) {
	if prometheusAgent == nil {
		return nil, fmt.Errorf("prometheusAgent provided to Apply must not be nil")
	}
	data, err := json.Marshal(prometheusAgent)
	if err != nil {
		return nil, err
	}
	name := prometheusAgent.Name
	if name == nil {
		return nil, fmt.Errorf("prometheusAgent.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.PrometheusAgent{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(prometheusagentsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrometheusAgent), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePrometheusAgents) ApplyStatus(ctx context.Context, prometheusAgent *monitoringv1alpha1.PrometheusAgentApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PrometheusAgent, err error) {
	if prometheusAgent == nil {
		return nil, fmt.Errorf("prometheusAgent provided to Apply must not be nil")
	}
	data, err := json.Marshal(prometheusAgent)
	if err != nil {
		return nil, err
	}
	name := prometheusAgent.Name
	if name == nil {
		return nil, fmt.Errorf("prometheusAgent.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.PrometheusAgent{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(prometheusagentsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrometheusAgent), err
}

// GetScale takes name of the prometheusAgent, and returns the corresponding scale object, and an error if there is any.
func (c *FakePrometheusAgents) GetScale(ctx context.Context, prometheusAgentName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithOptions(prometheusagentsResource, c.ns, "scale", prometheusAgentName, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakePrometheusAgents) UpdateScale(ctx context.Context, prometheusAgentName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(prometheusagentsResource, "scale", c.ns, scale, opts), &autoscalingv1.Scale{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}
