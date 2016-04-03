package cmd

import "github.com/codegangsta/cli"

var CmdCluster = cli.Command{
	Name:    "cluster",
	Aliases: []string{"c"},
	Usage:   "Cluster level operations",
	Subcommands: []cli.Command{
		CmdSystemServices,
		CmdClusterHealth,
	},
}
