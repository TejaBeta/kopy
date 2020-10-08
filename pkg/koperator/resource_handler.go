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

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetDeployments returns all the Deployments in the given namespace and clientset
func (kOpts *Options) GetDeployments() (result *appv1.DeploymentList, err error) {
	result, err = kOpts.clientset.
		AppsV1().
		Deployments(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteDeployment is a method to delete a provided deployment name
func (kOpts *Options) DeleteDeployment(name string) (err error) {
	err = kOpts.clientset.
		AppsV1().
		Deployments(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateDeployment method to create a deployment
func (kOpts *Options) CreateDeployment(deployment *appv1.Deployment) (result *appv1.Deployment, err error) {
	result, err = kOpts.clientset.
		AppsV1().
		Deployments(kOpts.namespace).
		Create(context.TODO(), deployment, metav1.CreateOptions{})
	return
}

// GetConfigMaps returns all the Configmaps in the given namespace and clientset
func (kOpts *Options) GetConfigMaps() (result *corev1.ConfigMapList, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		ConfigMaps(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteConfigMap is a method to delete a provided configmap name
func (kOpts *Options) DeleteConfigMap(name string) (err error) {
	err = kOpts.clientset.
		CoreV1().
		ConfigMaps(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateConfigMap is a method to create a configmap
func (kOpts *Options) CreateConfigMap(configmap *corev1.ConfigMap) (result *corev1.ConfigMap, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		ConfigMaps(kOpts.namespace).
		Create(context.TODO(), configmap, metav1.CreateOptions{})
	return
}

// GetIngress returns all the Ingresses in the given namespace and clientset
func (kOpts *Options) GetIngress() (result *v1beta1.IngressList, err error) {
	result, err = kOpts.clientset.
		ExtensionsV1beta1().
		Ingresses(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteIngress method deletes an ingress with the given name
func (kOpts *Options) DeleteIngress(name string) (err error) {
	err = kOpts.clientset.
		ExtensionsV1beta1().
		Ingresses(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateIngress method to create an ingress
func (kOpts *Options) CreateIngress(ingress *v1beta1.Ingress) (result *v1beta1.Ingress, err error) {
	result, err = kOpts.clientset.
		ExtensionsV1beta1().
		Ingresses(kOpts.namespace).
		Create(context.TODO(), ingress, metav1.CreateOptions{})
	return
}

// GetNS validates if the namespace exists or not
func (kOpts *Options) GetNS() (result *corev1.Namespace, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		Namespaces().
		Get(context.TODO(), kOpts.namespace, metav1.GetOptions{})
	return
}

// DeleteNS method to delete a namespace
func (kOpts *Options) DeleteNS(name string) (err error) {
	err = kOpts.clientset.
		CoreV1().
		Namespaces().
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateNS method to delete a namespace
func (kOpts *Options) CreateNS(namespace *v1.Namespace) (result *corev1.Namespace, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		Namespaces().
		Create(context.TODO(), namespace, metav1.CreateOptions{})
	return
}

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

// GetSecrets returns all the Secrets in the given namespace and clientset
func (kOpts *Options) GetSecrets() (result *corev1.SecretList, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		Secrets(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteSecret method deletes a secret with the name
func (kOpts *Options) DeleteSecret(name string) (err error) {
	err = kOpts.clientset.
		CoreV1().
		Secrets(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateSecret method creates a secret
func (kOpts *Options) CreateSecret(secret *corev1.Secret) (result *corev1.Secret, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		Secrets(kOpts.namespace).
		Create(context.TODO(), secret, metav1.CreateOptions{})
	return
}

// GetSVC returns all the Services in the given namespace and clientset
func (kOpts *Options) GetSVC() (result *corev1.ServiceList, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		Services(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteSVC method to delete a svc with the name
func (kOpts *Options) DeleteSVC(name string) (err error) {
	err = kOpts.clientset.
		CoreV1().
		Services(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateSVC method to create a svc
func (kOpts *Options) CreateSVC(service *corev1.Service) (result *corev1.Service, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		Services(kOpts.namespace).
		Create(context.TODO(), service, metav1.CreateOptions{})
	return
}
