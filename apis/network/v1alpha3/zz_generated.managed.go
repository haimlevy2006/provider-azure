/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha3

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

// GetBindingPhase of this Subnet.
func (mg *Subnet) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this Subnet.
func (mg *Subnet) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this Subnet.
func (mg *Subnet) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this Subnet.
func (mg *Subnet) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetProviderReference of this Subnet.
func (mg *Subnet) GetProviderReference() runtimev1alpha1.Reference {
	return mg.Spec.ProviderReference
}

// GetReclaimPolicy of this Subnet.
func (mg *Subnet) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this Subnet.
func (mg *Subnet) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this Subnet.
func (mg *Subnet) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this Subnet.
func (mg *Subnet) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this Subnet.
func (mg *Subnet) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this Subnet.
func (mg *Subnet) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetProviderReference of this Subnet.
func (mg *Subnet) SetProviderReference(r runtimev1alpha1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetReclaimPolicy of this Subnet.
func (mg *Subnet) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this Subnet.
func (mg *Subnet) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this VirtualNetwork.
func (mg *VirtualNetwork) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this VirtualNetwork.
func (mg *VirtualNetwork) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this VirtualNetwork.
func (mg *VirtualNetwork) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this VirtualNetwork.
func (mg *VirtualNetwork) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetProviderReference of this VirtualNetwork.
func (mg *VirtualNetwork) GetProviderReference() runtimev1alpha1.Reference {
	return mg.Spec.ProviderReference
}

// GetReclaimPolicy of this VirtualNetwork.
func (mg *VirtualNetwork) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this VirtualNetwork.
func (mg *VirtualNetwork) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this VirtualNetwork.
func (mg *VirtualNetwork) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this VirtualNetwork.
func (mg *VirtualNetwork) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this VirtualNetwork.
func (mg *VirtualNetwork) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this VirtualNetwork.
func (mg *VirtualNetwork) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetProviderReference of this VirtualNetwork.
func (mg *VirtualNetwork) SetProviderReference(r runtimev1alpha1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetReclaimPolicy of this VirtualNetwork.
func (mg *VirtualNetwork) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this VirtualNetwork.
func (mg *VirtualNetwork) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}