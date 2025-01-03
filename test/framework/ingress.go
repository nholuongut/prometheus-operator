// Copyright 2016 The Nho Luong DevOps
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

package framework

import (
	"context"
	"fmt"
	"os"
	"time"

	v1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func MakeBasicIngress(serviceName string, servicePort int) *networkv1.Ingress {
	return &networkv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: "monitoring",
		},
		Spec: networkv1.IngressSpec{
			Rules: []networkv1.IngressRule{
				{
					IngressRuleValue: networkv1.IngressRuleValue{
						HTTP: &networkv1.HTTPIngressRuleValue{
							Paths: []networkv1.HTTPIngressPath{
								{
									Backend: networkv1.IngressBackend{
										Service: &networkv1.IngressServiceBackend{
											Name: serviceName,
											Port: networkv1.ServiceBackendPort{
												Number: int32(servicePort),
											},
										},
									},
									Path: "/metrics",
								},
							},
						},
					},
				},
			},
		},
	}
}

func (f *Framework) CreateIngress(ctx context.Context, namespace string, i *networkv1.Ingress) error {
	if _, err := f.KubeClient.NetworkingV1().Ingresses(namespace).Create(ctx, i, metav1.CreateOptions{}); err != nil {
		return fmt.Errorf("creating ingress %v failed: %w", i.Name, err)
	}
	return nil
}

func (f *Framework) SetupNginxIngressControllerIncDefaultBackend(ctx context.Context, namespace string) error {
	// Create Nginx Ingress Replication Controller
	if err := f.createReplicationControllerViaYml(ctx, namespace, "./framework/resources/nxginx-ingress-controller.yml"); err != nil {
		return fmt.Errorf("creating nginx ingress replication controller failed: %w", err)
	}

	// Create Default HTTP Backend Replication Controller
	if err := f.createReplicationControllerViaYml(ctx, namespace, "./framework/resources/default-http-backend.yml"); err != nil {
		return fmt.Errorf("creating default http backend replication controller failed: %w", err)
	}

	// Create Default HTTP Backend Service
	manifest, err := os.Open("./framework/resources/default-http-backend-service.yml")
	if err != nil {
		return fmt.Errorf("reading default http backend service yaml failed: %w", err)
	}

	service := v1.Service{}
	err = yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&service)
	if err != nil {
		return fmt.Errorf("decoding http backend service yaml failed: %w", err)
	}

	_, err = f.KubeClient.CoreV1().Services(namespace).Create(ctx, &service, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("creating http backend service %v failed: %w", service.Name, err)
	}
	if err := f.WaitForServiceReady(ctx, namespace, service.Name); err != nil {
		return fmt.Errorf("waiting for http backend service %v timed out: %w", service.Name, err)
	}

	return nil
}

func (f *Framework) DeleteNginxIngressControllerIncDefaultBackend(ctx context.Context, namespace string) error {
	// Delete Nginx Ingress Replication Controller
	if err := f.deleteReplicationControllerViaYml(ctx, namespace, "./framework/resources/nxginx-ingress-controller.yml"); err != nil {
		return fmt.Errorf("deleting nginx ingress replication controller failed: %w", err)
	}

	// Delete Default HTTP Backend Replication Controller
	if err := f.deleteReplicationControllerViaYml(ctx, namespace, "./framework/resources/default-http-backend.yml"); err != nil {
		return fmt.Errorf("deleting default http backend replication controller failed: %w", err)
	}

	// Delete Default HTTP Backend Service
	manifest, err := os.Open("./framework/resources/default-http-backend-service.yml")
	if err != nil {
		return fmt.Errorf("reading default http backend service yaml failed: %w", err)
	}

	service := v1.Service{}
	err = yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&service)
	if err != nil {
		return fmt.Errorf("decoding http backend service yaml failed: %w", err)
	}

	if err := f.KubeClient.CoreV1().Services(namespace).Delete(ctx, service.Name, metav1.DeleteOptions{}); err != nil {
		return fmt.Errorf("deleting http backend service %v failed: %w", service.Name, err)
	}

	return nil
}

func (f *Framework) GetIngressIP(ctx context.Context, namespace string, ingressName string) (*string, error) {
	var ingress *networkv1.Ingress
	err := wait.PollUntilContextTimeout(ctx, time.Millisecond*500, time.Minute*5, false, func(ctx context.Context) (bool, error) {
		var err error
		ingress, err = f.KubeClient.NetworkingV1().Ingresses(namespace).Get(ctx, ingressName, metav1.GetOptions{})
		if err != nil {
			return false, fmt.Errorf("requesting the ingress %v failed: %w", ingressName, err)
		}
		ingresses := ingress.Status.LoadBalancer.Ingress
		if len(ingresses) != 0 {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return &ingress.Status.LoadBalancer.Ingress[0].IP, nil
}
