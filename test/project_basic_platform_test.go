package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestProjectPlatformBasicSuccess(t *testing.T) {

	organizationName := "me-urs"
	teamName := "me"

	type args struct {
		organization string
		team         string
		project      string
		platform     string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "validation platform go",
			args: args{
				organization: organizationName,
				team:         teamName,
				project:      "terraform-sentry-go-platform",
				platform:     "go",
			},
		},
		{
			name: "validation platform php",
			args: args{
				organization: organizationName,
				team:         teamName,
				project:      "terraform-sentry-php-platform",
				platform:     "php",
			},
		},
	}
	for _, tt := range tests {
    t.Parallel()
		t.Run(tt.name, func(t *testing.T) {
			terraformOptions := &terraform.Options{
				// The path to where your Terraform code is located
				TerraformDir: "project-basic",
				Upgrade:      true,
				Vars: map[string]interface{}{
					"name":         tt.args.project,
					"organization": tt.args.organization,
					"team":         tt.args.team,
					"platform":     tt.args.platform,
				},
			}

			// At the end of the test, run `terraform destroy` to clean up any resources that were created
			defer terraform.Destroy(t, terraformOptions)

			// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
			terraform.InitAndApply(t, terraformOptions)

			outputID := terraform.Output(t, terraformOptions, "id")
			outputProject := terraform.OutputMap(t, terraformOptions, "project")
			assert.NotEmpty(t, outputProject, outputProject)
			assert.NotEmpty(t, outputID, outputID)
		})
	}
}
