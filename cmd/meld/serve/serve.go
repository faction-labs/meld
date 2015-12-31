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
			Name:  "public-dir, s",
			Usage: "path to public media directory",
			Value: "public",
		},
	},
}

func serveAction(c *cli.Context) {
	log.Infof("%s %s", version.FullName(), version.FullVersion())

	listenAddr := c.String("listen")
	publicDir := c.String("public-dir")

	cfg := &api.APIConfig{
		ListenAddr: listenAddr,
		PublicDir:  publicDir,
	}

	a, err := api.NewAPI(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
