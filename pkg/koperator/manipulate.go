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
	"fmt"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
)

// ManipulateResource helps to manipulate resources
func ManipulateResource(x interface{}) {
	switch v := x.(type) {
	case *corev1.Namespace:
		input := x.(*corev1.Namespace)
		input.ResourceVersion = ""
	case *corev1.ConfigMap:
		input := x.(*corev1.ConfigMap)
		input.ResourceVersion = ""
	case *appv1.Deployment:
		input := x.(*appv1.Deployment)
		input.ResourceVersion = ""
	case *v1beta1.Ingress:
		input := x.(*v1beta1.Ingress)
		input.ResourceVersion = ""
	case *rbacv1.RoleBinding:
		input := x.(*rbacv1.RoleBinding)
		input.ResourceVersion = ""
	case *rbacv1.Role:
		input := x.(*rbacv1.Role)
		input.ResourceVersion = ""
	case *corev1.Secret:
		input := x.(*corev1.Secret)
		input.ResourceVersion = ""
	case *corev1.Service:
		input := x.(*corev1.Service)
		input.ResourceVersion, input.Spec.ClusterIP = "", ""
	default:
		fmt.Println("In default", v)
	}
	return
}
