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
	Port uint
	Storage StorageConfig
	Channels map[string]ChannelConfig
}
