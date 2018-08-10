package main

type StorageConfig struct {
	Name string
	DriverPath string
	Options map[string]interface{}
}

type ChannelConfig struct {
	DriverPath string
	Options map[string]interface{}
}

type AppConfig struct {
	Master bool
	Port uint
	Storage StorageConfig
	Channels map[string]ChannelConfig
}
