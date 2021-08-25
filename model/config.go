package model

import "gopkg.in/yaml.v2"

type Config struct {
	ServerAddr string `yaml:"server_addr"`
	DB Database `yaml:"database"`
}
type Database struct {
	User     string `yaml:"user"`
	PassWd   string `yaml:"pass_wd"`
	Addr     string `yaml:"addr"`
	Name     string `yaml:"name"`
}

func UnmarshalConfig(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}