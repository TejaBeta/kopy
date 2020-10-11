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
	v1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
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

	output, err := cs.AppsV1().Deployments(options.namespace).Get(context.TODO(), "unit-test-deployment", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-deployment" {
		t.Errorf("Error while retrieving created deployment")
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

	output, err := cs.CoreV1().ConfigMaps(options.namespace).Get(context.TODO(), "unit-test-configmap", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-configmap" {
		t.Errorf("Error while retrieving created configmap")
	}

	_, err = cs.CoreV1().ConfigMaps("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while creating duplicate configmap")
	}

}

func TestGetIngress(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1beta1.Ingress{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-ingress", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-ns",
	}

	_, err := cs.ExtensionsV1beta1().Ingresses("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetIngress()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-ingress" {
			t.Errorf("Error while getting ingress")
		}
	}

}

func TestDeleteIngress(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1beta1.Ingress{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-ingress", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-ns",
	}

	_, err := cs.ExtensionsV1beta1().Ingresses("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteIngress("unit-test-ingress")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteIngress("unit-test-ingress-1")
	if err == nil {
		t.Errorf("Error while deleting non existence ingress")
	}

}

func TestCreateIngress(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1beta1.Ingress{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-ingress", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-ns",
	}

	_, err := options.CreateIngress(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := cs.ExtensionsV1beta1().Ingresses("unit-test-ns").Get(context.TODO(), "unit-test-ingress", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-ingress" {
		t.Errorf("Error while retreiving created ingress")
	}

	_, err = cs.ExtensionsV1beta1().Ingresses("unit-test-ns").Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while creating duplicate ingress")
	}

}

func TestGetNS(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-namespace", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Namespaces().Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetNS()
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-namespace" {
		t.Errorf("Error while getting ingress")
	}

}

func TestDeleteNS(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-namespace", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Namespaces().Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteNS("unit-test-namespace")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteIngress("unit-test-namespace-1")
	if err == nil {
		t.Errorf("Error while deleting non existence namespace")
	}

}

func TestCreateNS(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-namespace", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := options.CreateNS(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := cs.CoreV1().Namespaces().Get(context.TODO(), "unit-test-namespace", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-namespace" {
		t.Errorf("Error while retrieving created namespace")
	}

	_, err = cs.CoreV1().Namespaces().Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while creating duplicate namespace")
	}

}

func TestGetRoleBindings(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &rbacv1.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-rolebinding", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.RbacV1().RoleBindings(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetRoleBindings()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-rolebinding" {
			t.Errorf("Error while getting rolebindings")
		}
	}

}

func TestDeleteRoleBindings(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &rbacv1.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-rolebinding", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.RbacV1().RoleBindings(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteRBinding("unit-test-rolebinding")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteRBinding("unit-test-rolebinding-1")
	if err == nil {
		t.Errorf("Error while deleting a non existence rolebinding")
	}

}

func TestCreateRoleBindings(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &rbacv1.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-rolebinding", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := options.CreateRBinding(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := cs.RbacV1().RoleBindings(options.namespace).Get(context.TODO(), "unit-test-rolebinding", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-rolebinding" {
		t.Errorf("Created rolebinding doesn't exist")
	}

	_, err = cs.RbacV1().RoleBindings(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while creating a duplicate rolebinding")
	}

}

func TestGetRoles(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-role", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.RbacV1().Roles(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetRoles()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-role" {
			t.Errorf("Error while getting role")
		}
	}

}

func TestDeleteRole(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-role", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.RbacV1().Roles(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteRole("unit-test-role")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteRole("unit-test-role-1")
	if err == nil {
		t.Errorf("Error while deleting non-existence role")
	}
}

func TestCreateRole(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-role", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := options.CreateRole(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := cs.RbacV1().Roles(options.namespace).Get(context.TODO(), "unit-test-role", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-role" {
		t.Errorf("Error while retreiving created role")
	}

	_, err = cs.RbacV1().Roles(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while creating duplicate role")
	}

}

func TestGetSecret(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-secret", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Secrets(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetSecrets()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-secret" {
			t.Errorf("Error while getting secret")
		}
	}

}

func TestDeleteSecret(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-secret", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Secrets(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteSecret("unit-test-secret")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteSecret("unit-test-secret-1")
	if err == nil {
		t.Errorf("Error while deleting non-existence secret")
	}

}

func TestCreateSecret(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-secret", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := options.CreateSecret(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := cs.CoreV1().Secrets(options.namespace).Get(context.TODO(), "unit-test-secret", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-secret" {
		t.Errorf("Error while retreiving already created secret")
	}

	_, err = cs.CoreV1().Secrets(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while creating duplicate secret")
	}

}

func TestGetSVC(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.Service{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-service", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Services(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetSVC()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-service" {
			t.Errorf("Error while getting service")
		}
	}

}

func TestDeleteSVC(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.Service{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-service", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Services(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteSVC("unit-test-service")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteSVC("unit-test-service-1")
	if err == nil {
		t.Errorf("Error while deleting non-existence service")
	}

}

func TestCreateSVC(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.Service{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-service", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := options.CreateSVC(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := cs.CoreV1().Services(options.namespace).Get(context.TODO(), "unit-test-service", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-service" {
		t.Errorf("Error while retreiving created service")
	}

	_, err = cs.CoreV1().Services(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while trying to create a duplicate service")
	}

}
