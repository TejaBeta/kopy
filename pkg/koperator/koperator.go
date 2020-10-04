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

	"k8s.io/client-go/rest"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Options to pass to all the methods
type Options struct {
	clientset *kubernetes.Clientset
	namespace string
}

// GetOpts generates required options
func GetOpts(context *rest.Config, ns string) (*Options, error) {
	cs, err := kubernetes.NewForConfig(context)
	if err != nil {
		return nil, err
	}

	return &Options{
		clientset: cs,
		namespace: ns,
	}, nil
}

// IsValidNS validates if the namespace exists or not
func (kOpts *Options) IsValidNS() bool {
	_, err := kOpts.clientset.CoreV1().Namespaces().Get(context.TODO(), kOpts.namespace, metav1.GetOptions{})
	if err != nil {
		return false
	}
	return true
}

// GetDeployments returns all the Deployments in the given namespace and clientset
func (kOpts *Options) GetDeployments() ([]appv1.Deployment, error) {
	deploymentsClient := kOpts.clientset.AppsV1().Deployments(kOpts.namespace)

	deploymentList, getErr := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return deploymentList.Items, nil
}

// GetConfigMaps returns all the Configmaps in the given namespace and clientset
func (kOpts *Options) GetConfigMaps() ([]corev1.ConfigMap, error) {
	configmapsList, getErr := kOpts.clientset.CoreV1().ConfigMaps(kOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return configmapsList.Items, nil
}

// GetClusterRoles returns all the ClusterRoles in the given namespace and clientset
func (kOpts *Options) GetClusterRoles() ([]rbacv1.ClusterRole, error) {
	clusterRoles, getErr := kOpts.clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return clusterRoles.Items, nil
}

// GetRoles returns all the Roles in the given namespace and clientset
func (kOpts *Options) GetRoles() ([]rbacv1.Role, error) {
	roles, getErr := kOpts.clientset.RbacV1().Roles(kOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return roles.Items, nil
}

// GetClusterRoleBindings returns all the ClusterRoleBindings in the given namespace and clientset
func (kOpts *Options) GetClusterRoleBindings() ([]rbacv1.ClusterRoleBinding, error) {
	clusterRoleBindings, getErr := kOpts.clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return clusterRoleBindings.Items, nil
}

// GetRoleBindings returns all the RoleBindings in the given namespace and clientset
func (kOpts *Options) GetRoleBindings() ([]rbacv1.RoleBinding, error) {
	roleBindings, getErr := kOpts.clientset.RbacV1().RoleBindings(kOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return roleBindings.Items, nil
}

// GetSecrets returns all the Secrets in the given namespace and clientset
func (kOpts *Options) GetSecrets() ([]corev1.Secret, error) {
	secrets, getErr := kOpts.clientset.CoreV1().Secrets(kOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return secrets.Items, nil
}

// GetSVC returns all the Services in the given namespace and clientset
func (kOpts *Options) GetSVC() ([]corev1.Service, error) {
	services, getErr := kOpts.clientset.CoreV1().Services(kOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return services.Items, nil
}

// GetIngress returns all the Ingresses in the given namespace and clientset
func (kOpts *Options) GetIngress() ([]v1beta1.Ingress, error) {
	ingresses, getErr := kOpts.clientset.ExtensionsV1beta1().Ingresses(kOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return ingresses.Items, nil
}
