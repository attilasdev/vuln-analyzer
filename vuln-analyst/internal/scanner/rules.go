package scanner

import (
	"path/filepath"
	"regexp"
	"strings"
)

type Rule struct {
	ID          string   `yaml:"id"`
	Name        string   `yaml:"name"`
	Pattern     string   `yaml:"pattern"`
	Description string   `yaml:"description"`
	Severity    string   `yaml:"severity"`
	Category    string   `yaml:"category"`
	Languages   []string `yaml:"languages"`
	regex       *regexp.Regexp
}

func (r *Rule) Compile() error {
	if r.regex != nil {
		return nil
	}

	regex, err := regexp.Compile(r.Pattern)
	if err != nil {
		return err
	}
	r.regex = regex
	return nil
}

func (r *Rule) Match(line string) bool {
	if r.regex == nil {
		r.Compile()
	}
	return r.regex != nil && r.regex.MatchString(line)
}

func (r *Rule) IsApplicableToFile(filePath string) bool {
	ext := strings.TrimPrefix(filepath.Ext(filePath), ".")

	// If rule applies to all languages
	for _, lang := range r.Languages {
		if lang == "all" {
			return true
		}
	}

	// Map file extensions to languages
	langMap := map[string][]string{
		"javascript": {"js", "jsx"},
		"typescript": {"ts", "tsx"},
		"python":     {"py"},
		"php":        {"php"},
		"java":       {"java"},
		"ruby":       {"rb"},
		"go":         {"go"},
	}

	// Check if file extension matches any of the specified languages
	for _, lang := range r.Languages {
		if extensions, ok := langMap[lang]; ok {
			for _, validExt := range extensions {
				if ext == validExt {
					return true
				}
			}
		}
	}

	return false
}
