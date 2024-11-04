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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ServiceMonitorLister helps list ServiceMonitors.
// All objects returned here must be treated as read-only.
type ServiceMonitorLister interface {
	// List lists all ServiceMonitors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ServiceMonitor, err error)
	// ServiceMonitors returns an object that can list and get ServiceMonitors.
	ServiceMonitors(namespace string) ServiceMonitorNamespaceLister
	ServiceMonitorListerExpansion
}

// serviceMonitorLister implements the ServiceMonitorLister interface.
type serviceMonitorLister struct {
	listers.ResourceIndexer[*v1.ServiceMonitor]
}

// NewServiceMonitorLister returns a new ServiceMonitorLister.
func NewServiceMonitorLister(indexer cache.Indexer) ServiceMonitorLister {
	return &serviceMonitorLister{listers.New[*v1.ServiceMonitor](indexer, v1.Resource("servicemonitor"))}
}

// ServiceMonitors returns an object that can list and get ServiceMonitors.
func (s *serviceMonitorLister) ServiceMonitors(namespace string) ServiceMonitorNamespaceLister {
	return serviceMonitorNamespaceLister{listers.NewNamespaced[*v1.ServiceMonitor](s.ResourceIndexer, namespace)}
}

// ServiceMonitorNamespaceLister helps list and get ServiceMonitors.
// All objects returned here must be treated as read-only.
type ServiceMonitorNamespaceLister interface {
	// List lists all ServiceMonitors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ServiceMonitor, err error)
	// Get retrieves the ServiceMonitor from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ServiceMonitor, error)
	ServiceMonitorNamespaceListerExpansion
}

// serviceMonitorNamespaceLister implements the ServiceMonitorNamespaceLister
// interface.
type serviceMonitorNamespaceLister struct {
	listers.ResourceIndexer[*v1.ServiceMonitor]
}
