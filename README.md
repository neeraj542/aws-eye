# aws-eye 🟧

[![Go Reference](https://pkg.go.dev/badge/github.com/neeraj542/aws-eye.svg)](https://pkg.go.dev/github.com/neeraj542/aws-eye)
![GitHub Release](https://img.shields.io/github/v/release/neeraj542/aws-eye)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
[![Documentation](https://img.shields.io/badge/docs-available-blue)](DOCUMENTATION.md)
[![MIT License](https://img.shields.io/badge/MIT-License-green)](LICENSE)
[![Contributing](https://img.shields.io/badge/contributions-welcome-brightgreen)](CONTRIBUTING.md)


![aws-eye (2)](https://github.com/user-attachments/assets/19257070-14dd-47c9-a808-4460c6f02b24)


A lightweight, interactive AWS EC2 utility that fetches instance details in a clean, readable format. Think of it as a user-friendly, interactive version of `aws ec2 describe-instances`.

## Overview

aws-eye is a command-line tool that provides an intuitive interface for querying AWS EC2 instances. It offers both interactive prompts and command-line flags, making it suitable for both beginners and automation scripts.

### Key Features

- **Interactive Mode**: Step-by-step prompts for easy usage
- **Flag Mode**: Quick command-line flags for automation
- **Color-coded Output**: Visual state indicators (green for running, yellow for stopped)
- **Multiple Formats**: Pretty-printed or JSON output
- **Instance Filtering**: Filter by instance ID
- **Comprehensive Information**: Instance ID, Name, Type, State, IPs, AZ, AMI, Architecture, Launch Time

## Installation

### Prerequisites

- Go 1.21 or later 
- AWS account with EC2 access
- AWS credentials configured (see [Documentation](DOCUMENTATION.md#aws-credentials-setup))

### Build from Source

```bash
# Clone the repository
git clone https://github.com/neeraj542/aws-eye.git
cd aws-eye

# Install dependencies
make install

# Build the binary
make build
```

### Install Globally (Optional)

```bash
sudo cp aws-eye /usr/local/bin/
# or
sudo ln -s $(pwd)/aws-eye /usr/local/bin/aws-eye
```

## Quick Start

### Interactive Mode

Run the command without flags to enter interactive mode:

```bash
./aws-eye describe
```

The CLI will prompt you for:
1. AWS region (default: eu-north-1)
2. Whether to filter by instance ID
3. Output format (Pretty/JSON)

**Example Interactive Session:**

<img width="872" height="240" alt="Interactive Session Screenshot" src="https://github.com/user-attachments/assets/afdd4b53-2505-4ea3-aac3-7b5466d34f2a" />

### Flag Mode

Use command-line flags for quick, non-interactive usage:

```bash
# List all instances in a region
./aws-eye describe --region eu-north-1

# Filter by instance ID
./aws-eye describe --region eu-north-1 --instance-id i-0abc1234

# JSON output
./aws-eye describe --region eu-north-1 --json
```

**Available Flags:**
- `--region, -r`: AWS region (e.g., eu-north-1, us-east-1)
- `--instance-id, -i`: Filter by specific instance ID
- `--json`: Output in JSON format instead of pretty-printed

## Output Format

### Pretty Format (Default)

The default format displays instance information in a clean, readable layout:

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
- Green: Running instances
- Yellow: Stopped instances

### JSON Format

Use the `--json` flag for machine-readable output:

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

## Development

### Makefile Commands

```bash
make build          # Build the application
make test           # Run tests
make format         # Format code
make run            # Run in interactive mode
make clean          # Clean build artifacts
make install        # Install dependencies
make test-coverage  # Run tests with coverage
```

### Project Structure

```
aws-eye/
├── cmd/              # CLI commands
│   ├── root.go       # Root command setup
│   └── describe.go   # Describe command
├── internal/         # Internal packages
│   ├── awsclient.go  # AWS client initialization
│   ├── ec2.go        # EC2 fetching and parsing
│   └── *_test.go     # Test files
├── main.go           # Entry point
├── Makefile          # Build automation
└── README.md         # This file
```

## Documentation

For detailed documentation, examples, troubleshooting, and API reference, see [DOCUMENTATION.md](DOCUMENTATION.md).

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Survey](https://github.com/AlecAivazis/survey) - Interactive prompts
- [Color](https://github.com/fatih/color) - Terminal colors
- [AWS SDK Go v2](https://github.com/aws/aws-sdk-go-v2) - AWS API client

---

Made with ❤️ for AWS users who want a simpler way to query EC2 instances
