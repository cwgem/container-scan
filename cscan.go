package main

import (
	"os"

	"github.com/cwgem/container-scan/cli"
)

func main() {
	app := cli.SetupCli()
	app.Run(os.Args)
}
