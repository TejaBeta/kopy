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

	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetDeployments returns all the Deployments in the given namespace and clientset
func (kOpts *Options) GetDeployments() (result *appv1.DeploymentList, err error) {
	result, err = kOpts.clientset.
		AppsV1().
		Deployments(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteDeployment is a method to delete a provided deployment name
func (kOpts *Options) DeleteDeployment(name string) (err error) {
	err = kOpts.clientset.
		AppsV1().
		Deployments(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateDeployment method to create a deployment
func (kOpts *Options) CreateDeployment(deployment *appv1.Deployment) (result *appv1.Deployment, err error) {
	result, err = kOpts.clientset.
		AppsV1().
		Deployments(kOpts.namespace).
		Create(context.TODO(), deployment, metav1.CreateOptions{})
	return
}
