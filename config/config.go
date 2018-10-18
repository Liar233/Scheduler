package config

import (
	"os"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type DriverConfig struct {
	DriverPath string                 `json:"driverPath" yaml:"driverPath"`
	Options    map[string]interface{} `json:"options" yaml:"options"`
}

type StorageConfig struct {
	DriverConfig `yaml:",inline"`
}

type ChannelConfig struct {
	DriverConfig `yaml:",inline"`
}

type AppConfig struct {
	Master     bool                     `json:"master" yaml:"master"`
	Port       uint                     `json:"port" yaml:"port"`
	Storage    StorageConfig            `json:"storage" yaml:"storage"`
	Channels   map[string]ChannelConfig `json:"channels" yaml:"channels"`
	ApiTimeout uint                     `json:"api-timeout" yaml:"api-timeout"`
}

func NewAppConfig() *AppConfig{
	return &AppConfig{}
}

// Trying to fill AppConfig from yaml file
func LoadYamlConfig(filename string, config *AppConfig) []error {
	var validation []error
	var err error
	var data []byte

	if _, err = os.Stat(filename); err != nil {
		return []error{
			fmt.Errorf("config file `%s` does not exist", filename),
		}
	}

	data, err = ioutil.ReadFile(filename)

	if err != nil {
		return []error{
			fmt.Errorf("can't read `%s` file with error: %s", filename, err),
		}
	}

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return []error{
			fmt.Errorf("can't parse yaml with error: %s", err),
		}
	}

	return validation
}