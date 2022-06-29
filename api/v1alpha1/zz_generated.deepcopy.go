//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
func (in *BuilderSpec) DeepCopyInto(out *BuilderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderSpec.
func (in *BuilderSpec) DeepCopy() *BuilderSpec {
	if in == nil {
		return nil
	}
	out := new(BuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebAppSpec) DeepCopyInto(out *WebAppSpec) {
	*out = *in
	if in.Builder != nil {
		in, out := &in.Builder, &out.Builder
		*out = new(BuilderSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebAppSpec.
func (in *WebAppSpec) DeepCopy() *WebAppSpec {
	if in == nil {
		return nil
	}
	out := new(WebAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebImageSpec) DeepCopyInto(out *WebImageSpec) {
	*out = *in
	if in.WebApp != nil {
		in, out := &in.WebApp, &out.WebApp
		*out = new(WebAppSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.WebServerHealthCheck != nil {
		in, out := &in.WebServerHealthCheck, &out.WebServerHealthCheck
		*out = new(WebServerHealthCheckSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebImageSpec.
func (in *WebImageSpec) DeepCopy() *WebImageSpec {
	if in == nil {
		return nil
	}
	out := new(WebImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebImageStreamSpec) DeepCopyInto(out *WebImageStreamSpec) {
	*out = *in
	if in.WebSources != nil {
		in, out := &in.WebSources, &out.WebSources
		*out = new(WebSourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.WebServerHealthCheck != nil {
		in, out := &in.WebServerHealthCheck, &out.WebServerHealthCheck
		*out = new(WebServerHealthCheckSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebImageStreamSpec.
func (in *WebImageStreamSpec) DeepCopy() *WebImageStreamSpec {
	if in == nil {
		return nil
	}
	out := new(WebImageStreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServer) DeepCopyInto(out *WebServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServer.
func (in *WebServer) DeepCopy() *WebServer {
	if in == nil {
		return nil
	}
	out := new(WebServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerHealthCheckSpec) DeepCopyInto(out *WebServerHealthCheckSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerHealthCheckSpec.
func (in *WebServerHealthCheckSpec) DeepCopy() *WebServerHealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(WebServerHealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerList) DeepCopyInto(out *WebServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerList.
func (in *WebServerList) DeepCopy() *WebServerList {
	if in == nil {
		return nil
	}
	out := new(WebServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerSpec) DeepCopyInto(out *WebServerSpec) {
	*out = *in
	if in.WebImage != nil {
		in, out := &in.WebImage, &out.WebImage
		*out = new(WebImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.WebImageStream != nil {
		in, out := &in.WebImageStream, &out.WebImageStream
		*out = new(WebImageStreamSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerSpec.
func (in *WebServerSpec) DeepCopy() *WebServerSpec {
	if in == nil {
		return nil
	}
	out := new(WebServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerStatus) DeepCopyInto(out *WebServerStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]PodStatus, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerStatus.
func (in *WebServerStatus) DeepCopy() *WebServerStatus {
	if in == nil {
		return nil
	}
	out := new(WebServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebSourcesParamsSpec) DeepCopyInto(out *WebSourcesParamsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebSourcesParamsSpec.
func (in *WebSourcesParamsSpec) DeepCopy() *WebSourcesParamsSpec {
	if in == nil {
		return nil
	}
	out := new(WebSourcesParamsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebSourcesSpec) DeepCopyInto(out *WebSourcesSpec) {
	*out = *in
	if in.WebSourcesParams != nil {
		in, out := &in.WebSourcesParams, &out.WebSourcesParams
		*out = new(WebSourcesParamsSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebSourcesSpec.
func (in *WebSourcesSpec) DeepCopy() *WebSourcesSpec {
	if in == nil {
		return nil
	}
	out := new(WebSourcesSpec)
	in.DeepCopyInto(out)
	return out
}
