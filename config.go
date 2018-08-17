package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"plugin"
	"path/filepath"
)

type DriverConfig struct {
	DriverPath string                 `json:"driverPath" yaml:"driverPath"`
	Options    map[string]interface{} `json:"options" yaml:"options"`
}

// Validating plugin config
func (d *DriverConfig) Validate(name string) []error {
	var validationErrors []error

	if d.DriverPath == "" {
		err := fmt.Errorf("module `%s` has no config parameter `driverPath`", name)

		return append(validationErrors, err)
	}

	if driverPath, err := filepath.Abs(d.DriverPath); err != nil {
		err := fmt.Errorf("module `%s` invalid config parameter `driverPath`")

		return append(validationErrors, err)
	} else {
		d.DriverPath = driverPath
	}

	if _, err := os.Stat(d.DriverPath); err != nil {
		err = fmt.Errorf("module `%s` driver file `%s` not found", name, d.DriverPath)

		return append(validationErrors, err)
	}

	if p, err := plugin.Open(d.DriverPath); err != nil {
		err = fmt.Errorf("module `%s` fail to load plugin `%s` error: %s", name, d.DriverPath, err)

		return append(validationErrors, err)
	} else {
		if pluginConfValidator, syncError := p.Lookup("ValidateConfig"); syncError != nil {
			syncError = fmt.Errorf("module `%s` fail to load plugin `%s` error: %s", name, d.DriverPath, syncError)

			return append(validationErrors, syncError)
		} else {
			pluginConfValidatorFunc, ok := pluginConfValidator.(func(map[string]interface{}) []error)

			if !ok {
				err := fmt.Errorf("module `%s` fail to load plugin `%s` error: %s", name, d.DriverPath, syncError)

				return append(validationErrors, err)
			}

			pluginConfigErrors := pluginConfValidatorFunc(d.Options)

			if len(pluginConfigErrors) > 0 {
				return pluginConfigErrors
			}
		}
	}

	return nil
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

	validation = append(validation, config.Storage.Validate("storage")...)

	for name, channelConfig := range config.Channels {
		validation = append(validation, channelConfig.Validate(name)...)
	}

	return validation
}
