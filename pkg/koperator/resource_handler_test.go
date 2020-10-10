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
