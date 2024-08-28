/*
Copyright Red Hat, Inc.

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

package v1

import (
	v1 "github.com/operator-framework/api/pkg/operators/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// OperatorLister helps list Operators.
// All objects returned here must be treated as read-only.
type OperatorLister interface {
	// List lists all Operators in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Operator, err error)
	// Get retrieves the Operator from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Operator, error)
	OperatorListerExpansion
}

// operatorLister implements the OperatorLister interface.
type operatorLister struct {
	listers.ResourceIndexer[*v1.Operator]
}

// NewOperatorLister returns a new OperatorLister.
func NewOperatorLister(indexer cache.Indexer) OperatorLister {
	return &operatorLister{listers.New[*v1.Operator](indexer, v1.Resource("operator"))}
}
