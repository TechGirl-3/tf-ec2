
package test

import (
	"testing"
	"regexp"
	"log"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../examples",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "instance-key")
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	output = re.ReplaceAllString(output, "")
	assert.Equal(t, "check", output)
}
