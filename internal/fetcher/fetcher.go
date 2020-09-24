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
	"errors"
	"fmt"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetResources(context *rest.Config, ns string) {
	clientset, err := kubernetes.NewForConfig(context)
	if err != nil {
		panic(err)
	}

	deployments, err := getDeployments(clientset, ns)
	if err != nil {
		panic(err)
	}
	fmt.Println(deployments)
}

func getDeployments(clientset *kubernetes.Clientset, ns string) ([]v1.Deployment, error) {
	deploymentsClient := clientset.AppsV1().Deployments(ns)

	deploymentList, getErr := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	}

	if len(deploymentList.Items) == 0 {
		return nil, errors.New("No deployments exists in the given namespace")
	}

	return deploymentList.Items, nil
}
