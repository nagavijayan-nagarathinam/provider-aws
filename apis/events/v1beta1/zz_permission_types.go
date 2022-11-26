/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConditionObservation struct {
}

type ConditionParameters struct {

	// Key for the condition. Valid values: aws:PrincipalOrgID.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Type of condition. Value values: StringEquals.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Value for the key.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PermissionObservation struct {

	// The statement ID of the EventBridge permission.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PermissionParameters struct {

	// The action that you are enabling the other account to perform. Defaults to events:PutEvents.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// The event bus to set the permissions on. If you omit this, the permissions are set on the default event bus.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/events/v1beta1.Bus
	// +crossplane:generate:reference:refFieldName=EventBusNameRefs
	// +crossplane:generate:reference:selectorFieldName=EventBusNameSelector
	// +kubebuilder:validation:Optional
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// Reference to a Bus in events to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameRefs *v1.Reference `json:"eventBusNameRefs,omitempty" tf:"-"`

	// Selector for a Bus in events to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameSelector *v1.Selector `json:"eventBusNameSelector,omitempty" tf:"-"`

	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify * to permit any account to put events to your default event bus, optionally limited by condition.
	// +kubebuilder:validation:Required
	Principal *string `json:"principal" tf:"principal,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// An identifier string for the external account that you are granting permissions to.
	// +kubebuilder:validation:Required
	StatementID *string `json:"statementId" tf:"statement_id,omitempty"`
}

// PermissionSpec defines the desired state of Permission
type PermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionParameters `json:"forProvider"`
}

// PermissionStatus defines the observed state of Permission.
type PermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Permission is the Schema for the Permissions API. Provides a resource to create an EventBridge permission to support cross-account events in the current account default event bus.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Permission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionSpec   `json:"spec"`
	Status            PermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionList contains a list of Permissions
type PermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Permission `json:"items"`
}

// Repository type metadata.
var (
	Permission_Kind             = "Permission"
	Permission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Permission_Kind}.String()
	Permission_KindAPIVersion   = Permission_Kind + "." + CRDGroupVersion.String()
	Permission_GroupVersionKind = CRDGroupVersion.WithKind(Permission_Kind)
)

func init() {
	SchemeBuilder.Register(&Permission{}, &PermissionList{})
}
