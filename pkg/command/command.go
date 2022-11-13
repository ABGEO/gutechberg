package command

import (
	"fmt"
	"io"

	"github.com/abgeo/gutechberg/pkg"
	"github.com/abgeo/gutechberg/pkg/config"
	"github.com/abgeo/gutechberg/pkg/platform"
	"github.com/spf13/cobra"
)

type Options struct {
	Writer            io.Writer
	Version           pkg.Version
	Platforms         []platform.Interface
	Sections          []config.Section
	SectionsOverrides map[string][]config.Section
}

type Interface interface {
	Complete(cmd *cobra.Command, args []string) error
	Validate() error
	Run() error
}

func GetRunner(command Interface) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(command.Complete(cmd, args))
		cobra.CheckErr(command.Validate())
		cobra.CheckErr(command.Run())
	}
}

func NewRootCommand(options Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gutechberg",
		Version: fmt.Sprintf("%s (%s)\n%s", options.Version.Number, options.Version.Date, options.Version.Commit),
		Short:   "CLI tool to improve lives of tech writers",
		Long: `  ____         _               _      _                       
 / ___| _   _ | |_  ___   ___ | |__  | |__    ___  _ __  __ _ 
| |  _ | | | || __|/ _ \ / __|| '_ \ | '_ \  / _ \| '__|/ _' |
| |_| || |_| || |_|  __/| (__ | | | || |_) ||  __/| |  | (_| |
 \____| \__,_| \__|\___| \___||_| |_||_.__/  \___||_|   \__, |
                                                        |___/ 
CLI tool to improve lives of tech writers.
		`,
	}

	return cmd
}
