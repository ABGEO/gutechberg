package twitter

import (
	"fmt"

	"github.com/abgeo/gutechberg/pkg/config"
	"github.com/abgeo/gutechberg/pkg/platform"
	"github.com/pkg/errors"
)

type iPlatform struct {
	Config *Config
}

type Config struct {
	platform.Config `mapstructure:",squash"`

	ConsumerKey    string `validate:"required_if=Enabled true"`
	ConsumerSecret string `validate:"required_if=Enabled true"`
	AccessToken    string `validate:"required_if=Enabled true"`
	AccessSecret   string `validate:"required_if=Enabled true"`
}

func New() (platform.Interface, error) {
	platformInstance := new(iPlatform)
	platformInstance.Config = new(Config)

	err := config.ReadPlatformConfig(platformInstance.GetKey(), platformInstance.Config)
	if err != nil {
		return platformInstance, errors.Wrapf(err, "unable to read config for platform %s", platformInstance.GetKey())
	}

	return platformInstance, nil
}

func (platform *iPlatform) GetKey() string {
	return "twitter"
}

func (platform *iPlatform) GetConfig() platform.ConfigInterface {
	return platform.Config
}

func (platform *iPlatform) Publish(content string) {
	fmt.Println(content)
}
