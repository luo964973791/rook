/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/rook/rook/pkg/apis/yugabytedb.rook.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// YBClusterLister helps list YBClusters.
type YBClusterLister interface {
	// List lists all YBClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.YBCluster, err error)
	// YBClusters returns an object that can list and get YBClusters.
	YBClusters(namespace string) YBClusterNamespaceLister
	YBClusterListerExpansion
}

// yBClusterLister implements the YBClusterLister interface.
type yBClusterLister struct {
	indexer cache.Indexer
}

// NewYBClusterLister returns a new YBClusterLister.
func NewYBClusterLister(indexer cache.Indexer) YBClusterLister {
	return &yBClusterLister{indexer: indexer}
}

// List lists all YBClusters in the indexer.
func (s *yBClusterLister) List(selector labels.Selector) (ret []*v1alpha1.YBCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.YBCluster))
	})
	return ret, err
}

// YBClusters returns an object that can list and get YBClusters.
func (s *yBClusterLister) YBClusters(namespace string) YBClusterNamespaceLister {
	return yBClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// YBClusterNamespaceLister helps list and get YBClusters.
type YBClusterNamespaceLister interface {
	// List lists all YBClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.YBCluster, err error)
	// Get retrieves the YBCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.YBCluster, error)
	YBClusterNamespaceListerExpansion
}

// yBClusterNamespaceLister implements the YBClusterNamespaceLister
// interface.
type yBClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all YBClusters in the indexer for a given namespace.
func (s yBClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.YBCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.YBCluster))
	})
	return ret, err
}

// Get retrieves the YBCluster from the indexer for a given namespace and name.
func (s yBClusterNamespaceLister) Get(name string) (*v1alpha1.YBCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ybcluster"), name)
	}
	return obj.(*v1alpha1.YBCluster), nil
}
