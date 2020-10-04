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

// GetClusterRoles returns all the ClusterRoles in the given namespace and clientset
func (kOpts *Options) GetClusterRoles() ([]rbacv1.ClusterRole, error) {
	clusterRoles, getErr := kOpts.clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return clusterRoles.Items, nil
}
