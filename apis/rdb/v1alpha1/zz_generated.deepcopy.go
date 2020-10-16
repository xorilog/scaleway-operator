// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBACL) DeepCopyInto(out *RDBACL) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]RDBACLRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBACL.
func (in *RDBACL) DeepCopy() *RDBACL {
	if in == nil {
		return nil
	}
	out := new(RDBACL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBACLRule) DeepCopyInto(out *RDBACLRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBACLRule.
func (in *RDBACLRule) DeepCopy() *RDBACLRule {
	if in == nil {
		return nil
	}
	out := new(RDBACLRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBDatabase) DeepCopyInto(out *RDBDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBDatabase.
func (in *RDBDatabase) DeepCopy() *RDBDatabase {
	if in == nil {
		return nil
	}
	out := new(RDBDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDBDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBDatabaseList) DeepCopyInto(out *RDBDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RDBDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBDatabaseList.
func (in *RDBDatabaseList) DeepCopy() *RDBDatabaseList {
	if in == nil {
		return nil
	}
	out := new(RDBDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDBDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBDatabaseSpec) DeepCopyInto(out *RDBDatabaseSpec) {
	*out = *in
	out.InstanceRef = in.InstanceRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBDatabaseSpec.
func (in *RDBDatabaseSpec) DeepCopy() *RDBDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(RDBDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBDatabaseStatus) DeepCopyInto(out *RDBDatabaseStatus) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBDatabaseStatus.
func (in *RDBDatabaseStatus) DeepCopy() *RDBDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(RDBDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBInstance) DeepCopyInto(out *RDBInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBInstance.
func (in *RDBInstance) DeepCopy() *RDBInstance {
	if in == nil {
		return nil
	}
	out := new(RDBInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDBInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBInstanceAutoBackup) DeepCopyInto(out *RDBInstanceAutoBackup) {
	*out = *in
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int32)
		**out = **in
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBInstanceAutoBackup.
func (in *RDBInstanceAutoBackup) DeepCopy() *RDBInstanceAutoBackup {
	if in == nil {
		return nil
	}
	out := new(RDBInstanceAutoBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBInstanceEndpoint) DeepCopyInto(out *RDBInstanceEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBInstanceEndpoint.
func (in *RDBInstanceEndpoint) DeepCopy() *RDBInstanceEndpoint {
	if in == nil {
		return nil
	}
	out := new(RDBInstanceEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBInstanceList) DeepCopyInto(out *RDBInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RDBInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBInstanceList.
func (in *RDBInstanceList) DeepCopy() *RDBInstanceList {
	if in == nil {
		return nil
	}
	out := new(RDBInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDBInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBInstancePassword) DeepCopyInto(out *RDBInstancePassword) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(RDBInstancePasswordValueFrom)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBInstancePassword.
func (in *RDBInstancePassword) DeepCopy() *RDBInstancePassword {
	if in == nil {
		return nil
	}
	out := new(RDBInstancePassword)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBInstancePasswordValueFrom) DeepCopyInto(out *RDBInstancePasswordValueFrom) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBInstancePasswordValueFrom.
func (in *RDBInstancePasswordValueFrom) DeepCopy() *RDBInstancePasswordValueFrom {
	if in == nil {
		return nil
	}
	out := new(RDBInstancePasswordValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBInstanceRef) DeepCopyInto(out *RDBInstanceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBInstanceRef.
func (in *RDBInstanceRef) DeepCopy() *RDBInstanceRef {
	if in == nil {
		return nil
	}
	out := new(RDBInstanceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBInstanceSpec) DeepCopyInto(out *RDBInstanceSpec) {
	*out = *in
	if in.InstanceFrom != nil {
		in, out := &in.InstanceFrom, &out.InstanceFrom
		*out = new(RDBInstanceRef)
		**out = **in
	}
	if in.AutoBackup != nil {
		in, out := &in.AutoBackup, &out.AutoBackup
		*out = new(RDBInstanceAutoBackup)
		(*in).DeepCopyInto(*out)
	}
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(RDBACL)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBInstanceSpec.
func (in *RDBInstanceSpec) DeepCopy() *RDBInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(RDBInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBInstanceStatus) DeepCopyInto(out *RDBInstanceStatus) {
	*out = *in
	out.Endpoint = in.Endpoint
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBInstanceStatus.
func (in *RDBInstanceStatus) DeepCopy() *RDBInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(RDBInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBPrivilege) DeepCopyInto(out *RDBPrivilege) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBPrivilege.
func (in *RDBPrivilege) DeepCopy() *RDBPrivilege {
	if in == nil {
		return nil
	}
	out := new(RDBPrivilege)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBUser) DeepCopyInto(out *RDBUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBUser.
func (in *RDBUser) DeepCopy() *RDBUser {
	if in == nil {
		return nil
	}
	out := new(RDBUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDBUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBUserList) DeepCopyInto(out *RDBUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RDBUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBUserList.
func (in *RDBUserList) DeepCopy() *RDBUserList {
	if in == nil {
		return nil
	}
	out := new(RDBUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDBUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBUserSpec) DeepCopyInto(out *RDBUserSpec) {
	*out = *in
	in.Password.DeepCopyInto(&out.Password)
	if in.Privileges != nil {
		in, out := &in.Privileges, &out.Privileges
		*out = make([]RDBPrivilege, len(*in))
		copy(*out, *in)
	}
	out.InstanceRef = in.InstanceRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBUserSpec.
func (in *RDBUserSpec) DeepCopy() *RDBUserSpec {
	if in == nil {
		return nil
	}
	out := new(RDBUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDBUserStatus) DeepCopyInto(out *RDBUserStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDBUserStatus.
func (in *RDBUserStatus) DeepCopy() *RDBUserStatus {
	if in == nil {
		return nil
	}
	out := new(RDBUserStatus)
	in.DeepCopyInto(out)
	return out
}