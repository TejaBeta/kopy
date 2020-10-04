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

	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
