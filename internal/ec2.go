package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// InstanceData represents EC2 instance information
type InstanceData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	State       string `json:"state"`
	PublicIP    string `json:"public_ip"`
	PrivateIP   string `json:"private_ip"`
	AZ          string `json:"availability_zone"`
	AMI         string `json:"ami_id"`
	Architecture string `json:"architecture"`
	LaunchTime  string `json:"launch_time"`
}

// FetchInstances retrieves EC2 instances from the specified region
// If instanceID is provided, it filters to that specific instance
func FetchInstances(region string, instanceID string) ([]InstanceData, error) {
	client, err := GetEC2Client(region)
	if err != nil {
		return nil, err
	}

	// Build the input
	input := &ec2.DescribeInstancesInput{}

	// If instance ID is provided, filter by it
	if instanceID != "" {
		input.InstanceIds = []string{instanceID}
	}

	// Describe instances
	result, err := client.DescribeInstances(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe instances: %w", err)
	}

	var instances []InstanceData

	// Extract instance data from reservations
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instanceData := parseInstance(instance)
			instances = append(instances, instanceData)
		}
	}

	return instances, nil
}

// parseInstance converts an EC2 instance to InstanceData
func parseInstance(instance types.Instance) InstanceData {
	data := InstanceData{
		ID:          aws.ToString(instance.InstanceId),
		Type:        string(instance.InstanceType),
		State:       string(instance.State.Name),
		AZ:          aws.ToString(instance.Placement.AvailabilityZone),
		AMI:         aws.ToString(instance.ImageId),
		Architecture: string(instance.Architecture),
	}

	// Get public IP
	if instance.PublicIpAddress != nil {
		data.PublicIP = aws.ToString(instance.PublicIpAddress)
	}

	// Get private IP
	if instance.PrivateIpAddress != nil {
		data.PrivateIP = aws.ToString(instance.PrivateIpAddress)
	}

	// Get name tag
	for _, tag := range instance.Tags {
		if aws.ToString(tag.Key) == "Name" {
			data.Name = aws.ToString(tag.Value)
			break
		}
	}

	// Format launch time
	if instance.LaunchTime != nil {
		data.LaunchTime = instance.LaunchTime.Format("2006-01-02 15:04:05 UTC")
	}

	return data
}

// FormatLaunchTime formats a time.Time to the required string format
func FormatLaunchTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05 UTC")
}

