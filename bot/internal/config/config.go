package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Conf struct {
	Apikey   string `yaml:"api_key"`
	AddrBot  string `yaml:"port"`
	AddrGrpc string `yaml:"port_grpc"`
}

func (c *Conf) LoadConfig(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Failed open config file :%v", err)
	}
	defer file.Close()
	read, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Failed reading file :%v", err)
	}
	err = yaml.Unmarshal(read, c)
	if err != nil {
		return fmt.Errorf("Yaml unmarshalling error :%v", err)
	}

	return nil
}
