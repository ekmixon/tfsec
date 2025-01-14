package cloudfront

import (
	"github.com/aquasecurity/defsec/rules/aws/cloudfront"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "AWS071",
		BadExample: []string{`
 resource "aws_cloudfront_distribution" "bad_example" {
 	// other config
 	// no logging_config
 }
 `},
		GoodExample: []string{`
 resource "aws_cloudfront_distribution" "good_example" {
 	// other config
 	logging_config {
 		include_cookies = false
 		bucket          = "mylogs.s3.amazonaws.com"
 		prefix          = "myprefix"
 	}
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_distribution#logging_config",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_cloudfront_distribution"},
		Base:           cloudfront.CheckEnableLogging,
	})
}
