package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ClientConfig struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Port int    `yaml:"port"`
}

func GetClientConfig() ClientConfig {
	f, err := os.ReadFile("./clientConfig.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	var clientConfig ClientConfig
	err = yaml.Unmarshal(f, &clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	if clientConfig.Host == "" {
		clientConfig.Host = "127.0.0.1"
	}
	if clientConfig.Port == 0 {
		clientConfig.Port = 623
	}
	return clientConfig
}
