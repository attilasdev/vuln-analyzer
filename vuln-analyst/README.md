# Vuln-Analyst 🛡️

Vuln-Analyst is a lightweight static code analysis tool built in Go that scans source code for potential security vulnerabilities. It helps developers identify common security issues early in the development process.

## Features 🚀

- **Static Code Analysis**: Scans code without execution
- **Multi-Language Support**: JavaScript, Python, PHP, Java, Ruby
- **Configurable Rules**: Custom rule definitions via YAML
- **Severity Levels**: HIGH, MEDIUM, LOW vulnerability classification
- **Easy Integration**: Simple command-line interface

## Installation 📥

### Prerequisites
- Go 1.16 or higher

### Steps
```bash
# Clone the repository
git clone https://github.com/attilasdev/vuln-analyst
cd vuln-analyst

# Build the project
go build -o vuln-analyst cmd/vuln-analyst/main.go
```

## Usage 💻

### Basic Scanning
```bash
# Scan current directory
./vuln-analyst

# Scan specific directory
./vuln-analyst -path /path/to/code

# Scan with specific severity level
./vuln-analyst -path /path/to/code -severity high
```

### Custom Rules
```bash
# Scan with custom rules file
./vuln-analyst -path /path/to/code -rules /path/to/rules.yaml
```

## Vulnerability Detection 🔍

Vuln-Analyst can detect various security vulnerabilities including:

- SQL Injection
- Cross-Site Scripting (XSS)
- Hardcoded Credentials
- Weak Cryptography
- Command Injection
- Path Traversal

## Example Output 📋

```
Starting scan of: /path/to/code
Found 2 potential vulnerabilities:

------------------
Vulnerability: Hardcoded Credentials
Severity: HIGH
Location: config.js:15
Description: Hardcoded credentials detected
Code: const apiKey = "1234567890abcdef"
------------------

------------------
Vulnerability: SQL Injection
Severity: HIGH
Location: query.php:25
Description: Possible SQL injection vulnerability
Code: $result = query("SELECT * FROM users WHERE id = " . $_GET['id']);
------------------
```

## Configuration ⚙️

### Rules Configuration
Create a `rules.yaml` file to define custom scanning rules:

```yaml
rules:
  - id: "HARDCODED_CREDENTIALS"
    name: "Hardcoded Credentials"
    pattern: '(?i)(password|passwd|pwd|secret|key|token|api[_-]?key)\s*[:=]\s*[''""][^''"'\s]{3,}[''""]'
    description: "Hardcoded credentials detected"
    severity: "HIGH"
    category: "AUTH"
    languages: ["all"]
```

## Project Structure 📁

```
codeguardian/
├── cmd/
│   └── codeguardian/
│       └── main.go
├── internal/
│   └── scanner/
│       ├── scanner.go
│       ├── rules.go
│       └── rule_loader.go
├── pkg/
│   └── models/
│       └── vulnerability.go
├── configs/
│   └── rules.yaml
└── README.md
```

## Contributing 🤝

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Security Considerations ⚠️

CodeGuardian is a static analysis tool designed for educational and development purposes. It:
- Does not execute scanned code
- Performs pattern-based analysis only
- Should not be the only security testing tool in your pipeline
- May produce false positives/negatives

Note: This is not a comprehensive security tool and should not be relied upon as the sole method of security testing.

## Limitations 📌

- Pattern-based detection may miss complex vulnerabilities
- Limited to static code analysis
- Language support is based on pattern matching
- Does not perform data flow analysis

## Disclaimer ⚖️

This tool is provided as-is for educational and development purposes. It should not be relied upon as the sole method of security testing. Always perform comprehensive security testing using multiple tools and methodologies.

## License
This is a portfolio project created for educational purposes.
Feel free to use, modify and learn from it.

---

Made with ❤️ by Attila
