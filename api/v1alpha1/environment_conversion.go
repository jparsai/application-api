//
// Copyright 2023 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"fmt"

	"github.com/redhat-appstudio/application-api/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this Memcached to the Hub version (vbeta1).
func (src *Environment) ConvertTo(dstRaw conversion.Hub) error {
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	// fetch v1beta1 version from Hub, converted values will be set in this object
	dst := dstRaw.(*v1beta1.Environment)

	// copy ObjectMeta from v1alpha1 to v1beta1 version
	dst.ObjectMeta = src.ObjectMeta

	// copy Spec fields from v1alpha1 to v1beta1 version
	dst.Spec = v1beta1.EnvironmentSpec{
		DisplayName:        src.Spec.DisplayName,
		DeploymentStrategy: v1beta1.DeploymentStrategyType(src.Spec.DeploymentStrategy),
		ParentEnvironment:  src.Spec.ParentEnvironment,
		Tags:               src.Spec.Tags,
	}

	fmt.Println("111.  $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	// if v1alpha1 version has src.Spec.Configuration.Env field then copy it to v1beta1
	if src.Spec.Configuration.Env != nil {
		dst.Spec.Configuration.Env = []v1beta1.EnvVarPair{}

		for _, env := range src.Spec.Configuration.Env {
			dst.Spec.Configuration.Env = append(dst.Spec.Configuration.Env, v1beta1.EnvVarPair(env))
		}
	}
	fmt.Println("222.  $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	fmt.Println("00. ####### src.Spec.Configuration.Target.DeploymentTargetClaim.ClaimName === ", src.Spec.Configuration.Target.DeploymentTargetClaim.ClaimName)

	// if v1alpha1 version has Spec.Configuration.Target field then copy it to v1beta1
	if src.Spec.Configuration.Target.DeploymentTargetClaim.ClaimName != "" {
		fmt.Println("00. ##### inside condition")
		// This filed is renamed and moved to Target in v1beta1
		dst.Spec.Target = &v1beta1.TargetConfiguration{
			Claim: v1beta1.TargetClaim{
				DeploymentTargetClaim: v1beta1.DeploymentTargetClaimConfig{
					ClaimName: src.Spec.Configuration.Target.DeploymentTargetClaim.ClaimName,
				},
			},
		}
	}

	// if v1alpha1 has Spec.UnstableConfigurationFields field then copy it to v1beta1
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("11. ####### src.Spec.UnstableConfigurationFields === ", src.Spec.UnstableConfigurationFields)
	if src.Spec.UnstableConfigurationFields != nil {

		fmt.Println("22. ####### inside condition")
		fmt.Println("33. ####### dst.Spec.Target === ", dst.Spec.Target)
		if dst.Spec.Target == nil {
			fmt.Println("44. ##### inside condition")
			dst.Spec.Target = &v1beta1.TargetConfiguration{}
		}

		fmt.Println("55. ##### src.Spec.UnstableConfigurationFields.ClusterType == ", src.Spec.UnstableConfigurationFields.ClusterType)

		dst.Spec.Target.ClusterType = v1beta1.ConfigurationClusterType(string(src.Spec.UnstableConfigurationFields.ClusterType))

		fmt.Println("66. ##### src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.TargetNamespace == ", src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.TargetNamespace)

		dst.Spec.Target.KubernetesClusterCredentials = v1beta1.KubernetesClusterCredentials{
			TargetNamespace:            src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.TargetNamespace,
			APIURL:                     src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.APIURL,
			IngressDomain:              src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.IngressDomain,
			ClusterCredentialsSecret:   src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.ClusterCredentialsSecret,
			AllowInsecureSkipTLSVerify: src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.AllowInsecureSkipTLSVerify,
			Namespaces:                 src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.Namespaces,
			ClusterResources:           src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.ClusterResources,
		}
	}

	// copy Status from v1alpha1 to v1beta1 version
	dst.Status = v1beta1.EnvironmentStatus(src.Status)

	return nil
}

// ConvertFrom converts from the Hub version (vbeta1) to this version.
func (dst *Environment) ConvertFrom(srcRaw conversion.Hub) error {

	src := srcRaw.(*v1beta1.Environment)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec = EnvironmentSpec{
		Type:               EnvironmentType(v1beta1.EnvironmentType_NonPOC),
		DisplayName:        src.Spec.DisplayName,
		DeploymentStrategy: DeploymentStrategyType(src.Spec.DeploymentStrategy),
		ParentEnvironment:  src.Spec.ParentEnvironment,
		Tags:               src.Spec.Tags,
	}

	if src.Spec.Configuration.Env != nil {
		dst.Spec.Configuration.Env = []EnvVarPair{}

		for _, env := range src.Spec.Configuration.Env {
			dst.Spec.Configuration.Env = append(dst.Spec.Configuration.Env, EnvVarPair(env))
		}
	}

	if src.Spec.Target != nil {
		if src.Spec.Target.Claim.DeploymentTargetClaim.ClaimName != "" {
			dst.Spec.Configuration.Target = EnvironmentTarget{
				DeploymentTargetClaim: DeploymentTargetClaimConfig{
					ClaimName: src.Spec.Target.Claim.DeploymentTargetClaim.ClaimName,
				},
			}
		}

		dst.Spec.UnstableConfigurationFields = &UnstableEnvironmentConfiguration{
			ClusterType: ConfigurationClusterType(string(src.Spec.Target.ClusterType)),
		}

		dst.Spec.UnstableConfigurationFields.KubernetesClusterCredentials = KubernetesClusterCredentials{
			TargetNamespace:            src.Spec.Target.KubernetesClusterCredentials.TargetNamespace,
			APIURL:                     src.Spec.Target.KubernetesClusterCredentials.APIURL,
			IngressDomain:              src.Spec.Target.KubernetesClusterCredentials.IngressDomain,
			ClusterCredentialsSecret:   src.Spec.Target.KubernetesClusterCredentials.ClusterCredentialsSecret,
			AllowInsecureSkipTLSVerify: src.Spec.Target.KubernetesClusterCredentials.AllowInsecureSkipTLSVerify,
			Namespaces:                 src.Spec.Target.KubernetesClusterCredentials.Namespaces,
			ClusterResources:           src.Spec.Target.KubernetesClusterCredentials.ClusterResources,
		}
	}

	dst.Status = EnvironmentStatus(src.Status)

	return nil
}
