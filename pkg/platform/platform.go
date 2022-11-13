package platform

type Interface interface {
	GetKey() string
	GetConfig() ConfigInterface
	Publish(content string)
}

type ConfigInterface interface {
	IsEnabled() bool
}

type Config struct {
	Enabled bool
}

func (conf *Config) IsEnabled() bool {
	return conf.Enabled
}
