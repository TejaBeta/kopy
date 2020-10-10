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

func TestGetDeployment(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &appv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-deployment", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-ns",
	}

	_, err := cs.AppsV1().Deployments("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetDeployments()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-deployment" {
			t.Errorf("Error while getting deployments")
		}
	}

}

func TestDeleteDeployment(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &appv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-deployment", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-ns",
	}

	_, err := cs.AppsV1().Deployments("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteDeployment("unit-test-deployment")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteDeployment("unit-test-deployment-1")
	if err == nil {
		t.Errorf("Error while deleting unexistence deployment")
	}

}

func TestCreateDeployment(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &appv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-deployment", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-ns",
	}

	_, err := options.CreateDeployment(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = cs.AppsV1().Deployments("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Fatal("Error while creating duplicate deployment")
	}

}

func TestGetConfigMap(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-configmap", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-ns",
	}

	_, err := cs.CoreV1().ConfigMaps("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetConfigMaps()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-configmap" {
			t.Errorf("Error while getting configmaps")
		}
	}

}

func TestDeleteConfigMap(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-configmap", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-ns",
	}

	_, err := cs.CoreV1().ConfigMaps("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteConfigMap("unit-test-configmap")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteConfigMap("unit-test-configmap-1")
	if err == nil {
		t.Fatal("Error while deleting known existence configmap")
	}

}

func TestCreateConfigMap(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-configmap", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-ns",
	}

	_, err := options.CreateConfigMap(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = cs.CoreV1().ConfigMaps("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Fatal("Error while creating duplicate configmap")
	}

}
