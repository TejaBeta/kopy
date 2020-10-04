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
