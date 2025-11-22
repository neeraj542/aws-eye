package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/51xneeraj/aws-eye/internal"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	regionFlag     string
	instanceIDFlag string
	jsonFlag       bool
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe EC2 instances",
	Long: `Describe EC2 instances with interactive prompts or command-line flags.
	
Interactive mode (default):
  aws-eye describe

Flag mode:
  aws-eye describe --region eu-north-1 --instance-id i-0abc1234 --json`,
	Run: runDescribe,
}

func init() {
	rootCmd.AddCommand(describeCmd)

	describeCmd.Flags().StringVarP(&regionFlag, "region", "r", "", "AWS region (e.g., eu-north-1)")
	describeCmd.Flags().StringVarP(&instanceIDFlag, "instance-id", "i", "", "Filter by instance ID")
	describeCmd.Flags().BoolVar(&jsonFlag, "json", false, "Output in JSON format")
}

func runDescribe(cmd *cobra.Command, args []string) {
	var region string
	var instanceID string
	var outputFormat string

	// Determine if we're in interactive mode (no flags provided)
	interactiveMode := regionFlag == "" && instanceIDFlag == "" && !jsonFlag

	if interactiveMode {
		// Interactive mode
		region = promptRegion()

		wantFilter := promptFilter()
		if wantFilter {
			instanceID = promptInstanceID()
		}

		outputFormat = promptOutputFormat()
	} else {
		// Flag mode
		if regionFlag == "" {
			region = "eu-north-1" // default
		} else {
			region = regionFlag
		}
		instanceID = instanceIDFlag

		if jsonFlag {
			outputFormat = "json"
		} else {
			outputFormat = "pretty"
		}
	}

	// Fetch instances
	instances, err := internal.FetchInstances(region, instanceID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching instances: %v\n", err)
		os.Exit(1)
	}

	if len(instances) == 0 {
		fmt.Println("No instances found.")
		return
	}

	// Output based on format
	if outputFormat == "json" {
		outputJSON(instances)
	} else {
		outputPretty(instances)
	}
}

func promptRegion() string {
	region := ""
	prompt := &survey.Input{
		Message: "Enter AWS region (default: eu-north-1):",
		Default: "eu-north-1",
	}
	survey.AskOne(prompt, &region)

	if strings.TrimSpace(region) == "" {
		return "eu-north-1"
	}
	return strings.TrimSpace(region)
}

func promptFilter() bool {
	wantFilter := false
	prompt := &survey.Confirm{
		Message: "Do you want to filter by instance ID?",
		Default: false,
	}
	survey.AskOne(prompt, &wantFilter)
	return wantFilter
}

func promptInstanceID() string {
	instanceID := ""
	prompt := &survey.Input{
		Message: "Enter instance ID:",
	}
	survey.AskOne(prompt, &instanceID)
	return strings.TrimSpace(instanceID)
}

func promptOutputFormat() string {
	choice := ""
	prompt := &survey.Select{
		Message: "Choose output format:",
		Options: []string{"Pretty", "JSON"},
		Default: "Pretty",
	}
	survey.AskOne(prompt, &choice)

	if strings.ToLower(choice) == "json" {
		return "json"
	}
	return "pretty"
}

func outputPretty(instances []internal.InstanceData) {
	for _, instance := range instances {
		fmt.Println(strings.Repeat("-", 50))
		fmt.Printf("Instance: %s\n", instance.ID)

		if instance.Name != "" {
			fmt.Printf("Name: %s\n", instance.Name)
		}

		fmt.Printf("Type: %s\n", instance.Type)

		// Color-coded state
		stateColor := color.New(color.FgYellow)
		if instance.State == "running" {
			stateColor = color.New(color.FgGreen)
		}
		stateColor.Printf("State: %s\n", strings.ToUpper(instance.State))

		if instance.PublicIP != "" {
			fmt.Printf("Public IP: %s\n", instance.PublicIP)
		} else {
			fmt.Printf("Public IP: —\n")
		}

		if instance.PrivateIP != "" {
			fmt.Printf("Private IP: %s\n", instance.PrivateIP)
		} else {
			fmt.Printf("Private IP: —\n")
		}

		fmt.Printf("AZ: %s\n", instance.AZ)
		fmt.Printf("AMI: %s\n", instance.AMI)
		fmt.Printf("Architecture: %s\n", instance.Architecture)
		fmt.Printf("Launched: %s\n", instance.LaunchTime)
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println()
	}
}

func outputJSON(instances []internal.InstanceData) {
	jsonData, err := json.MarshalIndent(instances, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
