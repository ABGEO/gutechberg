package publish

import (
	"github.com/abgeo/gutechberg/pkg/command"
	"github.com/abgeo/gutechberg/pkg/config"
	"github.com/abgeo/gutechberg/pkg/platform"
	"github.com/spf13/cobra"
)

type Command struct {
	platforms         []platform.Interface
	sections          []config.Section
	sectionsOverrides map[string][]config.Section
}

func NewCommand(options command.Options) *cobra.Command {
	cmd := &Command{
		platforms:         options.Platforms,
		sections:          options.Sections,
		sectionsOverrides: options.SectionsOverrides,
	}

	cobraCmd := &cobra.Command{
		Use:   "publish",
		Short: "Publish article",
		Args:  cobra.NoArgs,
		Run:   command.GetRunner(cmd),
	}

	return cobraCmd
}

func (command *Command) Complete(cmd *cobra.Command, args []string) error { return nil }

func (command *Command) Validate() error { return nil }

func (command *Command) Run() error {
	for _, item := range command.platforms {
		if item.GetConfig().IsEnabled() {
			var sectionOverrides []config.Section
			if overrides, ok := command.sectionsOverrides[item.GetKey()]; ok {
				sectionOverrides = overrides
			}

			item.Publish(getContent(command.sections, sectionOverrides))
		}
	}

	return nil
}

// @todo: move to helper.
func getContent(sections []config.Section, overrides []config.Section) string {
	content := ""

	overridesMap := make(map[string]string)
	for _, override := range overrides {
		overridesMap[override.ID] = override.Template
	}

	for _, section := range sections {
		if override, ok := overridesMap[section.ID]; ok {
			content += override
		} else {
			content += section.Template
		}
	}

	return content
}
