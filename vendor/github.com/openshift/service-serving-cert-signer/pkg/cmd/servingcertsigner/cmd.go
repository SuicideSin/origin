package servingcertsigner

import (
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/runtime"

	operatorv1alpha1 "github.com/openshift/api/operator/v1alpha1"
	servicecertsignerv1alpha1 "github.com/openshift/api/servicecertsigner/v1alpha1"
	"github.com/openshift/service-serving-cert-signer/pkg/boilerplate/controllercmd"
	"github.com/openshift/service-serving-cert-signer/pkg/cmd/scheme"
	"github.com/openshift/service-serving-cert-signer/pkg/controller/servingcert/starter"
	"github.com/openshift/service-serving-cert-signer/pkg/version"
)

const (
	componentName      = "openshift-service-serving-cert-signer-serving-ca"
	componentNamespace = "openshift-service-cert-signer"
)

func NewController() *cobra.Command {
	cmd := controllercmd.
		NewControllerCommandConfig(componentName, version.Get()).
		WithNamespace(componentNamespace).
		WithConfig(&servicecertsignerv1alpha1.ServiceServingCertSignerConfig{}, scheme.ConfigScheme, servicecertsignerv1alpha1.GroupVersion).
		WithControllerFunc(controllerFunc).
		NewCommand()
	cmd.Use = "serving-cert-signer"
	cmd.Short = "Start the Service Serving Cert Signer controller"
	return cmd
}

func controllerFunc(uncastConfig runtime.Object) (controllercmd.StartFunc, *operatorv1alpha1.GenericOperatorConfig, error) {
	config := uncastConfig.(*servicecertsignerv1alpha1.ServiceServingCertSignerConfig)

	startFunc, err := starter.ToStartFunc(config)
	if err != nil {
		return nil, nil, err
	}

	// TODO we should probably supply something useful in this config
	operatorConfig := &operatorv1alpha1.GenericOperatorConfig{}

	return startFunc, operatorConfig, nil
}
