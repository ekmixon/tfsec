package compute

import (
	"testing"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/testutil"
)

func Test_GoogleOpenInboundFirewallRule(t *testing.T) {
	expectedCode := "google-compute-no-public-ingress"

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode string
		mustExcludeResultCode string
	}{
		{
			name: "check google_compute_firewall ingress on 0.0.0.0/0",
			source: `
 resource "google_compute_firewall" "my-firewall" {
    allow {
        protocol = "tcp"
    }
 	source_ranges = ["0.0.0.0/0"]
 }`,
			mustIncludeResultCode: expectedCode,
		},
		{
			name: "check google_compute_firewall ingress on default sources",
			source: `
 resource "google_compute_firewall" "my-firewall" {
    allow {
        protocol = "tcp"
    }
 	source_ranges = []
 }`,
			mustIncludeResultCode: expectedCode,
		},
		{
			name: "check google_compute_firewall ingress on /32",
			source: `
 resource "google_compute_firewall" "my-firewall" {
    allow {
        protocol = "tcp"
    }
 	source_ranges = ["127.0.0.1/32"]
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
