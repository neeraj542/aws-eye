# aws-eye Documentation

Complete documentation for the aws-eye CLI tool with detailed steps and navigation.

## Table of Contents

1. [Overview](#overview)
2. [Installation](#installation)
3. [AWS Credentials Setup](#aws-credentials-setup)
4. [Usage](#usage)
5. [Commands](#commands)
6. [Output Formats](#output-formats)
7. [Examples](#examples)
8. [Troubleshooting](#troubleshooting)
9. [API Reference](#api-reference)
10. [Development](#development)

---

## Overview

aws-eye is a lightweight, interactive AWS EC2 utility that provides a user-friendly way to query and display EC2 instance information. It offers both interactive and programmatic interfaces for fetching instance details.

### Key Features

- **Interactive Mode**: Guided prompts for easy usage
- **Flag Mode**: Command-line flags for automation and scripting
- **Color Output**: Visual indicators for instance states
- **Multiple Formats**: Pretty-printed and JSON output
- **Filtering**: Filter instances by ID
- **Comprehensive Data**: All essential EC2 instance information

---

## Installation

### Prerequisites

- **Go**: Version 1.21 or later
- **AWS Account**: With EC2 access
- **AWS Credentials**: Configured locally

### Build from Source

```bash
# Clone the repository
git clone https://github.com/51xneeraj/aws-eye.git
cd aws-eye

# Install dependencies
go mod download

# Build the binary
go build -o aws-eye .

# Or use Makefile
make build
```

### Install Globally (Optional)

```bash
# Add to PATH or create symlink
sudo cp aws-eye /usr/local/bin/
# or
sudo ln -s $(pwd)/aws-eye /usr/local/bin/aws-eye
```

---

## AWS Credentials Setup

### What You Need

1. **AWS Account** (Free Tier is Sufficient)
   - You can use a free tier AWS account
   - No payment required for basic EC2 instance queries
   - The `DescribeInstances` API call is **FREE** (it only reads metadata, doesn't create resources)

2. **IAM User with EC2 Read Permissions**
   - You need an IAM user with permission to describe EC2 instances

3. **AWS Credentials**
   - Access Key ID and Secret Access Key for authentication

### Step-by-Step Setup

#### Step 1: Create an IAM User

1. **Log in to AWS Console**
   - Go to https://console.aws.amazon.com
   - Sign in with your AWS account

2. **Navigate to IAM**
   - Search for "IAM" in the top search bar
   - Click on "IAM" service

3. **Create a New User**
   - Click "Users" in the left sidebar
   - Click "Create user"
   - Enter a username (e.g., `aws-eye-user`)
   - Click "Next"

4. **Attach Permissions**
   - Select "Attach policies directly"
   - Search for and select: **`AmazonEC2ReadOnlyAccess`**
     - This policy allows reading EC2 instance information (no charges)
   - Click "Next", then "Create user"

#### Step 2: Create Access Keys

1. **Select Your User**
   - Click on the user you just created

2. **Create Access Key**
   - Go to the "Security credentials" tab
   - Scroll to "Access keys" section
   - Click "Create access key"
   - Select "Command Line Interface (CLI)" as the use case
   - Check the confirmation box and click "Next"
   - Optionally add a description tag, then click "Create access key"

3. **Save Your Credentials** ⚠️ **IMPORTANT**
   - **Access Key ID**: Copy and save this immediately
   - **Secret Access Key**: Copy and save this immediately
   - ⚠️ **You can only see the Secret Access Key once!** If you lose it, you'll need to create a new access key.

#### Step 3: Configure AWS Credentials Locally

aws-eye uses the AWS SDK's default credential chain. It will look for credentials in this order:

1. Environment variables
2. Shared credentials file (`~/.aws/credentials`)
3. Shared config file (`~/.aws/config`)
4. IAM roles (if running on EC2)

**Option A: Environment Variables (Recommended for Testing)**

```bash
export AWS_ACCESS_KEY_ID="your-access-key-id"
export AWS_SECRET_ACCESS_KEY="your-secret-access-key"
export AWS_DEFAULT_REGION="eu-north-1"  # Optional
```

**For macOS/Linux:**
Add these to your `~/.zshrc` or `~/.bashrc`:
```bash
echo 'export AWS_ACCESS_KEY_ID="your-access-key-id"' >> ~/.zshrc
echo 'export AWS_SECRET_ACCESS_KEY="your-secret-access-key"' >> ~/.zshrc
source ~/.zshrc
```

**Option B: AWS Credentials File (Recommended for Production)**

Create `~/.aws/credentials`:
```ini
[default]
aws_access_key_id = your-access-key-id
aws_secret_access_key = your-secret-access-key
```

Create `~/.aws/config`:
```ini
[default]
region = eu-north-1
```

Set proper permissions:
```bash
chmod 600 ~/.aws/credentials
```

### Required IAM Permissions

Your AWS user/role needs the following permission:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:DescribeInstances"
      ],
      "Resource": "*"
    }
  ]
}
```

Or attach the managed policy: `AmazonEC2ReadOnlyAccess`

### Free Tier Information

**What's FREE:**
- ✅ **EC2 DescribeInstances API calls** - Completely free (you're just reading metadata)
- ✅ **IAM users and policies** - Free
- ✅ **EC2 instance metadata queries** - Free

**What Costs Money:**
- ❌ **Running EC2 instances** - Free tier includes 750 hours/month of t2.micro/t3.micro for 12 months
- ❌ **Stopped instances** - No charges (but they still exist and can be queried)

**Important Notes:**
- **You can query instances even if you have none** - The API call will return an empty list
- **You can query instances in any region** - No charges for API calls
- **The CLI only READS data** - It never creates, modifies, or deletes anything

### Security Best Practices

1. **Never commit credentials to Git**
   - The `.gitignore` should exclude credential files
   - Use environment variables or AWS credentials file

2. **Use least privilege**
   - Only grant `AmazonEC2ReadOnlyAccess` (not full EC2 access)

3. **Rotate keys regularly**
   - Delete old access keys when creating new ones

4. **Use IAM roles for production**
   - For EC2 instances or Lambda, use IAM roles instead of access keys

---

## Usage

### Basic Usage

```bash
# Interactive mode
./aws-eye describe

# Flag mode
./aws-eye describe --region eu-north-1
```

### Getting Help

```bash
# General help
./aws-eye --help

# Command-specific help
./aws-eye describe --help

# Version
./aws-eye --version
```

---

## Commands

### `describe`

Describe EC2 instances in a region.

#### Interactive Mode

When run without flags, aws-eye enters interactive mode:

```bash
./aws-eye describe
```

Prompts:
1. **Region**: AWS region (default: eu-north-1)
2. **Filter by Instance ID**: Yes/No
3. **Instance ID**: If filtering (optional)
4. **Output Format**: Pretty or JSON

#### Flag Mode

```bash
./aws-eye describe [flags]
```

**Flags:**
- `--region, -r string`: AWS region (e.g., eu-north-1, us-east-1)
- `--instance-id, -i string`: Filter by specific instance ID
- `--json`: Output in JSON format instead of pretty-printed
- `--help, -h`: Show help for describe command

**Examples:**
```bash
# List all instances in a region
./aws-eye describe --region eu-north-1

# Get specific instance
./aws-eye describe --region eu-north-1 --instance-id i-0abc1234

# JSON output
./aws-eye describe --region eu-north-1 --json

# Combined flags
./aws-eye describe -r us-east-1 -i i-0abc1234 --json
```

---

## Output Formats

### Pretty Format (Default)

Human-readable format with color coding:

```
--------------------------------------------------
Instance: i-02ab34cd56
Name: opsa-server
Type: t3.micro
State: RUNNING
Public IP: 16.171.xx.xx
Private IP: 172.31.xx.xx
AZ: eu-north-1b
AMI: ami-0abcd1234
Architecture: x86_64
Launched: 2025-11-22 15:35:06 UTC
--------------------------------------------------
```

**State Colors:**
- 🟢 **Green**: Running instances
- 🟡 **Yellow**: Stopped instances

### JSON Format

Machine-readable JSON output:

```json
[
  {
    "id": "i-02ab34cd56",
    "name": "opsa-server",
    "type": "t3.micro",
    "state": "running",
    "public_ip": "16.171.xx.xx",
    "private_ip": "172.31.xx.xx",
    "availability_zone": "eu-north-1b",
    "ami_id": "ami-0abcd1234",
    "architecture": "x86_64",
    "launch_time": "2025-11-22 15:35:06 UTC"
  }
]
```

---

## Examples

### Example 1: Interactive Session

```bash
$ ./aws-eye describe
? Enter AWS region (default: eu-north-1): eu-north-1
? Do you want to filter by instance ID? No
? Choose output format: Pretty
  Pretty
  JSON
> Pretty

--------------------------------------------------
Instance: i-02ab34cd56
Name: opsa-server
Type: t3.micro
State: RUNNING
...
--------------------------------------------------
```

### Example 2: List All Instances

```bash
./aws-eye describe --region eu-north-1
```

### Example 3: Get Specific Instance

```bash
./aws-eye describe --region eu-north-1 --instance-id i-02ab34cd56
```

### Example 4: JSON Output for Scripting

```bash
# Get public IP of first instance
./aws-eye describe --region eu-north-1 --json | jq '.[0].public_ip'

# Count running instances
./aws-eye describe --region eu-north-1 --json | jq '[.[] | select(.state == "running")] | length'
```

### Example 5: Multiple Regions

```bash
# Check instances in multiple regions
for region in eu-north-1 us-east-1 us-west-2; do
  echo "=== $region ==="
  ./aws-eye describe --region $region
done
```

---

## Troubleshooting

### Common Issues

#### Error: "failed to load AWS config"

**Cause**: AWS credentials not configured.

**Solution**:
- Set environment variables, or
- Configure `~/.aws/credentials` file
- Verify credentials are correct

#### Error: "UnauthorizedOperation"

**Cause**: IAM user lacks required permissions.

**Solution**:
- Attach `AmazonEC2ReadOnlyAccess` policy to your IAM user
- Or grant `ec2:DescribeInstances` permission

#### Error: "No instances found"

**Cause**: No EC2 instances in the specified region.

**Solution**:
- This is normal if you have no instances
- Try a different region
- Verify instances exist in AWS Console

#### Colors Not Showing

**Cause**: Terminal doesn't support colors.

**Solution**:
- Functionality still works without colors
- Use `--json` flag for consistent output
- Check terminal color support

#### "command not found: aws-eye"

**Cause**: Binary not in PATH or not built.

**Solution**:
```bash
# Build the binary
make build

# Use with path
./aws-eye describe

# Or install globally
sudo cp aws-eye /usr/local/bin/
```

---

## API Reference

### Internal Packages

#### `internal` Package

##### `FetchInstances(region string, instanceID string) ([]InstanceData, error)`

Fetches EC2 instances from the specified region.

**Parameters:**
- `region` (string): AWS region (e.g., "eu-north-1")
- `instanceID` (string): Optional instance ID to filter by (empty string for all)

**Returns:**
- `[]InstanceData`: Slice of instance data
- `error`: Error if fetch fails

**Example:**
```go
instances, err := internal.FetchInstances("eu-north-1", "")
if err != nil {
    log.Fatal(err)
}
```

##### `InstanceData` Struct

```go
type InstanceData struct {
    ID           string `json:"id"`
    Name         string `json:"name"`
    Type         string `json:"type"`
    State        string `json:"state"`
    PublicIP     string `json:"public_ip"`
    PrivateIP    string `json:"private_ip"`
    AZ           string `json:"availability_zone"`
    AMI          string `json:"ami_id"`
    Architecture string `json:"architecture"`
    LaunchTime   string `json:"launch_time"`
}
```

##### `GetEC2Client(region string) (*ec2.Client, error)`

Creates and returns an EC2 client for the specified region.

**Parameters:**
- `region` (string): AWS region

**Returns:**
- `*ec2.Client`: EC2 client instance
- `error`: Error if client creation fails

---

## Development

### Building

```bash
make build
```

### Testing

```bash
make test
make test-coverage
```

### Formatting

```bash
make format
```

### Project Structure

```
aws-eye/
├── cmd/              # CLI commands
├── internal/         # Internal packages
├── main.go           # Entry point
├── Makefile          # Build automation
└── README.md         # Documentation
```

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/51xneeraj/aws-eye/issues)
- **Documentation**: See [README.md](README.md) for quick start
- **Contributing**: See [CONTRIBUTING.md](CONTRIBUTING.md)

---

**Last Updated**: 2025-01-XX
