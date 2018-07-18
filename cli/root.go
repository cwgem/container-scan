package cli

import (
	mainCli "gopkg.in/urfave/cli.v1"
)

// GlobalOptions holds CLI options that should be passed down to subcommands
type GlobalOptions struct {
	Basedir string
}

// SetupCli initializes the core cli
func SetupCli() *mainCli.App {
	// Base Information
	app := mainCli.NewApp()
	app.Name = "container-scan"
	app.Usage = "Scan Docker images and containers for malware and other security issues"
	app.Version = "0.1.0"

	// Flag Setup
	baseOptions := &GlobalOptions{}
	app.Flags = []mainCli.Flag{
		mainCli.StringFlag{
			Name:        "basedir, b",
			Usage:       "Use `BASEDIR` for unpacking and other operations (default: random temp directory)",
			Destination: &baseOptions.Basedir,
		},
	}

	// Subcommand Setup
	app.Commands = []mainCli.Command{
		SetupImageCommand(baseOptions),
	}

	return app
}
