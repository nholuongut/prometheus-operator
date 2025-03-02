// Copyright 2020 The Nho Luong DevOps
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

package operator

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	monitoringv1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1"
)

func MakeVolumeClaimTemplate(e monitoringv1.EmbeddedPersistentVolumeClaim) *v1.PersistentVolumeClaim {
	pvc := v1.PersistentVolumeClaim{
		TypeMeta: metav1.TypeMeta{
			APIVersion: e.APIVersion,
			Kind:       e.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        e.Name,
			Labels:      e.Labels,
			Annotations: e.Annotations,
		},
		Spec:   e.Spec,
		Status: e.Status,
	}
	return &pvc
}

// MakeHostAliases converts array of monitoringv1 HostAlias to array of corev1 HostAlias.
func MakeHostAliases(input []monitoringv1.HostAlias) []v1.HostAlias {
	if len(input) == 0 {
		return nil
	}

	output := make([]v1.HostAlias, len(input))

	for i, in := range input {
		output[i].Hostnames = in.Hostnames
		output[i].IP = in.IP
	}

	return output
}
