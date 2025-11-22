package internal

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/stretchr/testify/assert"
)

func TestFormatLaunchTime(t *testing.T) {
	tests := []struct {
		name     string
		input    *time.Time
		expected string
	}{
		{
			name:     "nil time",
			input:    nil,
			expected: "",
		},
		{
			name:     "valid time",
			input:    timePtr(time.Date(2025, 11, 22, 15, 35, 6, 0, time.UTC)),
			expected: "2025-11-22 15:35:06 UTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatLaunchTime(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseInstance(t *testing.T) {
	launchTime := time.Date(2025, 11, 22, 15, 35, 6, 0, time.UTC)

	instance := types.Instance{
		InstanceId:   stringPtr("i-02ab34cd56"),
		InstanceType: types.InstanceTypeT3Micro,
		State: &types.InstanceState{
			Name: types.InstanceStateNameRunning,
		},
		PublicIpAddress:  stringPtr("16.171.1.1"),
		PrivateIpAddress: stringPtr("172.31.1.1"),
		Placement: &types.Placement{
			AvailabilityZone: stringPtr("eu-north-1b"),
		},
		ImageId:      stringPtr("ami-0abcd1234"),
		Architecture: types.ArchitectureValuesX8664,
		LaunchTime:   &launchTime,
		Tags: []types.Tag{
			{
				Key:   stringPtr("Name"),
				Value: stringPtr("opsa-server"),
			},
		},
	}

	result := parseInstance(instance)

	assert.Equal(t, "i-02ab34cd56", result.ID)
	assert.Equal(t, "opsa-server", result.Name)
	assert.Equal(t, "t3.micro", result.Type)
	assert.Equal(t, "running", result.State)
	assert.Equal(t, "16.171.1.1", result.PublicIP)
	assert.Equal(t, "172.31.1.1", result.PrivateIP)
	assert.Equal(t, "eu-north-1b", result.AZ)
	assert.Equal(t, "ami-0abcd1234", result.AMI)
	assert.Equal(t, "x86_64", result.Architecture)
	assert.Equal(t, "2025-11-22 15:35:06 UTC", result.LaunchTime)
}

func TestParseInstanceWithoutTags(t *testing.T) {
	instance := types.Instance{
		InstanceId:   stringPtr("i-02ab34cd56"),
		InstanceType: types.InstanceTypeT3Micro,
		State: &types.InstanceState{
			Name: types.InstanceStateNameStopped,
		},
		Placement: &types.Placement{
			AvailabilityZone: stringPtr("eu-north-1b"),
		},
		ImageId:      stringPtr("ami-0abcd1234"),
		Architecture: types.ArchitectureValuesX8664,
	}

	result := parseInstance(instance)

	assert.Equal(t, "i-02ab34cd56", result.ID)
	assert.Equal(t, "", result.Name) // No name tag
	assert.Equal(t, "stopped", result.State)
	assert.Equal(t, "", result.PublicIP)  // No public IP
	assert.Equal(t, "", result.PrivateIP) // No private IP
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func timePtr(t time.Time) *time.Time {
	return &t
}
