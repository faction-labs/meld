package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/factionlabs/meld/cmd/serve"
	"github.com/factionlabs/meld/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "meld"
	app.Usage = "game server management"
	app.Version = version.FullVersion()
	app.Author = "Faction Labs"
	app.Email = "https://factionlabs.io"
	app.Before = func(c *cli.Context) error {
		// enable debug
		if c.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
			log.Debug("debug enabled")
		}

		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug",
		},
		cli.StringFlag{
			Name:   "db-addr, d",
			Usage:  "rethinkdb address",
			Value:  "127.0.0.1:28015",
			EnvVar: "RETHINKDB_ADDR",
		},
		cli.StringFlag{
			Name:   "db-name, n",
			Usage:  "rethinkdb db name",
			Value:  "meld",
			EnvVar: "RETHINKDB_NAME",
		},
		cli.StringFlag{
			Name:  "store-key, k",
			Usage: "store key for auth sessions",
			Value: "meld-store-key",
		},
	}
	app.Commands = []cli.Command{
		serve.Command,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
