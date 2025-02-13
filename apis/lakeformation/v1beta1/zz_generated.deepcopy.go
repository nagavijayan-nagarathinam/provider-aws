//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateDatabaseDefaultPermissionsObservation) DeepCopyInto(out *CreateDatabaseDefaultPermissionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateDatabaseDefaultPermissionsObservation.
func (in *CreateDatabaseDefaultPermissionsObservation) DeepCopy() *CreateDatabaseDefaultPermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(CreateDatabaseDefaultPermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateDatabaseDefaultPermissionsParameters) DeepCopyInto(out *CreateDatabaseDefaultPermissionsParameters) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateDatabaseDefaultPermissionsParameters.
func (in *CreateDatabaseDefaultPermissionsParameters) DeepCopy() *CreateDatabaseDefaultPermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(CreateDatabaseDefaultPermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateTableDefaultPermissionsObservation) DeepCopyInto(out *CreateTableDefaultPermissionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateTableDefaultPermissionsObservation.
func (in *CreateTableDefaultPermissionsObservation) DeepCopy() *CreateTableDefaultPermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(CreateTableDefaultPermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateTableDefaultPermissionsParameters) DeepCopyInto(out *CreateTableDefaultPermissionsParameters) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateTableDefaultPermissionsParameters.
func (in *CreateTableDefaultPermissionsParameters) DeepCopy() *CreateTableDefaultPermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(CreateTableDefaultPermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLakeSettings) DeepCopyInto(out *DataLakeSettings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLakeSettings.
func (in *DataLakeSettings) DeepCopy() *DataLakeSettings {
	if in == nil {
		return nil
	}
	out := new(DataLakeSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataLakeSettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLakeSettingsList) DeepCopyInto(out *DataLakeSettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataLakeSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLakeSettingsList.
func (in *DataLakeSettingsList) DeepCopy() *DataLakeSettingsList {
	if in == nil {
		return nil
	}
	out := new(DataLakeSettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataLakeSettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLakeSettingsObservation) DeepCopyInto(out *DataLakeSettingsObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLakeSettingsObservation.
func (in *DataLakeSettingsObservation) DeepCopy() *DataLakeSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(DataLakeSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLakeSettingsParameters) DeepCopyInto(out *DataLakeSettingsParameters) {
	*out = *in
	if in.Admins != nil {
		in, out := &in.Admins, &out.Admins
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CatalogID != nil {
		in, out := &in.CatalogID, &out.CatalogID
		*out = new(string)
		**out = **in
	}
	if in.CreateDatabaseDefaultPermissions != nil {
		in, out := &in.CreateDatabaseDefaultPermissions, &out.CreateDatabaseDefaultPermissions
		*out = make([]CreateDatabaseDefaultPermissionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateTableDefaultPermissions != nil {
		in, out := &in.CreateTableDefaultPermissions, &out.CreateTableDefaultPermissions
		*out = make([]CreateTableDefaultPermissionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TrustedResourceOwners != nil {
		in, out := &in.TrustedResourceOwners, &out.TrustedResourceOwners
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLakeSettingsParameters.
func (in *DataLakeSettingsParameters) DeepCopy() *DataLakeSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(DataLakeSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLakeSettingsSpec) DeepCopyInto(out *DataLakeSettingsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLakeSettingsSpec.
func (in *DataLakeSettingsSpec) DeepCopy() *DataLakeSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(DataLakeSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLakeSettingsStatus) DeepCopyInto(out *DataLakeSettingsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLakeSettingsStatus.
func (in *DataLakeSettingsStatus) DeepCopy() *DataLakeSettingsStatus {
	if in == nil {
		return nil
	}
	out := new(DataLakeSettingsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLocationObservation) DeepCopyInto(out *DataLocationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLocationObservation.
func (in *DataLocationObservation) DeepCopy() *DataLocationObservation {
	if in == nil {
		return nil
	}
	out := new(DataLocationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataLocationParameters) DeepCopyInto(out *DataLocationParameters) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ArnRef != nil {
		in, out := &in.ArnRef, &out.ArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ArnSelector != nil {
		in, out := &in.ArnSelector, &out.ArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.CatalogID != nil {
		in, out := &in.CatalogID, &out.CatalogID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataLocationParameters.
func (in *DataLocationParameters) DeepCopy() *DataLocationParameters {
	if in == nil {
		return nil
	}
	out := new(DataLocationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseObservation) DeepCopyInto(out *DatabaseObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseObservation.
func (in *DatabaseObservation) DeepCopy() *DatabaseObservation {
	if in == nil {
		return nil
	}
	out := new(DatabaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseParameters) DeepCopyInto(out *DatabaseParameters) {
	*out = *in
	if in.CatalogID != nil {
		in, out := &in.CatalogID, &out.CatalogID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameRef != nil {
		in, out := &in.NameRef, &out.NameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseParameters.
func (in *DatabaseParameters) DeepCopy() *DatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(DatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressionObservation) DeepCopyInto(out *ExpressionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressionObservation.
func (in *ExpressionObservation) DeepCopy() *ExpressionObservation {
	if in == nil {
		return nil
	}
	out := new(ExpressionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressionParameters) DeepCopyInto(out *ExpressionParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressionParameters.
func (in *ExpressionParameters) DeepCopy() *ExpressionParameters {
	if in == nil {
		return nil
	}
	out := new(ExpressionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LfTagObservation) DeepCopyInto(out *LfTagObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LfTagObservation.
func (in *LfTagObservation) DeepCopy() *LfTagObservation {
	if in == nil {
		return nil
	}
	out := new(LfTagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LfTagParameters) DeepCopyInto(out *LfTagParameters) {
	*out = *in
	if in.CatalogID != nil {
		in, out := &in.CatalogID, &out.CatalogID
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LfTagParameters.
func (in *LfTagParameters) DeepCopy() *LfTagParameters {
	if in == nil {
		return nil
	}
	out := new(LfTagParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LfTagPolicyObservation) DeepCopyInto(out *LfTagPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LfTagPolicyObservation.
func (in *LfTagPolicyObservation) DeepCopy() *LfTagPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(LfTagPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LfTagPolicyParameters) DeepCopyInto(out *LfTagPolicyParameters) {
	*out = *in
	if in.CatalogID != nil {
		in, out := &in.CatalogID, &out.CatalogID
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = make([]ExpressionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LfTagPolicyParameters.
func (in *LfTagPolicyParameters) DeepCopy() *LfTagPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(LfTagPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permissions) DeepCopyInto(out *Permissions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permissions.
func (in *Permissions) DeepCopy() *Permissions {
	if in == nil {
		return nil
	}
	out := new(Permissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Permissions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionsList) DeepCopyInto(out *PermissionsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Permissions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsList.
func (in *PermissionsList) DeepCopy() *PermissionsList {
	if in == nil {
		return nil
	}
	out := new(PermissionsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PermissionsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionsObservation) DeepCopyInto(out *PermissionsObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsObservation.
func (in *PermissionsObservation) DeepCopy() *PermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(PermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionsParameters) DeepCopyInto(out *PermissionsParameters) {
	*out = *in
	if in.CatalogID != nil {
		in, out := &in.CatalogID, &out.CatalogID
		*out = new(string)
		**out = **in
	}
	if in.CatalogResource != nil {
		in, out := &in.CatalogResource, &out.CatalogResource
		*out = new(bool)
		**out = **in
	}
	if in.DataLocation != nil {
		in, out := &in.DataLocation, &out.DataLocation
		*out = make([]DataLocationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = make([]DatabaseParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LfTag != nil {
		in, out := &in.LfTag, &out.LfTag
		*out = make([]LfTagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LfTagPolicy != nil {
		in, out := &in.LfTagPolicy, &out.LfTagPolicy
		*out = make([]LfTagPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PermissionsWithGrantOption != nil {
		in, out := &in.PermissionsWithGrantOption, &out.PermissionsWithGrantOption
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Table != nil {
		in, out := &in.Table, &out.Table
		*out = make([]TableParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TableWithColumns != nil {
		in, out := &in.TableWithColumns, &out.TableWithColumns
		*out = make([]TableWithColumnsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsParameters.
func (in *PermissionsParameters) DeepCopy() *PermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(PermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionsSpec) DeepCopyInto(out *PermissionsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsSpec.
func (in *PermissionsSpec) DeepCopy() *PermissionsSpec {
	if in == nil {
		return nil
	}
	out := new(PermissionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionsStatus) DeepCopyInto(out *PermissionsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsStatus.
func (in *PermissionsStatus) DeepCopy() *PermissionsStatus {
	if in == nil {
		return nil
	}
	out := new(PermissionsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Resource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceList) DeepCopyInto(out *ResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceList.
func (in *ResourceList) DeepCopy() *ResourceList {
	if in == nil {
		return nil
	}
	out := new(ResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceObservation) DeepCopyInto(out *ResourceObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceObservation.
func (in *ResourceObservation) DeepCopy() *ResourceObservation {
	if in == nil {
		return nil
	}
	out := new(ResourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceParameters) DeepCopyInto(out *ResourceParameters) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.RoleArnRef != nil {
		in, out := &in.RoleArnRef, &out.RoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleArnSelector != nil {
		in, out := &in.RoleArnSelector, &out.RoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceParameters.
func (in *ResourceParameters) DeepCopy() *ResourceParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatus) DeepCopyInto(out *ResourceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatus.
func (in *ResourceStatus) DeepCopy() *ResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableObservation) DeepCopyInto(out *TableObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableObservation.
func (in *TableObservation) DeepCopy() *TableObservation {
	if in == nil {
		return nil
	}
	out := new(TableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableParameters) DeepCopyInto(out *TableParameters) {
	*out = *in
	if in.CatalogID != nil {
		in, out := &in.CatalogID, &out.CatalogID
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Wildcard != nil {
		in, out := &in.Wildcard, &out.Wildcard
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableParameters.
func (in *TableParameters) DeepCopy() *TableParameters {
	if in == nil {
		return nil
	}
	out := new(TableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableWithColumnsObservation) DeepCopyInto(out *TableWithColumnsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableWithColumnsObservation.
func (in *TableWithColumnsObservation) DeepCopy() *TableWithColumnsObservation {
	if in == nil {
		return nil
	}
	out := new(TableWithColumnsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableWithColumnsParameters) DeepCopyInto(out *TableWithColumnsParameters) {
	*out = *in
	if in.CatalogID != nil {
		in, out := &in.CatalogID, &out.CatalogID
		*out = new(string)
		**out = **in
	}
	if in.ColumnNames != nil {
		in, out := &in.ColumnNames, &out.ColumnNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.ExcludedColumnNames != nil {
		in, out := &in.ExcludedColumnNames, &out.ExcludedColumnNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameRef != nil {
		in, out := &in.NameRef, &out.NameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Wildcard != nil {
		in, out := &in.Wildcard, &out.Wildcard
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableWithColumnsParameters.
func (in *TableWithColumnsParameters) DeepCopy() *TableWithColumnsParameters {
	if in == nil {
		return nil
	}
	out := new(TableWithColumnsParameters)
	in.DeepCopyInto(out)
	return out
}
