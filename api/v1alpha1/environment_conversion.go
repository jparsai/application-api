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
	fmt.Println("1111 $$$$$$$$$$$$$$$$$$")

	dst := dstRaw.(*v1beta1.Environment)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.DisplayName = "DisplayName111"
	dst.Spec.DeploymentStrategy = v1beta1.DeploymentStrategyType(src.Spec.DeploymentStrategy)
	dst.Spec.ParentEnvironment = "ParentEnvironment111"
	dst.Spec.Tags = src.Spec.Tags

	return nil
}

// ConvertFrom converts from the Hub version (vbeta1) to this version.
func (dst *Environment) ConvertFrom(srcRaw conversion.Hub) error {

	fmt.Println("2222 $$$$$$$$$$$$$$$$$$")
	fmt.Println("2222 $$$$$$$$$$$$$$$$$$")

	src := srcRaw.(*v1beta1.Environment)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.DisplayName = "DisplayName222"
	dst.Spec.DeploymentStrategy = DeploymentStrategyType(src.Spec.DeploymentStrategy)
	dst.Spec.ParentEnvironment = "ParentEnvironment222"
	dst.Spec.Tags = src.Spec.Tags

	return nil
}
