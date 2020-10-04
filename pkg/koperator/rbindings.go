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

package koperator

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetRoleBindings returns all the RoleBindings in the given namespace and clientset
func (kOpts *Options) GetRoleBindings() (result *rbacv1.RoleBindingList, err error) {
	result, err = kOpts.clientset.
		RbacV1().
		RoleBindings(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteRBinding method deletes a rolebindings with the given name
func (kOpts *Options) DeleteRBinding(name string) (err error) {
	err = kOpts.clientset.
		RbacV1().
		RoleBindings(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateRBinding method creates a rolebinding
func (kOpts *Options) CreateRBinding(rBinding *rbacv1.RoleBinding) (result *rbacv1.RoleBinding, err error) {
	result, err = kOpts.clientset.
		RbacV1().
		RoleBindings(kOpts.namespace).
		Create(context.TODO(), rBinding, metav1.CreateOptions{})
	return
}
