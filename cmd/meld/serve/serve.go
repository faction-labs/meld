package serve

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/factionlabs/meld/api"
	"github.com/factionlabs/meld/version"
)

var Command = cli.Command{
	Name:   "serve",
	Usage:  "start server",
	Action: serveAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "listen, l",
			Usage: "listen address",
			Value: ":8080",
		},
		cli.StringFlag{
			Name:  "rcon-address",
			Usage: "rcon address",
			Value: "127.0.0.1:28016",
		},
		cli.StringFlag{
			Name:  "rcon-password",
			Usage: "rcon password",
			Value: "m3ld",
		},
		cli.StringFlag{
			Name:  "public-dir, s",
			Usage: "path to public media directory",
			Value: "public",
		},
		cli.StringFlag{
			Name:  "db-path, d",
			Usage: "path to meld db",
			Value: "meld.db",
		},
	},
}

func serveAction(c *cli.Context) {
	log.Infof("%s %s", version.FullName(), version.FullVersion())

	listenAddr := c.String("listen")
	publicDir := c.String("public-dir")
	dbPath := c.String("db-path")
	rconAddress := c.String("rcon-address")
	rconPassword := c.String("rcon-password")

	cfg := &api.APIConfig{
		ListenAddr:   listenAddr,
		PublicDir:    publicDir,
		DBPath:       dbPath,
		RconAddress:  rconAddress,
		RconPassword: rconPassword,
	}

	a, err := api.NewAPI(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
