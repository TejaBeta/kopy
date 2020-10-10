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
	"testing"

	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestNamespaceManipulation(t *testing.T) {

	clientset := testclient.NewSimpleClientset()
	ns := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-ns", ResourceVersion: "12345"}}

	namespace, err := clientset.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	ManipulateResource(namespace)

	if namespace.ResourceVersion != "" {
		t.Errorf("Manipulation of Namespace is failing")
	}

}

func TestConfigMapManipulation(t *testing.T) {

	clientset := testclient.NewSimpleClientset()
	input := &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-configmap", ResourceVersion: "12345"}}

	configmap, err := clientset.CoreV1().ConfigMaps("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	ManipulateResource(configmap)

	if configmap.ResourceVersion != "" {
		t.Errorf("Manipulation of Configmap is failing")
	}

}

func TestDeploymentManipulation(t *testing.T) {

	clientset := testclient.NewSimpleClientset()
	input := &appv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-deployment", ResourceVersion: "12345"}}

	deployment, err := clientset.AppsV1().Deployments("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	ManipulateResource(deployment)

	if deployment.ResourceVersion != "" {
		t.Errorf("Manipulation of Deployment is failing")
	}

}
