//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainInitParameters) DeepCopyInto(out *DomainInitParameters) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.EncryptionKey != nil {
		in, out := &in.EncryptionKey, &out.EncryptionKey
		*out = new(string)
		**out = **in
	}
	if in.EncryptionKeyRef != nil {
		in, out := &in.EncryptionKeyRef, &out.EncryptionKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.EncryptionKeySelector != nil {
		in, out := &in.EncryptionKeySelector, &out.EncryptionKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainInitParameters.
func (in *DomainInitParameters) DeepCopy() *DomainInitParameters {
	if in == nil {
		return nil
	}
	out := new(DomainInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainObservation) DeepCopyInto(out *DomainObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AssetSizeBytes != nil {
		in, out := &in.AssetSizeBytes, &out.AssetSizeBytes
		*out = new(string)
		**out = **in
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.EncryptionKey != nil {
		in, out := &in.EncryptionKey, &out.EncryptionKey
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.RepositoryCount != nil {
		in, out := &in.RepositoryCount, &out.RepositoryCount
		*out = new(float64)
		**out = **in
	}
	if in.S3BucketArn != nil {
		in, out := &in.S3BucketArn, &out.S3BucketArn
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainObservation.
func (in *DomainObservation) DeepCopy() *DomainObservation {
	if in == nil {
		return nil
	}
	out := new(DomainObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainParameters) DeepCopyInto(out *DomainParameters) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.EncryptionKey != nil {
		in, out := &in.EncryptionKey, &out.EncryptionKey
		*out = new(string)
		**out = **in
	}
	if in.EncryptionKeyRef != nil {
		in, out := &in.EncryptionKeyRef, &out.EncryptionKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.EncryptionKeySelector != nil {
		in, out := &in.EncryptionKeySelector, &out.EncryptionKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainParameters.
func (in *DomainParameters) DeepCopy() *DomainParameters {
	if in == nil {
		return nil
	}
	out := new(DomainParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPermissionsPolicy) DeepCopyInto(out *DomainPermissionsPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPermissionsPolicy.
func (in *DomainPermissionsPolicy) DeepCopy() *DomainPermissionsPolicy {
	if in == nil {
		return nil
	}
	out := new(DomainPermissionsPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainPermissionsPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPermissionsPolicyInitParameters) DeepCopyInto(out *DomainPermissionsPolicyInitParameters) {
	*out = *in
	if in.DomainOwner != nil {
		in, out := &in.DomainOwner, &out.DomainOwner
		*out = new(string)
		**out = **in
	}
	if in.PolicyDocument != nil {
		in, out := &in.PolicyDocument, &out.PolicyDocument
		*out = new(string)
		**out = **in
	}
	if in.PolicyRevision != nil {
		in, out := &in.PolicyRevision, &out.PolicyRevision
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPermissionsPolicyInitParameters.
func (in *DomainPermissionsPolicyInitParameters) DeepCopy() *DomainPermissionsPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(DomainPermissionsPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPermissionsPolicyList) DeepCopyInto(out *DomainPermissionsPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainPermissionsPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPermissionsPolicyList.
func (in *DomainPermissionsPolicyList) DeepCopy() *DomainPermissionsPolicyList {
	if in == nil {
		return nil
	}
	out := new(DomainPermissionsPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainPermissionsPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPermissionsPolicyObservation) DeepCopyInto(out *DomainPermissionsPolicyObservation) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.DomainOwner != nil {
		in, out := &in.DomainOwner, &out.DomainOwner
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PolicyDocument != nil {
		in, out := &in.PolicyDocument, &out.PolicyDocument
		*out = new(string)
		**out = **in
	}
	if in.PolicyRevision != nil {
		in, out := &in.PolicyRevision, &out.PolicyRevision
		*out = new(string)
		**out = **in
	}
	if in.ResourceArn != nil {
		in, out := &in.ResourceArn, &out.ResourceArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPermissionsPolicyObservation.
func (in *DomainPermissionsPolicyObservation) DeepCopy() *DomainPermissionsPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(DomainPermissionsPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPermissionsPolicyParameters) DeepCopyInto(out *DomainPermissionsPolicyParameters) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.DomainOwner != nil {
		in, out := &in.DomainOwner, &out.DomainOwner
		*out = new(string)
		**out = **in
	}
	if in.DomainRef != nil {
		in, out := &in.DomainRef, &out.DomainRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DomainSelector != nil {
		in, out := &in.DomainSelector, &out.DomainSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PolicyDocument != nil {
		in, out := &in.PolicyDocument, &out.PolicyDocument
		*out = new(string)
		**out = **in
	}
	if in.PolicyRevision != nil {
		in, out := &in.PolicyRevision, &out.PolicyRevision
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPermissionsPolicyParameters.
func (in *DomainPermissionsPolicyParameters) DeepCopy() *DomainPermissionsPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(DomainPermissionsPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPermissionsPolicySpec) DeepCopyInto(out *DomainPermissionsPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPermissionsPolicySpec.
func (in *DomainPermissionsPolicySpec) DeepCopy() *DomainPermissionsPolicySpec {
	if in == nil {
		return nil
	}
	out := new(DomainPermissionsPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPermissionsPolicyStatus) DeepCopyInto(out *DomainPermissionsPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPermissionsPolicyStatus.
func (in *DomainPermissionsPolicyStatus) DeepCopy() *DomainPermissionsPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DomainPermissionsPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainStatus) DeepCopyInto(out *DomainStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainStatus.
func (in *DomainStatus) DeepCopy() *DomainStatus {
	if in == nil {
		return nil
	}
	out := new(DomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalConnectionsInitParameters) DeepCopyInto(out *ExternalConnectionsInitParameters) {
	*out = *in
	if in.ExternalConnectionName != nil {
		in, out := &in.ExternalConnectionName, &out.ExternalConnectionName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalConnectionsInitParameters.
func (in *ExternalConnectionsInitParameters) DeepCopy() *ExternalConnectionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ExternalConnectionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalConnectionsObservation) DeepCopyInto(out *ExternalConnectionsObservation) {
	*out = *in
	if in.ExternalConnectionName != nil {
		in, out := &in.ExternalConnectionName, &out.ExternalConnectionName
		*out = new(string)
		**out = **in
	}
	if in.PackageFormat != nil {
		in, out := &in.PackageFormat, &out.PackageFormat
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalConnectionsObservation.
func (in *ExternalConnectionsObservation) DeepCopy() *ExternalConnectionsObservation {
	if in == nil {
		return nil
	}
	out := new(ExternalConnectionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalConnectionsParameters) DeepCopyInto(out *ExternalConnectionsParameters) {
	*out = *in
	if in.ExternalConnectionName != nil {
		in, out := &in.ExternalConnectionName, &out.ExternalConnectionName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalConnectionsParameters.
func (in *ExternalConnectionsParameters) DeepCopy() *ExternalConnectionsParameters {
	if in == nil {
		return nil
	}
	out := new(ExternalConnectionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryInitParameters) DeepCopyInto(out *RepositoryInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DomainOwner != nil {
		in, out := &in.DomainOwner, &out.DomainOwner
		*out = new(string)
		**out = **in
	}
	if in.ExternalConnections != nil {
		in, out := &in.ExternalConnections, &out.ExternalConnections
		*out = new(ExternalConnectionsInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = make([]UpstreamInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryInitParameters.
func (in *RepositoryInitParameters) DeepCopy() *RepositoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryObservation) DeepCopyInto(out *RepositoryObservation) {
	*out = *in
	if in.AdministratorAccount != nil {
		in, out := &in.AdministratorAccount, &out.AdministratorAccount
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.DomainOwner != nil {
		in, out := &in.DomainOwner, &out.DomainOwner
		*out = new(string)
		**out = **in
	}
	if in.ExternalConnections != nil {
		in, out := &in.ExternalConnections, &out.ExternalConnections
		*out = new(ExternalConnectionsObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = make([]UpstreamObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryObservation.
func (in *RepositoryObservation) DeepCopy() *RepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(RepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryParameters) DeepCopyInto(out *RepositoryParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.DomainOwner != nil {
		in, out := &in.DomainOwner, &out.DomainOwner
		*out = new(string)
		**out = **in
	}
	if in.DomainRef != nil {
		in, out := &in.DomainRef, &out.DomainRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DomainSelector != nil {
		in, out := &in.DomainSelector, &out.DomainSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalConnections != nil {
		in, out := &in.ExternalConnections, &out.ExternalConnections
		*out = new(ExternalConnectionsParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = make([]UpstreamParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryParameters.
func (in *RepositoryParameters) DeepCopy() *RepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryPermissionsPolicy) DeepCopyInto(out *RepositoryPermissionsPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryPermissionsPolicy.
func (in *RepositoryPermissionsPolicy) DeepCopy() *RepositoryPermissionsPolicy {
	if in == nil {
		return nil
	}
	out := new(RepositoryPermissionsPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryPermissionsPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryPermissionsPolicyInitParameters) DeepCopyInto(out *RepositoryPermissionsPolicyInitParameters) {
	*out = *in
	if in.DomainOwner != nil {
		in, out := &in.DomainOwner, &out.DomainOwner
		*out = new(string)
		**out = **in
	}
	if in.PolicyDocument != nil {
		in, out := &in.PolicyDocument, &out.PolicyDocument
		*out = new(string)
		**out = **in
	}
	if in.PolicyRevision != nil {
		in, out := &in.PolicyRevision, &out.PolicyRevision
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryPermissionsPolicyInitParameters.
func (in *RepositoryPermissionsPolicyInitParameters) DeepCopy() *RepositoryPermissionsPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryPermissionsPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryPermissionsPolicyList) DeepCopyInto(out *RepositoryPermissionsPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RepositoryPermissionsPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryPermissionsPolicyList.
func (in *RepositoryPermissionsPolicyList) DeepCopy() *RepositoryPermissionsPolicyList {
	if in == nil {
		return nil
	}
	out := new(RepositoryPermissionsPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryPermissionsPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryPermissionsPolicyObservation) DeepCopyInto(out *RepositoryPermissionsPolicyObservation) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.DomainOwner != nil {
		in, out := &in.DomainOwner, &out.DomainOwner
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PolicyDocument != nil {
		in, out := &in.PolicyDocument, &out.PolicyDocument
		*out = new(string)
		**out = **in
	}
	if in.PolicyRevision != nil {
		in, out := &in.PolicyRevision, &out.PolicyRevision
		*out = new(string)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.ResourceArn != nil {
		in, out := &in.ResourceArn, &out.ResourceArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryPermissionsPolicyObservation.
func (in *RepositoryPermissionsPolicyObservation) DeepCopy() *RepositoryPermissionsPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(RepositoryPermissionsPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryPermissionsPolicyParameters) DeepCopyInto(out *RepositoryPermissionsPolicyParameters) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.DomainOwner != nil {
		in, out := &in.DomainOwner, &out.DomainOwner
		*out = new(string)
		**out = **in
	}
	if in.DomainRef != nil {
		in, out := &in.DomainRef, &out.DomainRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DomainSelector != nil {
		in, out := &in.DomainSelector, &out.DomainSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PolicyDocument != nil {
		in, out := &in.PolicyDocument, &out.PolicyDocument
		*out = new(string)
		**out = **in
	}
	if in.PolicyRevision != nil {
		in, out := &in.PolicyRevision, &out.PolicyRevision
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.RepositoryRef != nil {
		in, out := &in.RepositoryRef, &out.RepositoryRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RepositorySelector != nil {
		in, out := &in.RepositorySelector, &out.RepositorySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryPermissionsPolicyParameters.
func (in *RepositoryPermissionsPolicyParameters) DeepCopy() *RepositoryPermissionsPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryPermissionsPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryPermissionsPolicySpec) DeepCopyInto(out *RepositoryPermissionsPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryPermissionsPolicySpec.
func (in *RepositoryPermissionsPolicySpec) DeepCopy() *RepositoryPermissionsPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RepositoryPermissionsPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryPermissionsPolicyStatus) DeepCopyInto(out *RepositoryPermissionsPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryPermissionsPolicyStatus.
func (in *RepositoryPermissionsPolicyStatus) DeepCopy() *RepositoryPermissionsPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryPermissionsPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpec) DeepCopyInto(out *RepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpec.
func (in *RepositorySpec) DeepCopy() *RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryStatus) DeepCopyInto(out *RepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryStatus.
func (in *RepositoryStatus) DeepCopy() *RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamInitParameters) DeepCopyInto(out *UpstreamInitParameters) {
	*out = *in
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RepositoryNameRef != nil {
		in, out := &in.RepositoryNameRef, &out.RepositoryNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RepositoryNameSelector != nil {
		in, out := &in.RepositoryNameSelector, &out.RepositoryNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamInitParameters.
func (in *UpstreamInitParameters) DeepCopy() *UpstreamInitParameters {
	if in == nil {
		return nil
	}
	out := new(UpstreamInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamObservation) DeepCopyInto(out *UpstreamObservation) {
	*out = *in
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamObservation.
func (in *UpstreamObservation) DeepCopy() *UpstreamObservation {
	if in == nil {
		return nil
	}
	out := new(UpstreamObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamParameters) DeepCopyInto(out *UpstreamParameters) {
	*out = *in
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RepositoryNameRef != nil {
		in, out := &in.RepositoryNameRef, &out.RepositoryNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RepositoryNameSelector != nil {
		in, out := &in.RepositoryNameSelector, &out.RepositoryNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamParameters.
func (in *UpstreamParameters) DeepCopy() *UpstreamParameters {
	if in == nil {
		return nil
	}
	out := new(UpstreamParameters)
	in.DeepCopyInto(out)
	return out
}