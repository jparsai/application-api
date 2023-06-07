/*
Copyright 2023.

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

package v1alpha1

import (
	"github.com/redhat-appstudio/application-api/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

//func (r *Environment) SetupWebhookWithManager(mgr ctrl.Manager) error {
//	return ctrl.NewWebhookManagedBy(mgr).
//		For(r).
//		Complete()
//}

// Hub marks this type as a conversion hub.
// ConvertTo converts this ITS to the Hub version (v1beta1).
func (src *Environment) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.Environment)
	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.DisplayName = src.Spec.DisplayName
	dst.Spec.DeploymentStrategy = v1alpha2.DeploymentStrategyType(src.Spec.DeploymentStrategy)
	dst.Spec.ParentEnvironment = src.Spec.ParentEnvironment
	dst.Spec.Tags = src.Spec.Tags
	dst.Spec.DisplayName = src.Spec.DisplayName
	//dst.Spec.Configuration = src.Spec.Configuration
	//dst.Spec.Target = src.Spec.UnstableConfigurationFields

	return nil
}

func (dst *Environment) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.Environment)
	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.DisplayName = src.Spec.DisplayName
	dst.Spec.DeploymentStrategy = DeploymentStrategyType(src.Spec.DeploymentStrategy)
	dst.Spec.ParentEnvironment = src.Spec.ParentEnvironment
	dst.Spec.Tags = src.Spec.Tags
	dst.Spec.DisplayName = src.Spec.DisplayName
	//dst.Spec.Configuration = src.Spec.Configuration
	//dst.Spec.UnstableConfigurationFields = src.Spec.Target

	return nil
}
