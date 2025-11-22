# aws-eye

A lightweight, interactive AWS EC2 utility that fetches instance details in a clean, readable format. Think of it as a user-friendly, interactive version of `aws ec2 describe-instances`.

## ✨ Features

- 🎯 **Interactive Mode**: Step-by-step prompts for easy usage
- 🚀 **Flag Mode**: Quick command-line flags for automation
- 🎨 **Color-coded Output**: Visual state indicators (green for running, yellow for stopped)
- 📊 **Multiple Formats**: Pretty-printed or JSON output
- 🔍 **Instance Filtering**: Filter by instance ID
- 📋 **Comprehensive Info**: Instance ID, Name, Type, State, IPs, AZ, AMI, Architecture, Launch Time

## 📦 Installation

### Prerequisites

- Go 1.21 or later
- AWS credentials configured (via environment variables or `~/.aws/credentials`)

### Build from Source

```bash
# Clone or navigate to the project
cd aws-eye

# Install dependencies
make install
# or
go mod download

# Build the binary
make build
# or
go build -o aws-eye .
```

## 🚀 Usage

### Interactive Mode (Default)

Simply run the `describe` command without flags:

```bash
./aws-eye describe
```

The CLI will prompt you for:
1. AWS region (default: eu-north-1)
2. Whether to filter by instance ID (y/n)
3. Output format (Pretty/JSON)

**Example Interactive Session:**
```
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
Public IP: 16.171.xx.xx
Private IP: 172.31.xx.xx
AZ: eu-north-1b
AMI: ami-0abcd1234
Architecture: x86_64
Launched: 2025-11-22 15:35:06 UTC
--------------------------------------------------
```

### Flag Mode

Use command-line flags for quick, non-interactive usage:

```bash
# Basic usage with region
./aws-eye describe --region eu-north-1

# Filter by instance ID
./aws-eye describe --region eu-north-1 --instance-id i-0abc1234

# JSON output
./aws-eye describe --region eu-north-1 --json

# Combined flags
./aws-eye describe --region us-east-1 --instance-id i-0abc1234 --json
```

**Available Flags:**
- `--region, -r`: AWS region (e.g., eu-north-1, us-east-1)
- `--instance-id, -i`: Filter by specific instance ID
- `--json`: Output in JSON format instead of pretty-printed

## 📋 Output Format

### Pretty Format

The default pretty format displays instance information in a clean, readable layout:

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

Use `--json` flag or select JSON in interactive mode:

```bash
./aws-eye describe --json
```

**Example JSON Output:**
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

## 🛠️ Development

### Makefile Commands

```bash
# Build the application
make build

# Run the application (interactive mode)
make run

# Format the code
make format

# Run tests
make test

# Run tests with coverage
make test-coverage

# Clean build artifacts
make clean

# Install dependencies
make install
```

### Project Structure

```
aws-eye/
├── cmd/
│   ├── root.go          # Root command setup
│   └── describe.go      # Describe command (interactive + flags)
├── internal/
│   ├── awsclient.go     # AWS client initialization
│   ├── ec2.go           # EC2 fetching and parsing
│   ├── ec2_test.go      # EC2 tests
│   └── formatter_test.go # Formatter tests
├── main.go              # Entry point
├── Makefile             # Build automation
├── go.mod               # Go dependencies
└── README.md            # This file
```

## 🧪 Testing

Run the test suite:

```bash
make test
```

Or with coverage:

```bash
make test-coverage
```

This will generate a `coverage.html` file you can open in your browser.

## 🔐 AWS Credentials Setup

### Option 1: Environment Variables

```bash
export AWS_ACCESS_KEY_ID="your-access-key-id"
export AWS_SECRET_ACCESS_KEY="your-secret-access-key"
export AWS_DEFAULT_REGION="eu-north-1"  # Optional
```

### Option 2: AWS Credentials File

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

### Required IAM Permissions

Your AWS user needs the following permission:
- `ec2:DescribeInstances`

The easiest way is to attach the managed policy: `AmazonEC2ReadOnlyAccess`

## 📝 Examples

### Example 1: List all instances in a region

```bash
./aws-eye describe --region eu-north-1
```

### Example 2: Get specific instance details

```bash
./aws-eye describe --region eu-north-1 --instance-id i-02ab34cd56
```

### Example 3: Get JSON output for scripting

```bash
./aws-eye describe --region eu-north-1 --json | jq '.[0].public_ip'
```

### Example 4: Interactive mode with filtering

```bash
./aws-eye describe
# Follow the prompts to filter by instance ID
```

## 🐛 Troubleshooting

### Error: "failed to load AWS config"

**Solution**: Ensure your AWS credentials are configured correctly. Check:
- Environment variables are set, or
- `~/.aws/credentials` file exists and is readable

### Error: "UnauthorizedOperation"

**Solution**: Your IAM user needs `ec2:DescribeInstances` permission. Attach the `AmazonEC2ReadOnlyAccess` policy to your user.

### Error: "No instances found"

**Solution**: This is normal if you don't have EC2 instances in the specified region. Try a different region or check your AWS console.

### Colors not showing

**Solution**: Some terminals don't support colors. The functionality still works, just without color coding.

## 📄 License

This project is open source and available for use.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📚 Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Survey](https://github.com/AlecAivazis/survey) - Interactive prompts
- [Color](https://github.com/fatih/color) - Terminal colors
- [AWS SDK Go v2](https://github.com/aws/aws-sdk-go-v2) - AWS API client

---

**Made with ❤️ for AWS users who want a simpler way to query EC2 instances**

