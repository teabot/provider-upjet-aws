// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *CatalogTable) ResolveReferences( // ResolveReferences of this CatalogTable.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "CatalogDatabase", "CatalogDatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DatabaseNameRef,
			Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Connection.
func (mg *Connection) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.PhysicalConnectionRequirements != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PhysicalConnectionRequirements.AvailabilityZone),
				Extract:      resource.ExtractParamPath("availability_zone", false),
				Reference:    mg.Spec.ForProvider.PhysicalConnectionRequirements.AvailabilityZoneRef,
				Selector:     mg.Spec.ForProvider.PhysicalConnectionRequirements.AvailabilityZoneSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PhysicalConnectionRequirements.AvailabilityZone")
		}
		mg.Spec.ForProvider.PhysicalConnectionRequirements.AvailabilityZone = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PhysicalConnectionRequirements.AvailabilityZoneRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.PhysicalConnectionRequirements != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PhysicalConnectionRequirements.SubnetID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.PhysicalConnectionRequirements.SubnetIDRef,
				Selector:     mg.Spec.ForProvider.PhysicalConnectionRequirements.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PhysicalConnectionRequirements.SubnetID")
		}
		mg.Spec.ForProvider.PhysicalConnectionRequirements.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PhysicalConnectionRequirements.SubnetIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.PhysicalConnectionRequirements != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PhysicalConnectionRequirements.AvailabilityZone),
				Extract:      resource.ExtractParamPath("availability_zone", false),
				Reference:    mg.Spec.InitProvider.PhysicalConnectionRequirements.AvailabilityZoneRef,
				Selector:     mg.Spec.InitProvider.PhysicalConnectionRequirements.AvailabilityZoneSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PhysicalConnectionRequirements.AvailabilityZone")
		}
		mg.Spec.InitProvider.PhysicalConnectionRequirements.AvailabilityZone = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PhysicalConnectionRequirements.AvailabilityZoneRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.PhysicalConnectionRequirements != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PhysicalConnectionRequirements.SubnetID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.PhysicalConnectionRequirements.SubnetIDRef,
				Selector:     mg.Spec.InitProvider.PhysicalConnectionRequirements.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PhysicalConnectionRequirements.SubnetID")
		}
		mg.Spec.InitProvider.PhysicalConnectionRequirements.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PhysicalConnectionRequirements.SubnetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Crawler.
func (mg *Crawler) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CatalogTarget); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "CatalogDatabase", "CatalogDatabaseList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CatalogTarget[i3].DatabaseName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.CatalogTarget[i3].DatabaseNameRef,
				Selector:     mg.Spec.ForProvider.CatalogTarget[i3].DatabaseNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CatalogTarget[i3].DatabaseName")
		}
		mg.Spec.ForProvider.CatalogTarget[i3].DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CatalogTarget[i3].DatabaseNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "CatalogDatabase", "CatalogDatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DatabaseNameRef,
			Selector:     mg.Spec.ForProvider.DatabaseNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseName")
	}
	mg.Spec.ForProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.JdbcTarget); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Connection", "ConnectionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.JdbcTarget[i3].ConnectionName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.JdbcTarget[i3].ConnectionNameRef,
				Selector:     mg.Spec.ForProvider.JdbcTarget[i3].ConnectionNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.JdbcTarget[i3].ConnectionName")
		}
		mg.Spec.ForProvider.JdbcTarget[i3].ConnectionName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.JdbcTarget[i3].ConnectionNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.MongodbTarget); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Connection", "ConnectionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MongodbTarget[i3].ConnectionName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.MongodbTarget[i3].ConnectionNameRef,
				Selector:     mg.Spec.ForProvider.MongodbTarget[i3].ConnectionNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.MongodbTarget[i3].ConnectionName")
		}
		mg.Spec.ForProvider.MongodbTarget[i3].ConnectionName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.MongodbTarget[i3].ConnectionNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleRef,
			Selector:     mg.Spec.ForProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.CatalogTarget); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "CatalogDatabase", "CatalogDatabaseList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CatalogTarget[i3].DatabaseName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.CatalogTarget[i3].DatabaseNameRef,
				Selector:     mg.Spec.InitProvider.CatalogTarget[i3].DatabaseNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.CatalogTarget[i3].DatabaseName")
		}
		mg.Spec.InitProvider.CatalogTarget[i3].DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.CatalogTarget[i3].DatabaseNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "CatalogDatabase", "CatalogDatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DatabaseName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DatabaseNameRef,
			Selector:     mg.Spec.InitProvider.DatabaseNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DatabaseName")
	}
	mg.Spec.InitProvider.DatabaseName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatabaseNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.JdbcTarget); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Connection", "ConnectionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.JdbcTarget[i3].ConnectionName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.JdbcTarget[i3].ConnectionNameRef,
				Selector:     mg.Spec.InitProvider.JdbcTarget[i3].ConnectionNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.JdbcTarget[i3].ConnectionName")
		}
		mg.Spec.InitProvider.JdbcTarget[i3].ConnectionName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.JdbcTarget[i3].ConnectionNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.MongodbTarget); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Connection", "ConnectionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MongodbTarget[i3].ConnectionName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.MongodbTarget[i3].ConnectionNameRef,
				Selector:     mg.Spec.InitProvider.MongodbTarget[i3].ConnectionNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.MongodbTarget[i3].ConnectionName")
		}
		mg.Spec.InitProvider.MongodbTarget[i3].ConnectionName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.MongodbTarget[i3].ConnectionNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleRef,
			Selector:     mg.Spec.InitProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DataCatalogEncryptionSettings.
