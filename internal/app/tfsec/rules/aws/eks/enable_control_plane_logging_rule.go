package eks

import (
	"github.com/aquasecurity/defsec/rules/aws/eks"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "AWS067",
		BadExample: []string{`
 resource "aws_eks_cluster" "bad_example" {
     encryption_config {
         resources = [ "secrets" ]
         provider {
             key_arn = var.kms_arn
         }
     }
 
     name = "bad_example_cluster"
     role_arn = var.cluster_arn
     vpc_config {
         endpoint_public_access = false
     }
 }
 `},
		GoodExample: []string{`
 resource "aws_eks_cluster" "good_example" {
     encryption_config {
         resources = [ "secrets" ]
         provider {
             key_arn = var.kms_arn
         }
     }
 
 	enabled_cluster_log_types = ["api", "authenticator", "audit", "scheduler", "controllerManager"]
 
     name = "good_example_cluster"
     role_arn = var.cluster_arn
     vpc_config {
         endpoint_public_access = false
     }
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eks_cluster#enabled_cluster_log_types",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_eks_cluster"},
		Base:           eks.CheckEnableControlPlaneLogging,
	})
}
