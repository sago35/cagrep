package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "cagrep"
	app.Usage = "grep with cashed server"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "server, s",
			Usage: "server mode",
		},
		cli.IntFlag{
			Name:  "port, p",
			Value: 9999,
			Usage: "server port",
		},
	}
	app.Action = func(c *cli.Context) {
		if c.Bool("server") {
			cagrepServer(c)
		} else {
			queryGrep(c)
		}
	}

	app.Run(os.Args)
}
