# Changelog

All notable changes to aws-eye will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-01-XX

### Added
- Initial release of aws-eye
- Interactive mode with step-by-step prompts
- Flag mode for command-line automation
- Color-coded output (green for running, yellow for stopped)
- JSON output format support
- Instance filtering by instance ID
- Comprehensive EC2 instance information:
  - Instance ID
  - Name tag
  - Instance type
  - State
  - Public and private IP addresses
  - Availability Zone
  - AMI ID
  - Architecture
  - Launch time
- Cobra CLI framework integration
- Survey library for interactive prompts
- Unit tests for EC2 parsing and formatting
- Makefile with build, test, and format targets
- Comprehensive documentation (README, setup guides)

### Features
- Support for multiple AWS regions
- Default region: eu-north-1
- Pretty-printed and JSON output formats
- Error handling with clear error messages
- AWS credentials support via environment variables or credentials file

---

## [Unreleased]

### Planned
- Support for filtering by tags
- Support for filtering by instance state
- Export to CSV format
- Support for multiple regions in one query
- Caching of instance data
- Configuration file support

---

## Version History

- **1.0.0** - Initial release with interactive and flag modes

---

[1.0.0]: https://github.com/51xneeraj/aws-eye/releases/tag/v1.0.0

