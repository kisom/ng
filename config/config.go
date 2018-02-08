package config

import (
	"os"

	"gopkg.in/yaml.v1"
)

func getDefault(s *string, envName string, defaultValue string) {
	if *s != "" {
		return
	}

	env := os.Getenv(envName)
	if env != "" {
		*s = env
		return
	}

	*s = defaultValue
}

// Compilers stores any configured compilers (and linkers).
type Compilers struct {
	CC  string `json:"cc" yaml:"cc"`
	CXX string `json:"cxx" yaml: "cxx"`
	LD  string `json:"ld" yaml: "ld"`
}

func (c *Compilers) Setup() {
	getDefault(&c.CC, "CC", "clang")
	getDefault(&c.CXX, "CXX", "clang++")
	getDefault(&c.LD, "LD", "clang++")
}

type Targets struct {
	CC  map[string][]string `json:"c" yaml:"c"`
	CXX map[string][]string `json:"cc" yaml:"cc"`
}

type Config struct {
	Debug     bool       `json:"debug" yaml:"debug"`
	NoDeps    bool       `json:"no_deps" yaml:"no_deps"`
	Compilers *Compilers `json:"compilers" yaml:"compilers"`
	Targets   *Targets   `json:"targets" yaml:"targets"`
}

func (c *Config) RequiresCXX() bool {
	return len(c.Targets.CXX) != 0
}

func (c *Config) RequiresCC() bool {
	return len(c.Targets.CC) != 0
}

func (c *Config) Setup() {
	c.Compilers.Setup()
}

func Parse(in []byte) (*Config, error) {
	cfg := &Config{Debug: true}
	err := yaml.Unmarshal(in, cfg)
	if err != nil {
		return nil, err
	}

	cfg.Setup()
	return cfg, nil
}
