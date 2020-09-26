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

	log "github.com/sirupsen/logrus"

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

	if validateNS(clientset, ns) {
		deployments, err := getDeployments(clientset, ns)
		if err != nil {
			log.Errorln(err, ns)
			return
		}
		fmt.Println(deployments)
	}
}

func validateNS(clientset *kubernetes.Clientset, name string) bool {
	_, err := clientset.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Errorln("Error while retrieving namespace: " + name)
		return false
	}
	return true
}

func getDeployments(clientset *kubernetes.Clientset, ns string) ([]v1.Deployment, error) {
	deploymentsClient := clientset.AppsV1().Deployments(ns)

	deploymentList, getErr := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		log.Errorf("Failed to get latest version of Deployment: %v", getErr)
	}

	if len(deploymentList.Items) == 0 {
		return nil, errors.New("No deployments exists in the given namespace: ")
	}

	return deploymentList.Items, nil
}
