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

type RolePolicyAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RolePolicyAttachmentParameters struct {

	// +crossplane:generate:reference:type=Policy
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyArnRef *v1.Reference `json:"policyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PolicyArnSelector *v1.Selector `json:"policyArnSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Role
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`
}

// RolePolicyAttachmentSpec defines the desired state of RolePolicyAttachment
type RolePolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RolePolicyAttachmentParameters `json:"forProvider"`
}

// RolePolicyAttachmentStatus defines the observed state of RolePolicyAttachment.
type RolePolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RolePolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicyAttachment is the Schema for the RolePolicyAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RolePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RolePolicyAttachmentSpec   `json:"spec"`
	Status            RolePolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicyAttachmentList contains a list of RolePolicyAttachments
type RolePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RolePolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	RolePolicyAttachment_Kind             = "RolePolicyAttachment"
	RolePolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RolePolicyAttachment_Kind}.String()
	RolePolicyAttachment_KindAPIVersion   = RolePolicyAttachment_Kind + "." + CRDGroupVersion.String()
	RolePolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(RolePolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&RolePolicyAttachment{}, &RolePolicyAttachmentList{})
}