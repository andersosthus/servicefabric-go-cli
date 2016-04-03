package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/andersosthus/servicefabric-go-cli/helpers"
	"github.com/andersosthus/servicefabric-go-sdk/cluster"
	"os"
	"encoding/json"
)

var CmdClusterHealth = cli.Command{
	Name:    "clusterhealth",
	Aliases: []string{"health"},
	Usage:   "Gets the health of a Service Fabric cluster.",
	Action:  ClusterHealth,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "filter",
			Value: "",
			Usage: "show only service named",
		},
	},
}

func ClusterHealth(c *cli.Context) {
	auth, options := helpers.SettingsFromArgs(c)
	options.EndpointUrl = "/$/GetClusterHealthChunk"

	resp, err := cluster.GetClusterHealth(auth, options)

	if err != nil {
		panic(err)
	}

	pretty, err := json.MarshalIndent(*resp, "", "  ")

	os.Stdout.Write(pretty)
}
