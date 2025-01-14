package elb

import (
	"testing"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/testutil"
)

func Test_AWSOutdatedSSLPolicy(t *testing.T) {
	expectedCode := "aws-elb-use-secure-tls-policy"

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode string
		mustExcludeResultCode string
	}{
		{
			name: "check aws_alb_listener with outdated policy",
			source: `
 resource "aws_alb" "front_end" {
 }

 resource "aws_alb_listener" "my-resource" {
	load_balancer_arn = aws_alb.front_end.arn
 	ssl_policy = "ELBSecurityPolicy-TLS-1-1-2017-01"
 	protocol = "HTTPS"
 }`,
			mustIncludeResultCode: expectedCode,
		},
		{
			name: "check aws_lb_listener with outdated policy",
			source: `
 resource "aws_lb" "front_end" {
 }

 resource "aws_lb_listener" "my-resource" {
	load_balancer_arn = aws_lb.front_end.arn
 	ssl_policy = "ELBSecurityPolicy-TLS-1-1-2017-01"
 	protocol = "HTTPS"
 }`,
			mustIncludeResultCode: expectedCode,
		},
		{
			name: "check aws_alb_listener with ok policy",
			source: `
 resource "aws_alb" "front_end" {
 }

 resource "aws_alb_listener" "my-resource" {
	load_balancer_arn = aws_alb.front_end.arn
 	ssl_policy = "ELBSecurityPolicy-TLS-1-2-2017-01"
 	protocol = "HTTPS"
 }`,
			mustExcludeResultCode: expectedCode,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			results := testutil.ScanHCL(test.source, t)
			testutil.AssertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}
