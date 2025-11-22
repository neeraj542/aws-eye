package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceDataFields(t *testing.T) {
	instance := InstanceData{
		ID:           "i-02ab34cd56",
		Name:         "opsa-server",
		Type:         "t3.micro",
		State:        "running",
		PublicIP:     "16.171.1.1",
		PrivateIP:    "172.31.1.1",
		AZ:           "eu-north-1b",
		AMI:          "ami-0abcd1234",
		Architecture: "x86_64",
		LaunchTime:   "2025-11-22 15:35:06 UTC",
	}

	assert.Equal(t, "i-02ab34cd56", instance.ID)
	assert.Equal(t, "opsa-server", instance.Name)
	assert.Equal(t, "t3.micro", instance.Type)
	assert.Equal(t, "running", instance.State)
	assert.Equal(t, "16.171.1.1", instance.PublicIP)
	assert.Equal(t, "172.31.1.1", instance.PrivateIP)
	assert.Equal(t, "eu-north-1b", instance.AZ)
	assert.Equal(t, "ami-0abcd1234", instance.AMI)
	assert.Equal(t, "x86_64", instance.Architecture)
	assert.Equal(t, "2025-11-22 15:35:06 UTC", instance.LaunchTime)
}

func TestInstanceDataEmptyFields(t *testing.T) {
	instance := InstanceData{
		ID:    "i-02ab34cd56",
		State: "stopped",
	}

	assert.Equal(t, "i-02ab34cd56", instance.ID)
	assert.Equal(t, "", instance.Name)
	assert.Equal(t, "", instance.PublicIP)
	assert.Equal(t, "", instance.PrivateIP)
}

