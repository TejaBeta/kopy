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

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNS validates if the namespace exists or not
func (kOpts *Options) GetNS() (result *v1.Namespace, err error) {
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
func (kOpts *Options) CreateNS(namespace *v1.Namespace) (result *v1.Namespace, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		Namespaces().
		Create(context.TODO(), manNS(namespace), metav1.CreateOptions{})
	return
}

func manNS(ns *v1.Namespace) *v1.Namespace {
	ns.ResourceVersion = ""
	return ns
}
