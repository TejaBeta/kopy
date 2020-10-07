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
		Create(context.TODO(), manConfigMap(configmap), metav1.CreateOptions{})
	return
}

func manConfigMap(configmap *corev1.ConfigMap) *corev1.ConfigMap {
	configmap.ResourceVersion = ""
	return configmap
}
