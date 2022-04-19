package main

import (
	"fornever.org/app"
	"github.com/urfave/cli"
)

var commandEntry = cli.Command{
	Name:   "entry",
	Usage:  "program entry",
	Action: entry,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "addr",
			EnvVar: "LISTEN_ADDR",
			Value:  "0.0.0.0:8080",
		},
	},
}

func entry(c *cli.Context) error {

	inst := app.CreateApp(&app.WebAppParam{
		ServiceName: AppName,
		Version:     Version,
		Flag1:       false,
	})
	return inst.Run(c.String("addr"))

}
