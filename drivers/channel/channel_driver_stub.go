package main

import (
	"fmt"
)

const PluginName = "channel_driver_stub"
const PluginVersion = 1

var configFields = []string{
	"port",
	"host",
	"username",
	"password",
}

// Get plugin name and version
func GetDescription() (string, int) {
	return PluginName, PluginVersion
}

// Validate config from Scheduler
func ValidateConfig(config map[string]interface{}) error {
	for _, fieldName := range configFields {
		if _, ok := config[fieldName]; !ok {
			return fmt.Errorf("module `%s`, config field `%s` not found", PluginName, fieldName)
		}

		if config[fieldName] == "" {
			return fmt.Errorf("module `%s`, config field `%s` empty", PluginName, fieldName)
		}
	}

	return nil
}
