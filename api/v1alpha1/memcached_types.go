/*
Copyright 2024.

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

package v1alpha1

import (
	"github.com/PawelK2012/memcached-operator/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MemcachedSpec defines the desired state of Memcached
type MemcachedSpec struct {
	// +kubebuilder:validation:Minimum=0
	// Size is the size of the memcached deployment
	Size int32 `json:"size"`
}

// MemcachedStatus defines the observed state of Memcached
type MemcachedStatus struct {
	// Nodes are the names of the memcached pods
	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Memcached is the Schema for the memcacheds API
type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec   `json:"spec,omitempty"`
	Status MemcachedStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MemcachedList contains a list of Memcached
type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcached `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Memcached{}, &MemcachedList{})
}

// ConvertTo converts this Memcached to the Hub version (vbeta1).
func (src *Memcached) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Memcached)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.DisableEvictions = true //bools are false by default, so we'll set this to true so we can tell our hook is running
	dst.Spec.Size = src.Spec.Size
	dst.Status.Nodes = src.Status.Nodes

	return nil
}

// ConvertFrom converts from the Hub version (vbeta1) to this version.
func (dst *Memcached) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Memcached)
	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Size = src.Spec.Size

	dst.Status.Nodes = src.Status.Nodes

	return nil
}
