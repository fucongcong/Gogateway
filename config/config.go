package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	RConf *RegistryConfig
)

type RegistryConfig struct {
	RedisNetWork string `yaml:"registry.redis.network"`
	RedisAddr    string `yaml:"registry.redis.addr"`
}

func NewRegistryConfig(fp string) *RegistryConfig {
	confFile, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	c := &RegistryConfig{
		RedisNetWork: "tcp",
		RedisAddr:    "127.0.0.1:6379",
	}

	err = yaml.Unmarshal(confFile, c)
	if err != nil {
		panic(err)
	}

	return c
}

func GetRegistryConfig() *RegistryConfig {
	return RConf
}

func ParseConfig(path string) {
	RConf = NewRegistryConfig(path)
}
