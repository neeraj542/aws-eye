# Contributing to aws-eye

Thank you for your interest in contributing to aws-eye! This document provides guidelines and instructions for contributing.

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers and help them learn
- Focus on constructive feedback

## How to Contribute

### Reporting Bugs

If you find a bug, please open an issue with:

- **Description**: Clear description of the bug
- **Steps to Reproduce**: Step-by-step instructions
- **Expected Behavior**: What should happen
- **Actual Behavior**: What actually happens
- **Environment**: OS, Go version, AWS SDK version
- **Error Messages**: Full error output if applicable

### Suggesting Features

Feature suggestions are welcome! Please open an issue with:

- **Use Case**: Why this feature would be useful
- **Proposed Solution**: How you envision it working
- **Alternatives**: Other approaches you've considered

### Submitting Pull Requests

1. **Fork the repository**

2. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes**
   - Follow the existing code style
   - Add tests for new functionality
   - Update documentation as needed
   - Ensure all tests pass: `make test`

4. **Commit your changes**
   ```bash
   git commit -m "Add: description of your changes"
   ```
   Use clear, descriptive commit messages.

5. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```

6. **Open a Pull Request**
   - Provide a clear description of changes
   - Reference any related issues
   - Ensure CI checks pass

## Development Setup

### Prerequisites

- Go 1.21 or later
- AWS credentials configured (for testing)

### Setup

```bash
# Clone your fork
git clone https://github.com/your-username/aws-eye.git
cd aws-eye

# Install dependencies
make install

# Run tests
make test

# Build
make build
```

## Code Style

- Follow Go conventions and best practices
- Use `gofmt` or `go fmt` to format code
- Run `make format` before committing
- Keep functions focused and small
- Add comments for exported functions and types
- Write meaningful variable and function names

## Testing

- Add tests for new features
- Ensure existing tests still pass
- Aim for good test coverage
- Test both success and error cases

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

## Project Structure

```
aws-eye/
├── cmd/              # CLI commands
├── internal/         # Internal packages
├── main.go           # Entry point
├── Makefile          # Build automation
└── README.md         # Documentation
```

## Commit Message Guidelines

Use clear, descriptive commit messages:

- **Add**: New features
- **Fix**: Bug fixes
- **Update**: Updates to existing features
- **Refactor**: Code refactoring
- **Docs**: Documentation changes
- **Test**: Test additions or changes

Example:
```
Add: Support for filtering by instance tags
Fix: Handle nil pointer in instance parsing
Update: Improve error messages for AWS authentication
```

## Review Process

- All PRs require review before merging
- Address review comments promptly
- Keep PRs focused and reasonably sized
- Update documentation as needed

## Questions?

Feel free to open an issue for questions or discussions about contributions.

Thank you for contributing to aws-eye! 🎉

