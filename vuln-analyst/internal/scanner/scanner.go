package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"vuln-analyst/pkg/models"
)

type Scanner struct {
	rules []*Rule
}

func NewScanner(rules []*Rule) *Scanner {
	if rules == nil {
		rules = GetDefaultRules()
	}
	return &Scanner{
		rules: rules,
	}
}

func (s *Scanner) ScanDirectory(root string) ([]models.Vulnerability, error) {
	var allVulnerabilities []models.Vulnerability

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-code files
		if info.IsDir() || !isCodeFile(path) {
			return nil
		}

		vulns, err := s.ScanFile(path)
		if err != nil {
			return err
		}

		allVulnerabilities = append(allVulnerabilities, vulns...)
		return nil
	})

	return allVulnerabilities, err
}

func (s *Scanner) ScanFile(path string) ([]models.Vulnerability, error) {
	var vulnerabilities []models.Vulnerability

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		for _, rule := range s.rules {
			if rule == nil {
				continue
			}

			if !rule.IsApplicableToFile(path) {
				continue
			}

			if rule.Match(line) {
				vuln := models.Vulnerability{
					ID:          rule.ID,
					Name:        rule.Name,
					Description: rule.Description,
					Severity:    rule.Severity,
					Location:    path,
					LineNumber:  lineNum,
					Code:        line,
				}
				vulnerabilities = append(vulnerabilities, vuln)
			}
		}
	}

	return vulnerabilities, scanner.Err()
}

func isCodeFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	codeExtensions := map[string]bool{
		".go":   true,
		".js":   true,
		".py":   true,
		".java": true,
		".php":  true,
		".rb":   true,
		".cpp":  true,
		".c":    true,
		".ts":   true,
		".jsx":  true,
		".tsx":  true,
	}
	return codeExtensions[ext]
}
