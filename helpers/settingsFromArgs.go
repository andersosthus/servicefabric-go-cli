package helpers

import (
	"github.com/codegangsta/cli"
	"github.com/andersosthus/servicefabric-go-sdk/cluster"
)

func SettingsFromArgs(c *cli.Context) (cluster.Auther, cluster.CallOptions) {
	options := cluster.CallOptions{}

	if c.GlobalIsSet(CLUSTERENDPOINT) {
		options.BaseUrl = c.GlobalString(CLUSTERENDPOINT)
	}

	if c.GlobalIsSet(AUTHPFXFILE) {
		auth := &cluster.PfxFileAuth{
			PfxFile: c.GlobalString(AUTHPFXFILE),
			PfxPasscode: c.GlobalString(AUTHPFXCODE),
		}

		return auth, options
	}

	return &cluster.PfxFileAuth{}, options
}
