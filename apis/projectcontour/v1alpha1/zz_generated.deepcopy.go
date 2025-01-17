//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright Project Contour Authors

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
	"github.com/projectcontour/contour/apis/projectcontour/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContourConfiguration) DeepCopyInto(out *ContourConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContourConfiguration.
func (in *ContourConfiguration) DeepCopy() *ContourConfiguration {
	if in == nil {
		return nil
	}
	out := new(ContourConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContourConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContourConfigurationList) DeepCopyInto(out *ContourConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContourConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContourConfigurationList.
func (in *ContourConfigurationList) DeepCopy() *ContourConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ContourConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContourConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContourConfigurationSpec) DeepCopyInto(out *ContourConfigurationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContourConfigurationSpec.
func (in *ContourConfigurationSpec) DeepCopy() *ContourConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ContourConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContourConfigurationStatus) DeepCopyInto(out *ContourConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.DetailedCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContourConfigurationStatus.
func (in *ContourConfigurationStatus) DeepCopy() *ContourConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ContourConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContourDeployment) DeepCopyInto(out *ContourDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContourDeployment.
func (in *ContourDeployment) DeepCopy() *ContourDeployment {
	if in == nil {
		return nil
	}
	out := new(ContourDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContourDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContourDeploymentList) DeepCopyInto(out *ContourDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContourDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContourDeploymentList.
func (in *ContourDeploymentList) DeepCopy() *ContourDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ContourDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContourDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContourDeploymentSpec) DeepCopyInto(out *ContourDeploymentSpec) {
	*out = *in
	out.Config = in.Config
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContourDeploymentSpec.
func (in *ContourDeploymentSpec) DeepCopy() *ContourDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ContourDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContourDeploymentStatus) DeepCopyInto(out *ContourDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.DetailedCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContourDeploymentStatus.
func (in *ContourDeploymentStatus) DeepCopy() *ContourDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ContourDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionService) DeepCopyInto(out *ExtensionService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionService.
func (in *ExtensionService) DeepCopy() *ExtensionService {
	if in == nil {
		return nil
	}
	out := new(ExtensionService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtensionService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionServiceList) DeepCopyInto(out *ExtensionServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExtensionService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionServiceList.
func (in *ExtensionServiceList) DeepCopy() *ExtensionServiceList {
	if in == nil {
		return nil
	}
	out := new(ExtensionServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtensionServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionServiceSpec) DeepCopyInto(out *ExtensionServiceSpec) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ExtensionServiceTarget, len(*in))
		copy(*out, *in)
	}
	if in.UpstreamValidation != nil {
		in, out := &in.UpstreamValidation, &out.UpstreamValidation
		*out = new(v1.UpstreamValidation)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerPolicy != nil {
		in, out := &in.LoadBalancerPolicy, &out.LoadBalancerPolicy
		*out = new(v1.LoadBalancerPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeoutPolicy != nil {
		in, out := &in.TimeoutPolicy, &out.TimeoutPolicy
		*out = new(v1.TimeoutPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionServiceSpec.
func (in *ExtensionServiceSpec) DeepCopy() *ExtensionServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ExtensionServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionServiceStatus) DeepCopyInto(out *ExtensionServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.DetailedCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionServiceStatus.
func (in *ExtensionServiceStatus) DeepCopy() *ExtensionServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ExtensionServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionServiceTarget) DeepCopyInto(out *ExtensionServiceTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionServiceTarget.
func (in *ExtensionServiceTarget) DeepCopy() *ExtensionServiceTarget {
	if in == nil {
		return nil
	}
	out := new(ExtensionServiceTarget)
	in.DeepCopyInto(out)
	return out
}
