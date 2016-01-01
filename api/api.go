package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	mdb "github.com/factionlabs/meld/db"
	"github.com/gorilla/mux"
)

const (
	dbBucketConfig    = "config"
	dbKeySteamCmdPath = "config.steamcmd.path"
	dbKeyRustPath     = "config.rust.path"
)

type API struct {
	config *APIConfig
	db     *mdb.DB
}

type APIConfig struct {
	ListenAddr string
	PublicDir  string
	DBPath     string
}

func NewAPI(config *APIConfig) (*API, error) {
	return &API{
		config: config,
	}, nil
}

func (a *API) Run() error {
	dbConfig := &mdb.DBConfig{
		DBPath: a.config.DBPath,
	}
	bdb, err := mdb.NewDB(dbConfig)
	if err != nil {
		return err
	}

	a.db = bdb

	globalMux := http.NewServeMux()

	r := mux.NewRouter()
	r.HandleFunc("/api/version", a.version)
	r.HandleFunc("/api/info", a.info)
	r.HandleFunc("/api/install/steam", a.installSteamCmd)
	r.HandleFunc("/api/install/rust", a.installRust)
	r.HandleFunc("/api/install/oxide", a.installOxide)
	r.HandleFunc("/api/update/rust", a.updateRust)

	// static handler
	globalMux.Handle("/", http.FileServer(http.Dir(a.config.PublicDir)))
	globalMux.Handle("/api/", r)

	s := &http.Server{
		Addr:    a.config.ListenAddr,
		Handler: globalMux,
	}

	log.Infof("api serving: addr=%s", a.config.ListenAddr)

	return s.ListenAndServe()
}