func (mg *DataCatalogEncryptionSettings) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.DataCatalogEncryptionSettings != nil {
		if mg.Spec.ForProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyID),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.ForProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyIDRef,
					Selector:     mg.Spec.ForProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyID")
			}
			mg.Spec.ForProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyIDRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.DataCatalogEncryptionSettings != nil {
		if mg.Spec.ForProvider.DataCatalogEncryptionSettings.EncryptionAtRest != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyID),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.ForProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyIDRef,
					Selector:     mg.Spec.ForProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyID")
			}
			mg.Spec.ForProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyIDRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.DataCatalogEncryptionSettings != nil {
		if mg.Spec.InitProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyID),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.InitProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyIDRef,
					Selector:     mg.Spec.InitProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyID")
			}
			mg.Spec.InitProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.DataCatalogEncryptionSettings.ConnectionPasswordEncryption.AwsKMSKeyIDRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.DataCatalogEncryptionSettings != nil {
		if mg.Spec.InitProvider.DataCatalogEncryptionSettings.EncryptionAtRest != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyID),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.InitProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyIDRef,
					Selector:     mg.Spec.InitProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyID")
			}
			mg.Spec.InitProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.DataCatalogEncryptionSettings.EncryptionAtRest.SseAwsKMSKeyIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Job.
func (mg *Job) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleArnRef,
			Selector:     mg.Spec.ForProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleArnRef,
			Selector:     mg.Spec.InitProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityConfiguration.
func (mg *SecurityConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.EncryptionConfiguration != nil {
		if mg.Spec.ForProvider.EncryptionConfiguration.CloudwatchEncryption != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArn),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.ForProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArnRef,
					Selector:     mg.Spec.ForProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArn")
			}
			mg.Spec.ForProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.EncryptionConfiguration != nil {
		if mg.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArn),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArnRef,
					Selector:     mg.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArn")
			}
			mg.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.EncryptionConfiguration != nil {
		if mg.Spec.ForProvider.EncryptionConfiguration.S3Encryption != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionConfiguration.S3Encryption.KMSKeyArn),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.ForProvider.EncryptionConfiguration.S3Encryption.KMSKeyArnRef,
					Selector:     mg.Spec.ForProvider.EncryptionConfiguration.S3Encryption.KMSKeyArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionConfiguration.S3Encryption.KMSKeyArn")
			}
			mg.Spec.ForProvider.EncryptionConfiguration.S3Encryption.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EncryptionConfiguration.S3Encryption.KMSKeyArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.EncryptionConfiguration != nil {
		if mg.Spec.InitProvider.EncryptionConfiguration.CloudwatchEncryption != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArn),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.InitProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArnRef,
					Selector:     mg.Spec.InitProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArn")
			}
			mg.Spec.InitProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EncryptionConfiguration.CloudwatchEncryption.KMSKeyArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.EncryptionConfiguration != nil {
		if mg.Spec.InitProvider.EncryptionConfiguration.JobBookmarksEncryption != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArn),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.InitProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArnRef,
					Selector:     mg.Spec.InitProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArn")
			}
			mg.Spec.InitProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.EncryptionConfiguration != nil {
		if mg.Spec.InitProvider.EncryptionConfiguration.S3Encryption != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EncryptionConfiguration.S3Encryption.KMSKeyArn),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.InitProvider.EncryptionConfiguration.S3Encryption.KMSKeyArnRef,
					Selector:     mg.Spec.InitProvider.EncryptionConfiguration.S3Encryption.KMSKeyArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EncryptionConfiguration.S3Encryption.KMSKeyArn")
			}
			mg.Spec.InitProvider.EncryptionConfiguration.S3Encryption.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EncryptionConfiguration.S3Encryption.KMSKeyArnRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Trigger.
