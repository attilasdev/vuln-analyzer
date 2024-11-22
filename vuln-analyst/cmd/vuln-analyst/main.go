package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"vuln-analyst/internal/scanner"
)

func main() {
	targetPath := flag.String("path", ".", "Path to scan")
	rulesPath := flag.String("rules", "configs/rules.yaml", "Path to rules configuration")
	flag.Parse()

	// Load rules
	ruleLoader := scanner.NewRuleLoader(*rulesPath)
	rules, err := ruleLoader.LoadRules()
	if err != nil {
		log.Fatalf("Error loading rules: %v", err)
	}

	if len(rules) == 0 {
		log.Println("Warning: No rules loaded. Using default rules.")
	}

	// Create scanner with loaded rules
	scannerInstance := scanner.NewScanner(rules)
	if scannerInstance == nil {
		log.Fatal("Failed to create scanner instance")
	}

	absPath, err := filepath.Abs(*targetPath)
	if err != nil {
		log.Fatalf("Error resolving path: %v", err)
	}

	fmt.Printf("Starting scan of: %s\n", absPath)
	vulnerabilities, err := scannerInstance.ScanDirectory(absPath)
	if err != nil {
		log.Fatalf("Error scanning directory: %v", err)
	}

	// Print results
	if len(vulnerabilities) == 0 {
		fmt.Println("No vulnerabilities found!")
		return
	}

	fmt.Printf("\nFound %d potential vulnerabilities:\n\n", len(vulnerabilities))
	for _, vuln := range vulnerabilities {
		fmt.Printf("------------------\n")
		fmt.Printf("Vulnerability: %s\n", vuln.Name)
		fmt.Printf("Severity: %s\n", vuln.Severity)
		fmt.Printf("Location: %s:%d\n", vuln.Location, vuln.LineNumber)
		fmt.Printf("Description: %s\n", vuln.Description)
		fmt.Printf("Code: %s\n", vuln.Code)
		fmt.Printf("------------------\n\n")
	}
}
