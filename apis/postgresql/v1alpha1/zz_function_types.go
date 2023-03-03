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

type ArgObservation struct {
}

type ArgParameters struct {

	// An expression to be used as default value if the parameter is not specified.
	// An expression to be used as default value if the parameter is not specified.
	// +kubebuilder:validation:Optional
	Default *string `json:"default,omitempty" tf:"default,omitempty"`

	// Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
	// The argument mode. One of: IN, OUT, INOUT, or VARIADIC
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the function.
	// The argument name. The name may be required for some languages or depending on the argument mode.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of the argument.
	// The argument type.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type FunctionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FunctionParameters struct {

	// List of arguments for the function.
	// Function argument definitions.
	// +kubebuilder:validation:Optional
	Arg []ArgParameters `json:"arg,omitempty" tf:"arg,omitempty"`

	// Function body.
	// This should be everything after the return type in the function definition.
	// Body of the function.
	// +kubebuilder:validation:Required
	Body *string `json:"body" tf:"body,omitempty"`

	// The database where the function is located.
	// If not specified, the function is created in the current database.
	// The database where the function is located. If not specified, the provider default database is used.
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// True to automatically drop objects that depend on the function (such as
	// operators or triggers), and in turn all objects that depend on those objects. Default is false.
	// Automatically drop objects that depend on the function (such as operators or triggers), and in turn all objects that depend on those objects.
	// +kubebuilder:validation:Optional
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	// The name of the function.
	// Name of the function.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Type that the function returns.
	// Function return type.
	// +kubebuilder:validation:Optional
	Returns *string `json:"returns,omitempty" tf:"returns,omitempty"`

	// The schema where the function is located.
	// If not specified, the function is created in the current schema.
	// Schema where the function is located. If not specified, the provider default schema is used.
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API. Creates and manages a function on a PostgreSQL server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,postgresql}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
