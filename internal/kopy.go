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
package internal

import (
	log "github.com/sirupsen/logrus"
	"github.com/tejabeta/kopy/internal/options"
	"github.com/tejabeta/kopy/pkg/koperator"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
)

// KopyResources as name suggests a struct type to hold all the resources
type kopyResources struct {
	Deployments  []appv1.Deployment
	ConfigMaps   []corev1.ConfigMap
	Roles        []rbacv1.Role
	RoleBindings []rbacv1.RoleBinding
	Secrets      []corev1.Secret
	Services     []corev1.Service
	Ingresses    []v1beta1.Ingress
}

// Kopy functionality goes here
func Kopy(kopyOptions *options.KopyOptions) {
	sourceFOpts, err := koperator.GetOpts(kopyOptions.CurrentContext, kopyOptions.Namespace)
	if err != nil {
		log.Errorln(err)
		return
	}

	destFOpts, err := koperator.GetOpts(kopyOptions.DestinationContext, kopyOptions.Namespace)
	if err != nil {
		log.Errorln(err)
	}

	if sourceFOpts.IsValidNS() {
		sourceResources, err := getKopyResources(sourceFOpts)
		if err != nil {
			log.Fatalln(err)
			return
		}

		log.Println(sourceResources, destFOpts)

		if destFOpts.IsValidNS() {
			log.Info("Namespace ", kopyOptions.Namespace, " exists at destination all resource will be overwritten.")
		} else {
			log.Info("No namespace ", kopyOptions.Namespace, " found in destination context.")
			log.Info("Namespace and resources will be created in the destination context.")
		}
	} else {
		log.Info("No namespace ", kopyOptions.Namespace, " found in source context.")
	}
}

func getKopyResources(kOpts *koperator.Options) (*kopyResources, error) {
	deployments, err := kOpts.GetDeployments()
	if err != nil {
		return nil, err
	}

	configMaps, err := kOpts.GetConfigMaps()
	if err != nil {
		return nil, err
	}

	roles, err := kOpts.GetRoles()
	if err != nil {
		return nil, err
	}

	roleBindings, err := kOpts.GetRoleBindings()
	if err != nil {
		return nil, err
	}

	secrets, err := kOpts.GetSecrets()
	if err != nil {
		return nil, err
	}

	services, err := kOpts.GetSVC()
	if err != nil {
		return nil, err
	}

	ingresses, err := kOpts.GetIngress()
	if err != nil {
		return nil, err
	}

	kopyResources := kopyResources{
		Deployments:  deployments.Items,
		ConfigMaps:   configMaps.Items,
		Roles:        roles.Items,
		RoleBindings: roleBindings.Items,
		Secrets:      secrets.Items,
		Services:     services.Items,
		Ingresses:    ingresses.Items,
	}

	return &kopyResources, nil
}
