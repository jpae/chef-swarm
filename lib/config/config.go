package config

import (
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
)

type Config struct {
	Server        string
	Organizations map[string]OrgConfig
}

type OrgConfig struct {
	ValidatorName string
	ValidatorKey  string
	KeyDirectory  string
	Nodes         int
	Runs          int
	Stagger       bool
}

func LoadConfig(file string) (cfg Config, err error) {
	f, err := os.Open(file)
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return cfg, err
	}

	err = toml.Unmarshal(buf, &cfg)

	return cfg, err
}
