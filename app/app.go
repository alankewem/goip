package app

import (
	"goip/cmd"

	"github.com/urfave/cli"
)

func New() *cli.App {
	app := cli.NewApp()

	app.Author = "Alan Kewem"
	app.Name = "GoIP"

	app.Commands = []cli.Command{{
		Name:  "ip",
		Usage: "Search public IP address of the host",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "host", Required: true},
		},
		Action: cmd.SearchPublicIpAddress,
	}}

	return app
}
