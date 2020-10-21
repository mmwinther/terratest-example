package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/terraform-hello-world-example",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "hello_world")

	assert.Equal(t, "Hello, World!", output)
	
}