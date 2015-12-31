package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type API struct {
	config *APIConfig
}

type APIConfig struct {
	ListenAddr string
	PublicDir  string
}

func NewAPI(config *APIConfig) (*API, error) {
	return &API{
		config: config,
	}, nil
}

func (a *API) Run() error {
	globalMux := http.NewServeMux()

	r := mux.NewRouter()
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
