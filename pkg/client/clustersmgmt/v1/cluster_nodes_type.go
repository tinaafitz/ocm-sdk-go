/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// ClusterNodes represents the values of the 'cluster_nodes' type.
//
// Counts of different classes of nodes inside a cluster.
type ClusterNodes struct {
	total   *int
	master  *int
	infra   *int
	compute *int
}

// Total returns the value of the 'total' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Total number of nodes of the cluster.
func (o *ClusterNodes) Total() int {
	if o != nil && o.total != nil {
		return *o.total
	}
	return 0
}

// GetTotal returns the value of the 'total' attribute and
// a flag indicating if the attribute has a value.
//
// Total number of nodes of the cluster.
func (o *ClusterNodes) GetTotal() (value int, ok bool) {
	ok = o != nil && o.total != nil
	if ok {
		value = *o.total
	}
	return
}

// Master returns the value of the 'master' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of master nodes of the cluster.
func (o *ClusterNodes) Master() int {
	if o != nil && o.master != nil {
		return *o.master
	}
	return 0
}

// GetMaster returns the value of the 'master' attribute and
// a flag indicating if the attribute has a value.
//
// Number of master nodes of the cluster.
func (o *ClusterNodes) GetMaster() (value int, ok bool) {
	ok = o != nil && o.master != nil
	if ok {
		value = *o.master
	}
	return
}

// Infra returns the value of the 'infra' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of infrastructure nodes of the cluster.
func (o *ClusterNodes) Infra() int {
	if o != nil && o.infra != nil {
		return *o.infra
	}
	return 0
}

// GetInfra returns the value of the 'infra' attribute and
// a flag indicating if the attribute has a value.
//
// Number of infrastructure nodes of the cluster.
func (o *ClusterNodes) GetInfra() (value int, ok bool) {
	ok = o != nil && o.infra != nil
	if ok {
		value = *o.infra
	}
	return
}

// Compute returns the value of the 'compute' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of compute nodes of the cluster.
func (o *ClusterNodes) Compute() int {
	if o != nil && o.compute != nil {
		return *o.compute
	}
	return 0
}

// GetCompute returns the value of the 'compute' attribute and
// a flag indicating if the attribute has a value.
//
// Number of compute nodes of the cluster.
func (o *ClusterNodes) GetCompute() (value int, ok bool) {
	ok = o != nil && o.compute != nil
	if ok {
		value = *o.compute
	}
	return
}

// ClusterNodesList is a list of values of the 'cluster_nodes' type.
type ClusterNodesList struct {
	items []*ClusterNodes
}

// Len returns the length of the list.
func (l *ClusterNodesList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterNodesList) Slice() []*ClusterNodes {
	var slice []*ClusterNodes
	if l == nil {
		slice = make([]*ClusterNodes, 0)
	} else {
		slice = make([]*ClusterNodes, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterNodesList) Each(f func(item *ClusterNodes) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ClusterNodesList) Range(f func(index int, item *ClusterNodes) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
