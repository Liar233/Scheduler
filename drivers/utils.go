package drivers

import (
	"github.com/Liar233/Scheduler/config"
	"plugin"
)

func LoadPluginDriver(conf *config.DriverConfig) (interface{}, error) {
	p, err := plugin.Open(conf.DriverPath)

	if err != nil {
		return nil, err
	}

	return p, nil
}
