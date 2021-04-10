package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestProjectBasicSuccess(t *testing.T) {
	t.Parallel()

	const NAME string = "test-project-sentry"
	const ORGANIZATION string = "me-urs"
	const TEAM string = "me"
  const PLATFORM string = "go"

	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "project-basic",
		Upgrade:      true,
		Vars: map[string]interface{}{
			"name":  NAME,
			"organization": ORGANIZATION,
			"team": TEAM,
      "platform": PLATFORM,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	outputId := terraform.Output(t, terraformOptions, "id")
	outputProject := terraform.OutputMap(t, terraformOptions, "project")
	assert.NotEmpty(t, outputProject, outputProject)
	assert.NotEmpty(t, outputId, outputId)
}
