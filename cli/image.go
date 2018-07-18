package cli

import (
	mainCli "gopkg.in/urfave/cli.v1"
)

// ImageOptions holds the options for the image command
type ImageOptions struct {
	Global *GlobalOptions
	Name   string
}

// SetupImageCommand initializes the image scan command and its respective flags
func SetupImageCommand(globalOptions *GlobalOptions) mainCli.Command {
	// Basic command info
	command := mainCli.Command{
		Name:  "image",
		Usage: "Runs scans against images",
	}

	// Arguments
	options := &ImageOptions{
		Global: globalOptions,
	}
	command.Flags = []mainCli.Flag{
		mainCli.StringFlag{
			Name:        "name, n",
			Usage:       "Scan image with with the name `NAME`",
			Destination: &options.Name,
		},
	}

	// Command processing
	command.Action = func(c *mainCli.Context) error {
		return nil
	}

	return command
}
