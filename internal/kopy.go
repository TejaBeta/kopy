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
	"github.com/tejabeta/kopy/internal/fetcher"
	"github.com/tejabeta/kopy/internal/options"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
)

// KopyResources as name suggests a struct type to hold all the resources
type kopyResources struct {
	Deployments         []appv1.Deployment
	ConfigMaps          []corev1.ConfigMap
	ClusterRoles        []rbacv1.ClusterRole
	Roles               []rbacv1.Role
	ClusterRoleBindings []rbacv1.ClusterRoleBinding
	RoleBindings        []rbacv1.RoleBinding
	Secrets             []corev1.Secret
	Services            []corev1.Service
	Ingresses           []v1beta1.Ingress
}

// Kopy functionality goes here
func Kopy(kopyOptions *options.KopyOptions) {
	sourceFOpts, err := fetcher.GetFetchOpts(kopyOptions.CurrentContext, kopyOptions.Namespace)
	if err != nil {
		log.Errorln(err)
		return
	}

	destFOpts, err := fetcher.GetFetchOpts(kopyOptions.DestinationContext, kopyOptions.Namespace)
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

func getKopyResources(fetcherOpts *fetcher.FetchOpts) (*kopyResources, error) {
	deployments, err := fetcherOpts.GetDeployments()
	if err != nil {
		return nil, err
	}

	configMaps, err := fetcherOpts.GetConfigMaps()
	if err != nil {
		return nil, err
	}

	clusterRoles, err := fetcherOpts.GetClusterRoles()
	if err != nil {
		return nil, err
	}

	clusterRoleBindings, err := fetcherOpts.GetClusterRoleBindings()
	if err != nil {
		return nil, err
	}

	roles, err := fetcherOpts.GetRoles()
	if err != nil {
		return nil, err
	}

	roleBindings, err := fetcherOpts.GetRoleBindings()
	if err != nil {
		return nil, err
	}

	secrets, err := fetcherOpts.GetSecrets()
	if err != nil {
		return nil, err
	}

	services, err := fetcherOpts.GetSVC()
	if err != nil {
		return nil, err
	}

	ingresses, err := fetcherOpts.GetIngress()
	if err != nil {
		return nil, err
	}

	kopyResources := kopyResources{
		Deployments:         deployments,
		ConfigMaps:          configMaps,
		ClusterRoles:        clusterRoles,
		ClusterRoleBindings: clusterRoleBindings,
		Roles:               roles,
		RoleBindings:        roleBindings,
		Secrets:             secrets,
		Services:            services,
		Ingresses:           ingresses,
	}

	return &kopyResources, nil
}
