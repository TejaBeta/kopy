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
package fetcher

import (
	"context"

	"k8s.io/client-go/rest"

	log "github.com/sirupsen/logrus"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// FetchOpts to pass to all the methods
type FetchOpts struct {
	clientset *kubernetes.Clientset
	namespace string
}

// GetFetchOpts generates required options
func GetFetchOpts(context *rest.Config, ns string) (*FetchOpts, error) {
	cs, err := kubernetes.NewForConfig(context)
	if err != nil {
		return nil, err
	}

	return &FetchOpts{
		clientset: cs,
		namespace: ns,
	}, nil
}

// IsValidNS validates if the namespace exists or not
func (fOpts *FetchOpts) IsValidNS() bool {
	_, err := fOpts.clientset.CoreV1().Namespaces().Get(context.TODO(), fOpts.namespace, metav1.GetOptions{})
	if err != nil {
		log.Errorln(err)
		return false
	}
	return true
}

// GetDeployments returns all the Deployments in the given namespace and clientset
func (fOpts *FetchOpts) GetDeployments() ([]appv1.Deployment, error) {
	deploymentsClient := fOpts.clientset.AppsV1().Deployments(fOpts.namespace)

	deploymentList, getErr := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return deploymentList.Items, nil
}

// GetConfigMaps returns all the Configmaps in the given namespace and clientset
func (fOpts *FetchOpts) GetConfigMaps() ([]corev1.ConfigMap, error) {
	configmapsList, getErr := fOpts.clientset.CoreV1().ConfigMaps(fOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return configmapsList.Items, nil
}

// GetClusterRoles returns all the ClusterRoles in the given namespace and clientset
func (fOpts *FetchOpts) GetClusterRoles() ([]rbacv1.ClusterRole, error) {
	clusterRoles, getErr := fOpts.clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return clusterRoles.Items, nil
}

// GetRoles returns all the Roles in the given namespace and clientset
func (fOpts *FetchOpts) GetRoles() ([]rbacv1.Role, error) {
	roles, getErr := fOpts.clientset.RbacV1().Roles(fOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return roles.Items, nil
}

// GetClusterRoleBindings returns all the ClusterRoleBindings in the given namespace and clientset
func (fOpts *FetchOpts) GetClusterRoleBindings() ([]rbacv1.ClusterRoleBinding, error) {
	clusterRoleBindings, getErr := fOpts.clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return clusterRoleBindings.Items, nil
}

// GetRoleBindings returns all the RoleBindings in the given namespace and clientset
func (fOpts *FetchOpts) GetRoleBindings() ([]rbacv1.RoleBinding, error) {
	roleBindings, getErr := fOpts.clientset.RbacV1().RoleBindings(fOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return roleBindings.Items, nil
}

// GetSecrets returns all the Secrets in the given namespace and clientset
func (fOpts *FetchOpts) GetSecrets() ([]corev1.Secret, error) {
	secrets, getErr := fOpts.clientset.CoreV1().Secrets(fOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return secrets.Items, nil
}

// GetSVC returns all the Services in the given namespace and clientset
func (fOpts *FetchOpts) GetSVC() ([]corev1.Service, error) {
	services, getErr := fOpts.clientset.CoreV1().Services(fOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return services.Items, nil
}

// GetIngress returns all the Ingresses in the given namespace and clientset
func (fOpts *FetchOpts) GetIngress() ([]v1beta1.Ingress, error) {
	ingresses, getErr := fOpts.clientset.ExtensionsV1beta1().Ingresses(fOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return ingresses.Items, nil
}
