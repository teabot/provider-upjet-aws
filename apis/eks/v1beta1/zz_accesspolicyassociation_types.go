// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessPolicyAssociationInitParameters struct {
	AccessScope *AccessScopeInitParameters `json:"accessScope,omitempty" tf:"access_scope,omitempty"`
}

type AccessPolicyAssociationObservation struct {
	AccessScope *AccessScopeObservation `json:"accessScope,omitempty" tf:"access_scope,omitempty"`

	AssociatedAt *string `json:"associatedAt,omitempty" tf:"associated_at,omitempty"`

	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ModifiedAt *string `json:"modifiedAt,omitempty" tf:"modified_at,omitempty"`

	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`
}

type AccessPolicyAssociationParameters struct {

	// +kubebuilder:validation:Optional
	AccessScope *AccessScopeParameters `json:"accessScope,omitempty" tf:"access_scope,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/eks/v1beta2.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in eks to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in eks to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	PolicyArn *string `json:"policyArn" tf:"policy_arn,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/eks/v1beta1.AccessEntry
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("principal_arn",true)
	// +kubebuilder:validation:Optional
	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`

	// Reference to a AccessEntry in eks to populate principalArn.
	// +kubebuilder:validation:Optional
	PrincipalArnRef *v1.Reference `json:"principalArnRef,omitempty" tf:"-"`

	// Selector for a AccessEntry in eks to populate principalArn.
	// +kubebuilder:validation:Optional
	PrincipalArnSelector *v1.Selector `json:"principalArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type AccessScopeInitParameters struct {

	// +listType=set
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AccessScopeObservation struct {

	// +listType=set
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AccessScopeParameters struct {

	// +kubebuilder:validation:Optional
	// +listType=set
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// AccessPolicyAssociationSpec defines the desired state of AccessPolicyAssociation
type AccessPolicyAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessPolicyAssociationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AccessPolicyAssociationInitParameters `json:"initProvider,omitempty"`
}

// AccessPolicyAssociationStatus defines the observed state of AccessPolicyAssociation.
type AccessPolicyAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessPolicyAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessPolicyAssociation is the Schema for the AccessPolicyAssociations API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AccessPolicyAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessScope) || (has(self.initProvider) && has(self.initProvider.accessScope))",message="spec.forProvider.accessScope is a required parameter"
	Spec   AccessPolicyAssociationSpec   `json:"spec"`
	Status AccessPolicyAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPolicyAssociationList contains a list of AccessPolicyAssociations
type AccessPolicyAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPolicyAssociation `json:"items"`
}

// Repository type metadata.
var (
	AccessPolicyAssociation_Kind             = "AccessPolicyAssociation"
	AccessPolicyAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPolicyAssociation_Kind}.String()
	AccessPolicyAssociation_KindAPIVersion   = AccessPolicyAssociation_Kind + "." + CRDGroupVersion.String()
	AccessPolicyAssociation_GroupVersionKind = CRDGroupVersion.WithKind(AccessPolicyAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessPolicyAssociation{}, &AccessPolicyAssociationList{})
}