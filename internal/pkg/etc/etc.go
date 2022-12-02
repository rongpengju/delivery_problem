package etc

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Calculate _Calculate `yaml:"calculate" validate:"required"`
	Database  _Database  `yaml:"database" validate:"required"`
}

type _Calculate struct {
	BasicPrice    int8 `yaml:"basic_price"`
	OverPrice     int8 `yaml:"over_price"`
	BasicWeight   int8 `yaml:"basic_weight"`
	InsuranceRate int8 `yaml:"insurance_rate"`
}

type _Database struct {
	// Dot use
}

func LoadFromFile(f string) (*Config, error) {
	buf, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("read file %q: %w", f, err)
	}

	ret := &Config{}
	if err := yaml.Unmarshal(buf, ret); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return ret, nil
}
