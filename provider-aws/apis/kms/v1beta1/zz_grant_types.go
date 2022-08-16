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

type ConstraintsObservation struct {
}

type ConstraintsParameters struct {

	// A list of key-value pairs that must match the encryption context in subsequent cryptographic operation requests. The grant allows the operation only when the encryption context in the request is the same as the encryption context specified in this constraint. Conflicts with encryption_context_subset.
	// +kubebuilder:validation:Optional
	EncryptionContextEquals map[string]*string `json:"encryptionContextEquals,omitempty" tf:"encryption_context_equals,omitempty"`

	// A list of key-value pairs that must be included in the encryption context of subsequent cryptographic operation requests. The grant allows the cryptographic operation only when the encryption context in the request includes the key-value pairs specified in this constraint, although it can include additional key-value pairs. Conflicts with encryption_context_equals.
	// +kubebuilder:validation:Optional
	EncryptionContextSubset map[string]*string `json:"encryptionContextSubset,omitempty" tf:"encryption_context_subset,omitempty"`
}

type GrantObservation struct {

	// The unique identifier for the grant.
	GrantID *string `json:"grantId,omitempty" tf:"grant_id,omitempty"`

	// The grant token for the created grant. For more information, see Grant Tokens.
	GrantToken *string `json:"grantToken,omitempty" tf:"grant_token,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GrantParameters struct {

	// A structure that you can use to allow certain operations in the grant only when the desired encryption context is present. For more information about encryption context, see Encryption Context.
	// +kubebuilder:validation:Optional
	Constraints []ConstraintsParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// A list of grant tokens to be used when creating the grant. See Grant Tokens for more information about grant tokens.
	// +kubebuilder:validation:Optional
	GrantCreationTokens []*string `json:"grantCreationTokens,omitempty" tf:"grant_creation_tokens,omitempty"`

	// The principal that is given permission to perform the operations that the grant permits in ARN format. Note that due to eventual consistency issues around IAM principals, terraform's state may not always be refreshed to reflect what is true in AWS.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	GranteePrincipal *string `json:"granteePrincipal,omitempty" tf:"grantee_principal,omitempty"`

	// +kubebuilder:validation:Optional
	GranteePrincipalRef *v1.Reference `json:"granteePrincipalRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	GranteePrincipalSelector *v1.Selector `json:"granteePrincipalSelector,omitempty" tf:"-"`

	// The unique identifier for the customer master key  that the grant applies to. Specify the key ID or the Amazon Resource Name  of the CMK. To specify a CMK in a different AWS account, you must use the key ARN.
	// +crossplane:generate:reference:type=Key
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`

	// A friendly name for identifying the grant.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of operations that the grant permits. The permitted values are: Decrypt, Encrypt, GenerateDataKey, GenerateDataKeyWithoutPlaintext, ReEncryptFrom, ReEncryptTo, Sign, Verify, GetPublicKey, CreateGrant, RetireGrant, DescribeKey, GenerateDataKeyPair, or GenerateDataKeyPairWithoutPlaintext.
	// +kubebuilder:validation:Required
	Operations []*string `json:"operations" tf:"operations,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Defaults to false, Forces new resources) If set to false  the grants will be revoked upon deletion, and if set to true the grants will try to be retired upon deletion. Note that retiring grants requires special permissions, hence why we default to revoking grants.
	// See RetireGrant for more information.
	// +kubebuilder:validation:Optional
	RetireOnDelete *bool `json:"retireOnDelete,omitempty" tf:"retire_on_delete,omitempty"`

	// The principal that is given permission to retire the grant by using RetireGrant operation in ARN format. Note that due to eventual consistency issues around IAM principals, terraform's state may not always be refreshed to reflect what is true in AWS.
	// +kubebuilder:validation:Optional
	RetiringPrincipal *string `json:"retiringPrincipal,omitempty" tf:"retiring_principal,omitempty"`
}

// GrantSpec defines the desired state of Grant
type GrantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GrantParameters `json:"forProvider"`
}

// GrantStatus defines the observed state of Grant.
type GrantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Grant is the Schema for the Grants API. Provides a resource-based access control mechanism for KMS Customer Master Keys.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Grant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GrantSpec   `json:"spec"`
	Status            GrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrantList contains a list of Grants
type GrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Grant `json:"items"`
}

// Repository type metadata.
var (
	Grant_Kind             = "Grant"
	Grant_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Grant_Kind}.String()
	Grant_KindAPIVersion   = Grant_Kind + "." + CRDGroupVersion.String()
	Grant_GroupVersionKind = CRDGroupVersion.WithKind(Grant_Kind)
)

func init() {
	SchemeBuilder.Register(&Grant{}, &GrantList{})
}
