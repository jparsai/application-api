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
	"testing"

	"github.com/redhat-appstudio/application-api/api/v1beta1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func TestV1toV2Conversion(t *testing.T) {

	tests := []struct {
		name              string
		v1EnvTestData     Environment
		v2EnvExpectedData *v1beta1.Environment
	}{
		{
			name: "Convert Environment v1alpha1 version to v1beta1 version with all fields.",
			v1EnvTestData: Environment{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
				},
				Spec: EnvironmentSpec{
					Type:               EnvironmentType_POC,
					DisplayName:        "Test Name",
					DeploymentStrategy: DeploymentStrategy_AppStudioAutomated,
					ParentEnvironment:  "dev",
					Tags:               []string{"test-tag-1", "test-tag-2"},
					Configuration: EnvironmentConfiguration{
						Env: []EnvVarPair{
							{Name: "env-key-1", Value: "env-value-1"},
							{Name: "env-key-2", Value: "env-value-2"},
						},
						Target: EnvironmentTarget{
							DeploymentTargetClaim: DeploymentTargetClaimConfig{
								ClaimName: "test-claim",
							},
						},
					},
					UnstableConfigurationFields: &UnstableEnvironmentConfiguration{
						ClusterType: ConfigurationClusterType_OpenShift,
						KubernetesClusterCredentials: KubernetesClusterCredentials{
							TargetNamespace:            "test-namespace",
							APIURL:                     "https://api.com:6443",
							IngressDomain:              "test-domain",
							ClusterCredentialsSecret:   "test-secret",
							AllowInsecureSkipTLSVerify: true,
							Namespaces:                 []string{"namespace-1", "namespace-2"},
							ClusterResources:           true,
						},
					},
				},
			},
			v2EnvExpectedData: &v1beta1.Environment{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
				},
				Spec: v1beta1.EnvironmentSpec{
					Type:               v1beta1.EnvironmentType_POC,
					DisplayName:        "Test Name",
					DeploymentStrategy: v1beta1.DeploymentStrategy_AppStudioAutomated,
					ParentEnvironment:  "dev",
					Tags:               []string{"test-tag-1", "test-tag-2"},
					Configuration: v1beta1.EnvironmentConfiguration{
						Env: []v1beta1.EnvVarPair{
							{Name: "env-key-1", Value: "env-value-1"},
							{Name: "env-key-2", Value: "env-value-2"},
						},
						Target: v1beta1.EnvironmentTarget{
							DeploymentTargetClaim: v1beta1.DeploymentTargetClaimConfig{
								ClaimName: "test-claim",
							},
						},
					},
					Target: &v1beta1.TargetConfiguration{
						ClusterType: v1beta1.ConfigurationClusterType_OpenShift,
						KubernetesClusterCredentials: v1beta1.KubernetesClusterCredentials{
							TargetNamespace:            "test-namespace",
							APIURL:                     "https://api.com:6443",
							IngressDomain:              "test-domain",
							ClusterCredentialsSecret:   "test-secret",
							AllowInsecureSkipTLSVerify: true,
							Namespaces:                 []string{"namespace-1", "namespace-2"},
							ClusterResources:           true,
						},
					},
				},
			},
		}, {
			name: "Convert Environment v1alpha1 version to v1beta1 version with missing fields (test1).",
			// Missing Spec.Configuration.Target field.
			v1EnvTestData: Environment{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
				},
				Spec: EnvironmentSpec{
					Type:               EnvironmentType_POC,
					DisplayName:        "Test Name",
					DeploymentStrategy: DeploymentStrategy_AppStudioAutomated,
					ParentEnvironment:  "dev",
					Tags:               []string{"test-tag-1", "test-tag-2"},
					Configuration: EnvironmentConfiguration{
						Env: []EnvVarPair{
							{Name: "env-key-1", Value: "env-value-1"},
							{Name: "env-key-2", Value: "env-value-2"},
						},
					},
					UnstableConfigurationFields: &UnstableEnvironmentConfiguration{
						ClusterType: ConfigurationClusterType_OpenShift,
						KubernetesClusterCredentials: KubernetesClusterCredentials{
							TargetNamespace:            "test-namespace",
							APIURL:                     "https://api.com:6443",
							IngressDomain:              "test-domain",
							ClusterCredentialsSecret:   "test-secret",
							AllowInsecureSkipTLSVerify: true,
							Namespaces:                 []string{"namespace-1", "namespace-2"},
							ClusterResources:           true,
						},
					},
				},
			},
			v2EnvExpectedData: &v1beta1.Environment{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
				},
				Spec: v1beta1.EnvironmentSpec{
					Type:               v1beta1.EnvironmentType_POC,
					DisplayName:        "Test Name",
					DeploymentStrategy: v1beta1.DeploymentStrategy_AppStudioAutomated,
					ParentEnvironment:  "dev",
					Tags:               []string{"test-tag-1", "test-tag-2"},
					Configuration: v1beta1.EnvironmentConfiguration{
						Env: []v1beta1.EnvVarPair{
							{Name: "env-key-1", Value: "env-value-1"},
							{Name: "env-key-2", Value: "env-value-2"},
						},
					},
					Target: &v1beta1.TargetConfiguration{
						ClusterType: v1beta1.ConfigurationClusterType_OpenShift,
						KubernetesClusterCredentials: v1beta1.KubernetesClusterCredentials{
							TargetNamespace:            "test-namespace",
							APIURL:                     "https://api.com:6443",
							IngressDomain:              "test-domain",
							ClusterCredentialsSecret:   "test-secret",
							AllowInsecureSkipTLSVerify: true,
							Namespaces:                 []string{"namespace-1", "namespace-2"},
							ClusterResources:           true,
						},
					},
				},
			},
		}, {
			name: "Convert Environment v1alpha1 version to v1beta1 version with missing fields (test2).",
			// Missing Spec.UnstableConfigurationFields.ClusterType field.
			v1EnvTestData: Environment{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
				},
				Spec: EnvironmentSpec{
					Type:               EnvironmentType_POC,
					DisplayName:        "Test Name",
					DeploymentStrategy: DeploymentStrategy_AppStudioAutomated,
					ParentEnvironment:  "dev",
					Tags:               []string{"test-tag-1", "test-tag-2"},
					Configuration: EnvironmentConfiguration{
						Env: []EnvVarPair{
							{Name: "env-key-1", Value: "env-value-1"},
							{Name: "env-key-2", Value: "env-value-2"},
						},
					},
					UnstableConfigurationFields: &UnstableEnvironmentConfiguration{
						KubernetesClusterCredentials: KubernetesClusterCredentials{
							TargetNamespace:            "test-namespace",
							APIURL:                     "https://api.com:6443",
							IngressDomain:              "test-domain",
							ClusterCredentialsSecret:   "test-secret",
							AllowInsecureSkipTLSVerify: true,
							Namespaces:                 []string{"namespace-1", "namespace-2"},
							ClusterResources:           true,
						},
					},
				},
			},
			v2EnvExpectedData: &v1beta1.Environment{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
				},
				Spec: v1beta1.EnvironmentSpec{
					Type:               v1beta1.EnvironmentType_POC,
					DisplayName:        "Test Name",
					DeploymentStrategy: v1beta1.DeploymentStrategy_AppStudioAutomated,
					ParentEnvironment:  "dev",
					Tags:               []string{"test-tag-1", "test-tag-2"},
					Configuration: v1beta1.EnvironmentConfiguration{
						Env: []v1beta1.EnvVarPair{
							{Name: "env-key-1", Value: "env-value-1"},
							{Name: "env-key-2", Value: "env-value-2"},
						},
					},
					Target: &v1beta1.TargetConfiguration{
						KubernetesClusterCredentials: v1beta1.KubernetesClusterCredentials{
							TargetNamespace:            "test-namespace",
							APIURL:                     "https://api.com:6443",
							IngressDomain:              "test-domain",
							ClusterCredentialsSecret:   "test-secret",
							AllowInsecureSkipTLSVerify: true,
							Namespaces:                 []string{"namespace-1", "namespace-2"},
							ClusterResources:           true,
						},
					},
				},
			},
		}, {
			name: "Convert Environment v1alpha1 version to v1beta1 version with missing fields (test3).",
			// Missing Spec.UnstableConfigurationFields.KubernetesClusterCredentials field.
			v1EnvTestData: Environment{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
				},
				Spec: EnvironmentSpec{
					Type:               EnvironmentType_POC,
					DisplayName:        "Test Name",
					DeploymentStrategy: DeploymentStrategy_AppStudioAutomated,
					ParentEnvironment:  "dev",
					Tags:               []string{"test-tag-1", "test-tag-2"},
					Configuration: EnvironmentConfiguration{
						Env: []EnvVarPair{
							{Name: "env-key-1", Value: "env-value-1"},
							{Name: "env-key-2", Value: "env-value-2"},
						},
					},
					UnstableConfigurationFields: &UnstableEnvironmentConfiguration{
						ClusterType: ConfigurationClusterType_OpenShift,
					},
				},
			},
			v2EnvExpectedData: &v1beta1.Environment{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
				},
				Spec: v1beta1.EnvironmentSpec{
					Type:               v1beta1.EnvironmentType_POC,
					DisplayName:        "Test Name",
					DeploymentStrategy: v1beta1.DeploymentStrategy_AppStudioAutomated,
					ParentEnvironment:  "dev",
					Tags:               []string{"test-tag-1", "test-tag-2"},
					Configuration: v1beta1.EnvironmentConfiguration{
						Env: []v1beta1.EnvVarPair{
							{Name: "env-key-1", Value: "env-value-1"},
							{Name: "env-key-2", Value: "env-value-2"},
						},
					},
					Target: &v1beta1.TargetConfiguration{
						ClusterType: v1beta1.ConfigurationClusterType_OpenShift,
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Set v1beta1 object in Hub, converted values will be set in this object.
			var hub conversion.Hub = &v1beta1.Environment{}

			// Call ConvertTo function to convert v1alpha1 version to v1beta1
			test.v1EnvTestData.ConvertTo(hub)

			// Fetch the converted object
			result := hub.(*v1beta1.Environment)

			// Compare converted object with expected.
			assert.Equal(t, test.v2EnvExpectedData, result)
		})
	}
}
