package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Section struct {
	ID       string
	Template string
}

type sectionItem struct {
	ID       string `validate:"required"`
	Template string `validate:"required_without=Include"`
	Include  string `validate:"required_without=Template"`
}

type content struct {
	Sections  []sectionItem            `validate:"required,dive,required"`
	Overrides map[string][]sectionItem `validate:"dive,dive,required"`
}

func ReadPlatformConfig(platform string, conf any) error {
	return readKey("platforms."+platform, conf)
}

func readKey(key string, conf any) error {
	viperInstance := viper.New()

	viperInstance.SetConfigName(".gutechberg")
	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath(".")

	err := viperInstance.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "unable to read config")
	}

	err = viperInstance.UnmarshalKey(key, conf)
	if err != nil {
		return errors.Wrapf(err, "unable to unmarshal key: %s", key)
	}

	validate := validator.New()
	if err = validate.Struct(conf); err != nil {
		return errors.Wrapf(err, "config under key '%s' is not valid", key)
	}

	return nil
}

func GetContent() (sections []Section, overrides map[string][]Section, err error) {
	overrides = make(map[string][]Section)
	configStruct := new(content)

	err = readKey("content", configStruct)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to read content")
	}

	sections, err = composeSectionFromRaw(configStruct.Sections)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to read content sections")
	}

	for platformID, platformOverrides := range configStruct.Overrides {
		overrides[platformID], err = composeSectionFromRaw(platformOverrides)
		if err != nil {
			return nil, nil, errors.Wrap(err, "unable to read content overrides")
		}
	}

	return sections, overrides, nil
}

func composeSectionFromRaw(items []sectionItem) (sections []Section, err error) {
	for _, item := range items {
		template := item.Template

		if item.Include != "" {
			data, err := os.ReadFile(item.Include)
			if err != nil {
				return nil, errors.Wrap(err, "unable to read include file")
			}

			template = string(data)
		}

		sections = append(sections, Section{
			ID:       item.ID,
			Template: template,
		})
	}

	return sections, nil
}
