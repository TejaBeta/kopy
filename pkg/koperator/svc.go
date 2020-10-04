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

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
