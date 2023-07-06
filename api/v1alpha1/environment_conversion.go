package v1alpha1

import (
	"fmt"

	"github.com/redhat-appstudio/application-api/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

//func (r *Environment) SetupWebhookWithManager(mgr ctrl.Manager) error {
//	return ctrl.NewWebhookManagedBy(mgr).
//		For(r).
//		Complete()
//}

// ConvertTo converts this Memcached to the Hub version (vbeta1).
func (src *Environment) ConvertTo(dstRaw conversion.Hub) error {

	fmt.Println("1111 $$$$$$$$$$$$$$$$$$")

	dst := dstRaw.(*v1beta1.Environment)

	dst.ObjectMeta = src.ObjectMeta
	dst.Spec = v1beta1.EnvironmentSpec{
		Type:               v1beta1.EnvironmentType(src.Spec.Type),
		DisplayName:        src.Spec.DisplayName,
		DeploymentStrategy: v1beta1.DeploymentStrategyType(src.Spec.DeploymentStrategy),
		ParentEnvironment:  src.Spec.ParentEnvironment,
		Tags:               src.Spec.Tags,
		Configuration: v1beta1.EnvironmentConfiguration{
			Target: v1beta1.EnvironmentTarget{
				DeploymentTargetClaim: v1beta1.DeploymentTargetClaimConfig{
					ClaimName: src.Spec.Configuration.Target.DeploymentTargetClaim.ClaimName,
				},
			},
		},
		Target: &v1beta1.TargetConfiguration{
			ClusterType: v1beta1.ConfigurationClusterType(src.Spec.UnstableConfigurationFields.ClusterType),
			KubernetesClusterCredentials: v1beta1.KubernetesClusterCredentials{
				TargetNamespace:            src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.TargetNamespace,
				APIURL:                     src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.APIURL,
				IngressDomain:              src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.IngressDomain,
				ClusterCredentialsSecret:   src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.ClusterCredentialsSecret,
				AllowInsecureSkipTLSVerify: src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.AllowInsecureSkipTLSVerify,
				Namespaces:                 src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.Namespaces,
				ClusterResources:           src.Spec.UnstableConfigurationFields.KubernetesClusterCredentials.ClusterResources,
			},
		},
	}

	if src.Spec.Configuration.Env != nil {
		for _, env := range src.Spec.Configuration.Env {
			dst.Spec.Configuration.Env = append(dst.Spec.Configuration.Env, v1beta1.EnvVarPair(env))
		}
	}

	return nil
}

// ConvertFrom converts from the Hub version (vbeta1) to this version.
func (dst *Environment) ConvertFrom(srcRaw conversion.Hub) error {

	fmt.Println("2222 $$$$$$$$$$$$$$$$$$")

	src := srcRaw.(*v1beta1.Environment)

	dst.ObjectMeta = src.ObjectMeta
	dst.Spec = EnvironmentSpec{
		Type:               EnvironmentType(src.Spec.Type),
		DisplayName:        src.Spec.DisplayName,
		DeploymentStrategy: DeploymentStrategyType(src.Spec.DeploymentStrategy),
		ParentEnvironment:  src.Spec.ParentEnvironment,
		Tags:               src.Spec.Tags,
		Configuration: EnvironmentConfiguration{
			Target: EnvironmentTarget{
				DeploymentTargetClaim: DeploymentTargetClaimConfig{
					ClaimName: src.Spec.Configuration.Target.DeploymentTargetClaim.ClaimName,
				},
			},
		},
		UnstableConfigurationFields: &UnstableEnvironmentConfiguration{
			ClusterType: ConfigurationClusterType(src.Spec.Target.ClusterType),
			KubernetesClusterCredentials: KubernetesClusterCredentials{
				TargetNamespace:            src.Spec.Target.KubernetesClusterCredentials.TargetNamespace,
				APIURL:                     src.Spec.Target.KubernetesClusterCredentials.APIURL,
				IngressDomain:              src.Spec.Target.KubernetesClusterCredentials.IngressDomain,
				ClusterCredentialsSecret:   src.Spec.Target.KubernetesClusterCredentials.ClusterCredentialsSecret,
				AllowInsecureSkipTLSVerify: src.Spec.Target.KubernetesClusterCredentials.AllowInsecureSkipTLSVerify,
				Namespaces:                 src.Spec.Target.KubernetesClusterCredentials.Namespaces,
				ClusterResources:           src.Spec.Target.KubernetesClusterCredentials.ClusterResources,
			},
		},
	}

	if src.Spec.Configuration.Env != nil {
		for _, env := range src.Spec.Configuration.Env {
			dst.Spec.Configuration.Env = append(dst.Spec.Configuration.Env, EnvVarPair(env))
		}
	}

	return nil
}
