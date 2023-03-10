/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ReplicationSlotObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReplicationSlotParameters struct {
}

// ReplicationSlotSpec defines the desired state of ReplicationSlot
type ReplicationSlotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationSlotParameters `json:"forProvider"`
}

// ReplicationSlotStatus defines the observed state of ReplicationSlot.
type ReplicationSlotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationSlotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationSlot is the Schema for the ReplicationSlots API. Creates and manages a physical replication slot on a PostgreSQL server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,postgresql}
type ReplicationSlot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationSlotSpec   `json:"spec"`
	Status            ReplicationSlotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationSlotList contains a list of ReplicationSlots
type ReplicationSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationSlot `json:"items"`
}

// Repository type metadata.
var (
	ReplicationSlot_Kind             = "ReplicationSlot"
	ReplicationSlot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicationSlot_Kind}.String()
	ReplicationSlot_KindAPIVersion   = ReplicationSlot_Kind + "." + CRDGroupVersion.String()
	ReplicationSlot_GroupVersionKind = CRDGroupVersion.WithKind(ReplicationSlot_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicationSlot{}, &ReplicationSlotList{})
}
