//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
func (in *SleepCycle) DeepCopyInto(out *SleepCycle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepCycle.
func (in *SleepCycle) DeepCopy() *SleepCycle {
	if in == nil {
		return nil
	}
	out := new(SleepCycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SleepCycle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepCycleList) DeepCopyInto(out *SleepCycleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SleepCycle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepCycleList.
func (in *SleepCycleList) DeepCopy() *SleepCycleList {
	if in == nil {
		return nil
	}
	out := new(SleepCycleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SleepCycleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepCycleSpec) DeepCopyInto(out *SleepCycleSpec) {
	*out = *in
	if in.ShutdownTimeZone != nil {
		in, out := &in.ShutdownTimeZone, &out.ShutdownTimeZone
		*out = new(string)
		**out = **in
	}
	if in.WakeUp != nil {
		in, out := &in.WakeUp, &out.WakeUp
		*out = new(string)
		**out = **in
	}
	if in.WakeupTimeZone != nil {
		in, out := &in.WakeupTimeZone, &out.WakeupTimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepCycleSpec.
func (in *SleepCycleSpec) DeepCopy() *SleepCycleSpec {
	if in == nil {
		return nil
	}
	out := new(SleepCycleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepCycleStatus) DeepCopyInto(out *SleepCycleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepCycleStatus.
func (in *SleepCycleStatus) DeepCopy() *SleepCycleStatus {
	if in == nil {
		return nil
	}
	out := new(SleepCycleStatus)
	in.DeepCopyInto(out)
	return out
}