func (mg *Trigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Actions); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Crawler", "CrawlerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Actions[i3].CrawlerName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Actions[i3].CrawlerNameRef,
				Selector:     mg.Spec.ForProvider.Actions[i3].CrawlerNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Actions[i3].CrawlerName")
		}
		mg.Spec.ForProvider.Actions[i3].CrawlerName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Actions[i3].CrawlerNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Actions); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Job", "JobList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Actions[i3].JobName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Actions[i3].JobNameRef,
				Selector:     mg.Spec.ForProvider.Actions[i3].JobNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Actions[i3].JobName")
		}
		mg.Spec.ForProvider.Actions[i3].JobName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Actions[i3].JobNameRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.Predicate != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Predicate.Conditions); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Crawler", "CrawlerList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Predicate.Conditions[i4].CrawlerName),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.Predicate.Conditions[i4].CrawlerNameRef,
					Selector:     mg.Spec.ForProvider.Predicate.Conditions[i4].CrawlerNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Predicate.Conditions[i4].CrawlerName")
			}
			mg.Spec.ForProvider.Predicate.Conditions[i4].CrawlerName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Predicate.Conditions[i4].CrawlerNameRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.Predicate != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Predicate.Conditions); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Job", "JobList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Predicate.Conditions[i4].JobName),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.Predicate.Conditions[i4].JobNameRef,
					Selector:     mg.Spec.ForProvider.Predicate.Conditions[i4].JobNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Predicate.Conditions[i4].JobName")
			}
			mg.Spec.ForProvider.Predicate.Conditions[i4].JobName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Predicate.Conditions[i4].JobNameRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Actions); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Crawler", "CrawlerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Actions[i3].CrawlerName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.Actions[i3].CrawlerNameRef,
				Selector:     mg.Spec.InitProvider.Actions[i3].CrawlerNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Actions[i3].CrawlerName")
		}
		mg.Spec.InitProvider.Actions[i3].CrawlerName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Actions[i3].CrawlerNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Actions); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Job", "JobList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Actions[i3].JobName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.Actions[i3].JobNameRef,
				Selector:     mg.Spec.InitProvider.Actions[i3].JobNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Actions[i3].JobName")
		}
		mg.Spec.InitProvider.Actions[i3].JobName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Actions[i3].JobNameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.Predicate != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Predicate.Conditions); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Crawler", "CrawlerList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Predicate.Conditions[i4].CrawlerName),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.Predicate.Conditions[i4].CrawlerNameRef,
					Selector:     mg.Spec.InitProvider.Predicate.Conditions[i4].CrawlerNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Predicate.Conditions[i4].CrawlerName")
			}
			mg.Spec.InitProvider.Predicate.Conditions[i4].CrawlerName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Predicate.Conditions[i4].CrawlerNameRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.Predicate != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Predicate.Conditions); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("glue.aws.upbound.io", "v1beta2", "Job", "JobList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Predicate.Conditions[i4].JobName),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.Predicate.Conditions[i4].JobNameRef,
					Selector:     mg.Spec.InitProvider.Predicate.Conditions[i4].JobNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Predicate.Conditions[i4].JobName")
			}
			mg.Spec.InitProvider.Predicate.Conditions[i4].JobName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Predicate.Conditions[i4].JobNameRef = rsp.ResolvedReference

		}
	}

	return nil
}