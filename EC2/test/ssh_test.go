package test

import (
        "testing"
	//"strings"
        "fmt"
	"github.com/gruntwork-io/terratest/modules/aws"
        "github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsExample(t *testing.T) {
        t.Parallel()
        awsRegion := "us-east-1"
        expectedKey := "gotest"
        terraformOptions := &terraform.Options{
                TerraformDir: "../examples",
                // Environment variables to set when running Terraform
                EnvVars: map[string]string{
                        "AWS_DEFAULT_REGION": awsRegion,
			"GOCACHE" : "off",
                },
        }
        defer terraform.Destroy(t, terraformOptions)
        terraform.InitAndApply(t, terraformOptions)
        instanceKey := terraform.Output(t, terraformOptions, "instance-key")
        //fmt.Printf("%s\n", instanceKey)
	//instanceKey = strings.Replace(instanceKey, "\n", "", -1)
	assert.Equal(t, expectedKey, instanceKey)

	//Tags
	instanceID := terraform.Output(t, terraformOptions, "instance_id")
	//aws.AddTagsToResource(t, awsRegion, instanceID, map[string]string{"Name": "testing-tag-value"})
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)
	fmt.Printf("%s Tags \n", instanceTags)
	
	nameTag, containsNameTag := instanceTags["Name"]
	assert.True(t, containsNameTag)
	assert.Equal(t, "testtera", nameTag)
}

