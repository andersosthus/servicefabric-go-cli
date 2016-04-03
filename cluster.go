package main

import (
	"github.com/andersosthus/servicefabric-go-cli/cmd"
	"github.com/codegangsta/cli"
	"github.com/andersosthus/servicefabric-go-cli/helpers"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "servicefabric-cli"

	app.Commands = []cli.Command{
		cmd.CmdCluster,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: helpers.AUTHPFXFILE,
			Value: "",
			Usage: "Path to pfx file",
		},
		cli.StringFlag{
			Name: helpers.AUTHPFXCODE,
			Value: "",
			Usage: "Passcode to PFX",
		},
		cli.StringFlag{
			Name: helpers.CLUSTERENDPOINT,
			Value: "http://localhost:19000",
			Usage: "Cluster endpoint",
		},
	}

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	app.Run(os.Args)
}
