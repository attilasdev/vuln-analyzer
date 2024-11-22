package scanner

import (
	"os"

	"gopkg.in/yaml.v3"
)

type RuleLoader struct {
	rulesPath string
}

type RuleSet struct {
	Rules []*Rule `yaml:"rules"`
}

func NewRuleLoader(rulesPath string) *RuleLoader {
	return &RuleLoader{
		rulesPath: rulesPath,
	}
}

func (rl *RuleLoader) LoadRules() ([]*Rule, error) {
	// Try to load rules from YAML file
	data, err := os.ReadFile(rl.rulesPath)
	if err != nil {
		// If file not found or other error, return default rules
		return GetDefaultRules(), nil
	}

	// Parse YAML
	var ruleSet RuleSet
	if err := yaml.Unmarshal(data, &ruleSet); err != nil {
		return nil, err
	}

	// Compile regex patterns
	for _, rule := range ruleSet.Rules {
		if rule != nil {
			if err := rule.Compile(); err != nil {
				return nil, err
			}
		}
	}

	return ruleSet.Rules, nil
}

func GetDefaultRules() []*Rule {
	rules := []*Rule{
		{
			ID:          "HARDCODED_CREDENTIALS",
			Name:        "Hardcoded Credentials",
			Pattern:     `(?i)(password|passwd|pwd|secret|key|token|api[_-]?key)\s*[:=]\s*['"][^'"\s]{3,}['"]`,
			Description: "Hardcoded credentials detected",
			Severity:    "HIGH",
			Category:    "AUTH",
			Languages:   []string{"all"},
		},
		{
			ID:          "SQL_INJECTION",
			Name:        "SQL Injection",
			Pattern:     `(?i)(execute|query|select|insert|update|delete)\s*[(\s].*?([\+\.]|concat).*?(\$|@|#)[a-zA-Z0-9_]+`,
			Description: "Possible SQL injection vulnerability",
			Severity:    "HIGH",
			Category:    "INJECTION",
			Languages:   []string{"php", "python", "java", "ruby"},
		},
		{
			ID:          "XSS_VULNERABILITY",
			Name:        "Cross-site Scripting (XSS)",
			Pattern:     `(?i)(innerHTML|outerHTML|document\.write)\s*=\s*`,
			Description: "Possible XSS vulnerability",
			Severity:    "HIGH",
			Category:    "INJECTION",
			Languages:   []string{"javascript", "typescript"},
		},
	}

	// Compile the default rules
	for _, rule := range rules {
		rule.Compile()
	}

	return rules
}
