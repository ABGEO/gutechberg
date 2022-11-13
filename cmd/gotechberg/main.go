package main

import (
	"os"

	"github.com/abgeo/gutechberg/pkg"
	"github.com/abgeo/gutechberg/pkg/command"
	"github.com/abgeo/gutechberg/pkg/command/publish"
	"github.com/abgeo/gutechberg/pkg/config"
	"github.com/abgeo/gutechberg/pkg/platform"
	"github.com/abgeo/gutechberg/pkg/platform/twitter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "unknown" //nolint:gochecknoglobals
	date    = "unknown" //nolint:gochecknoglobals
)

func main() {
	var err error

	appVersion := pkg.Version{
		Number: version,
		Commit: commit,
		Date:   date,
	}
	writer := os.Stdout
	options := command.Options{
		Writer:  writer,
		Version: appVersion,
	}

	options.Platforms, err = registerPlatforms(
		twitter.New,
	)
	cobra.CheckErr(err)

	options.Sections, options.SectionsOverrides, err = config.GetContent()
	cobra.CheckErr(err)

	cmd := registerCommands(options)
	cobra.CheckErr(cmd.Execute())
}

func registerPlatforms(constructors ...func() (platform.Interface, error)) (platforms []platform.Interface, err error) {
	for _, constructor := range constructors {
		platformInstance, err := constructor()
		if err != nil {
			return platforms, errors.Wrap(err, "")
		}

		platforms = append(platforms, platformInstance)
	}

	return platforms, nil
}

func registerCommands(options command.Options) *cobra.Command {
	cmdRoot := command.NewRootCommand(options)

	cmdPublish := publish.NewCommand(options)

	cmdRoot.AddCommand(cmdPublish)

	return cmdRoot
}
