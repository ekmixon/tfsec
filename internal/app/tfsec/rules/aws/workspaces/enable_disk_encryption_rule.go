package workspaces

import (
	"github.com/aquasecurity/defsec/rules/aws/workspaces"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "AWS084",
		BadExample: []string{`
 resource "aws_workspaces_workspace" "bad_example" {
 	directory_id = aws_workspaces_directory.test.id
 	bundle_id    = data.aws_workspaces_bundle.value_windows_10.id
 	user_name    = "Administrator"
   
 	workspace_properties {
 	  compute_type_name                         = "VALUE"
 	  user_volume_size_gib                      = 10
 	  root_volume_size_gib                      = 80
 	  running_mode                              = "AUTO_STOP"
 	  running_mode_auto_stop_timeout_in_minutes = 60
 	}
   }
 `},
		GoodExample: []string{`	
 resource "aws_workspaces_workspace" "good_example" {
 		directory_id 				   = aws_workspaces_directory.test.id
 		bundle_id    				   = data.aws_workspaces_bundle.value_windows_10.id
 		user_name    				   = "Administrator"
 		root_volume_encryption_enabled = true
 		user_volume_encryption_enabled = true
 	  
 		workspace_properties {
 		  compute_type_name                         = "VALUE"
 		  user_volume_size_gib                      = 10
 		  root_volume_size_gib                      = 80
 		  running_mode                              = "AUTO_STOP"
 		  running_mode_auto_stop_timeout_in_minutes = 60
 		}
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/workspaces_workspace#root_volume_encryption_enabled",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_workspaces_workspace"},
		Base:           workspaces.CheckEnableDiskEncryption,
	})
}
