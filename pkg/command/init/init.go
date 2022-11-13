package init

import (
	"io"
	"os"

	"github.com/abgeo/gutechberg/pkg/command"
	"github.com/abgeo/gutechberg/pkg/config"
	"github.com/abgeo/gutechberg/pkg/platform"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const ExampleConfig = `# This is an example .gutechberg.yaml file with some sensible defaults.
platforms:
  twitter:
    enabled: true
    consumerKey: "some-consumer-key"
    consumerSecret: "some-consumer-secret"
    accessToken: "some-access-token"
    accessSecret: "some-access-secret"

content:
  sections:
    - id: foo
      template: |+
        foo,
        
        bar, baz!

    - id: bar
      include: ./bar.md
    - id: baz
      template: baz, bar, foo...

  overrides:
    twitter:
      - id: bar
        template: |+
          bar,
          
          baz,
          
          foo

`

type Command struct {
	writer            io.Writer
	platforms         []platform.Interface
	sections          []config.Section
	sectionsOverrides map[string][]config.Section
}

func NewCommand(options command.Options) *cobra.Command {
	cmd := &Command{
		writer:            options.Writer,
		platforms:         options.Platforms,
		sections:          options.Sections,
		sectionsOverrides: options.SectionsOverrides,
	}

	cobraCmd := &cobra.Command{
		Use:   "init",
		Short: "Generates a .gutechberg.yaml file",
		Args:  cobra.NoArgs,
		Run:   command.GetRunner(cmd),
	}

	return cobraCmd
}

func (command *Command) Complete(cmd *cobra.Command, args []string) error { return nil }

func (command *Command) Validate() error { return nil }

func (command *Command) Run() error {
	file, err := os.Create(".gutechberg.yaml")
	if err != nil {
		return errors.Wrap(err, "unable to create .gutechberg.yaml")
	}

	defer file.Close()

	_, err = file.WriteString(ExampleConfig)
	if err != nil {
		return errors.Wrap(err, "unable to write into .gutechberg.yaml")
	}

	pterm.DefaultParagraph.
		WithWriter(command.writer).
		Println("Example .gutechberg.yaml file has been created successfully.")

	return nil
}
