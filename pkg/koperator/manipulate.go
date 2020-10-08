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

	corev1 "k8s.io/api/core/v1"
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
	default:
		fmt.Println("In default", v)
	}
	return
}
