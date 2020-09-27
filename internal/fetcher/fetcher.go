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
	"fmt"

	log "github.com/sirupsen/logrus"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type fetchOpts struct {
	clientset *kubernetes.Clientset
	namespace string
}

func GetResources(context *rest.Config, ns string) {
	clientset, err := kubernetes.NewForConfig(context)
	if err != nil {
		panic(err)
	}

	fOpts := fetchOpts{clientset: clientset, namespace: ns}

	if fOpts.isValidateNS() {
		deployments, err := fOpts.getDeployments()
		if err != nil {
			log.Errorln(err)
			return
		}

		configmaps, err := fOpts.getConfigMaps()
		if err != nil {
			log.Errorln(err)
			return
		}

		fmt.Printf("%v, %v", deployments, configmaps)
	}
}

func (fOpts *fetchOpts) isValidateNS() bool {
	_, err := fOpts.clientset.CoreV1().Namespaces().Get(context.TODO(), fOpts.namespace, metav1.GetOptions{})
	if err != nil {
		log.Errorln(err)
		return false
	}
	return true
}

func (fOpts *fetchOpts) getDeployments() ([]appv1.Deployment, error) {
	deploymentsClient := fOpts.clientset.AppsV1().Deployments(fOpts.namespace)

	deploymentList, getErr := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}

	return deploymentList.Items, nil
}

func (fOpts *fetchOpts) getConfigMaps() ([]corev1.ConfigMap, error) {
	configmapsList, getErr := fOpts.clientset.CoreV1().ConfigMaps(fOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if getErr != nil {
		return nil, getErr
	}

	return configmapsList.Items, nil
}
