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

package alertmanager

import (
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/tools/cache"

	"github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1"
)

var (
	descAlertmanagerSpecReplicas = prometheus.NewDesc(
		"prometheus_operator_spec_replicas",
		"Number of expected replicas for the object.",
		[]string{
			"namespace",
			"name",
		}, nil,
	)
)

type alertmanagerCollector struct {
	stores []cache.Store
}

func newAlertmanagerCollectorForStores(s ...cache.Store) *alertmanagerCollector {
	return &alertmanagerCollector{stores: s}
}

// Describe implements the prometheus.Collector interface.
func (c *alertmanagerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- descAlertmanagerSpecReplicas
}

// Collect implements the prometheus.Collector interface.
func (c *alertmanagerCollector) Collect(ch chan<- prometheus.Metric) {
	for _, s := range c.stores {
		for _, p := range s.List() {
			c.collectAlertmanager(ch, p.(*v1.Alertmanager))
		}
	}
}

func (c *alertmanagerCollector) collectAlertmanager(ch chan<- prometheus.Metric, a *v1.Alertmanager) {
	replicas := float64(minReplicas)
	if a.Spec.Replicas != nil {
		replicas = float64(*a.Spec.Replicas)
	}
	ch <- prometheus.MustNewConstMetric(descAlertmanagerSpecReplicas, prometheus.GaugeValue, replicas, a.Namespace, a.Name)
}
