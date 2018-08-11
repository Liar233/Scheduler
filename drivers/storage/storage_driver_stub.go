package main

import (
	"fmt"
)

const PluginName = "storage_driver_stub"
const PluginVersion = 1

var configFields = []string{
	"port",
	"host",
	"username",
	"password",
	"database",
}

// Get plugin name and version
func GetDescription() (string, int) {
	return PluginName, PluginVersion
}

// Validate config from Scheduler
func ValidateConfig(config map[string]interface{}) []error {
	var validationErrors []error
	var err error

	for _, fieldName := range configFields {
		if _, ok := config[fieldName]; !ok {
			err = fmt.Errorf("module `%s`, config field `%s` not found", PluginName, fieldName)

			validationErrors = append(validationErrors, err)
		}

		if config[fieldName] == "" {
			err = fmt.Errorf("module `%s`, config field `%s` empty", PluginName, fieldName)

			validationErrors = append(validationErrors, err)
		}
	}

	return validationErrors
}
