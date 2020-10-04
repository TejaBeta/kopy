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

// GetRoles returns all the Roles in the given namespace and clientset
func (kOpts *Options) GetRoles() (result *rbacv1.RoleList, err error) {
	result, err = kOpts.clientset.
		RbacV1().
		Roles(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteRole method deletes a role based on the name provided
func (kOpts *Options) DeleteRole(name string) (err error) {
	err = kOpts.clientset.
		RbacV1().
		Roles(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateRole method creates a role
func (kOpts *Options) CreateRole(role *rbacv1.Role) (result *rbacv1.Role, err error) {
	result, err = kOpts.clientset.
		RbacV1().
		Roles(kOpts.namespace).
		Create(context.TODO(), role, metav1.CreateOptions{})
	return
}
