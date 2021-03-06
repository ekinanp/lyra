package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	awshandler "github.com/lyraproj/lyra/cmd/goplugin-tf-aws/handler"
	googlehandler "github.com/lyraproj/lyra/cmd/goplugin-tf-google/handler"
	kuberneteshandler "github.com/lyraproj/lyra/cmd/goplugin-tf-kubernetes/handler"
	"github.com/lyraproj/lyra/pkg/bridge"
	"github.com/lyraproj/puppet-evaluator/eval"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/terraform-providers/terraform-provider-google/google"
	"github.com/terraform-providers/terraform-provider-kubernetes/kubernetes"
)

func main() {
	eval.Puppet.Do(func(c eval.Context) {

		// AWS - you will need default AWS credentials available or this step will fail
		bridge.GeneratePP(c, awshandler.Server(c, aws.Provider().(*schema.Provider)), "TerraformAws", "plugins/aaa-terraform-aws.pp")

		// Kubernetes - you will need default Kubernetes credentials available or this step will fail
		bridge.GeneratePP(c, kuberneteshandler.Server(c, kubernetes.Provider().(*schema.Provider)), "TerraformKubernetes", "plugins/aaa-terraform-kubernetes.pp")

		// Google - you will need Gcloud credentials available e.g. export GOOGLE_APPLICATION_CREDENTIALS=~/gcp.json
		bridge.GeneratePP(c, googlehandler.Server(c, google.Provider().(*schema.Provider)), "TerraformGoogle", "plugins/aaa-terraform-google.pp")

		// // Azure - you will need login credentials available or this step will fail
		// bridge.GeneratePP(c, azurermhandler.Server(c, azurerm.Provider().(*schema.Provider)), "TerraformAzureRM", "plugins/aaa-terraform-azurerm.pp")

	})
}
