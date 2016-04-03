package cmd

import (
	"encoding/json"
	"github.com/andersosthus/servicefabric-go-sdk/cluster"
	"github.com/codegangsta/cli"
	"os"
	"regexp"
	"github.com/andersosthus/servicefabric-go-cli/helpers"
)

var CmdSystemServices = cli.Command{
	Name:    "systemservices",
	Aliases: []string{"system"},
	Usage:   "Gets the Service Fabric system services.",
	Action:  SystemServices,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "filter",
			Value: "",
			Usage: "show only service named",
		},
	},
}

func SystemServices(c *cli.Context) {
	filter := ""
	if c.IsSet("filter") {
		filter = c.String("filter")
	}

	auth, options := helpers.SettingsFromArgs(c)
	options.EndpointUrl = "/Applications/System/$/GetServices"

	resp, err := cluster.GetSystemServices(auth, options)

	if err != nil {
		panic(err)
	}

	var result []cluster.SystemService

	if filter != "" {
		re, _ := regexp.Compile("(?i)" + filter)
		for _, item := range *resp {
			if re.MatchString(item.Name) {
				result = append(result, item)
			}
		}
	} else {
		result = *resp
	}

	pretty, err := json.MarshalIndent(result, "", "  ")
	os.Stdout.Write(pretty)
}
