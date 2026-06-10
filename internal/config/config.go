package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/prometheus/alertmanager/config"
)

type TestCase struct {
	Name              string            `yaml:"name"`
	Labels            map[string]string `yaml:"labels"`
	ExpectedReceivers []string          `yaml:"expected_receivers"`
}

type TestSuite struct {
	Tests []TestCase `yaml:"tests"`
}

func LoadAlertmanagerConfig(path string) (*config.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var cfg config.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}

	return &cfg, nil
}

func LoadTestCases(path string) ([]TestCase, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var suite TestSuite
	if err := yaml.Unmarshal(data, &suite); err != nil {
		return nil, fmt.Errorf("error parsing test cases: %w", err)
	}

	return suite.Tests, nil
}
