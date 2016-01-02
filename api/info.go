package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/factionlabs/meld/version"
)

type InstallInfo struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	SteamCMDPath string `json:"steamCMDPath"`
	RustPath     string `json:"rustPath"`
}

func (a *API) info(w http.ResponseWriter, r *http.Request) {
	steamPath, err := a.db.Get(dbBucketConfig, dbKeySteamCmdPath)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error getting configuration: %s", err), http.StatusInternalServerError)
		return
	}

	rustPath, err := a.db.Get(dbBucketConfig, dbKeyRustPath)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error getting configuration: %s", err), http.StatusInternalServerError)
		return
	}

	info := InstallInfo{
		Name:         version.FullName(),
		Version:      version.FullVersion(),
		SteamCMDPath: string(steamPath),
		RustPath:     string(rustPath),
	}

	if err := json.NewEncoder(w).Encode(info); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error encoding: %s", err), http.StatusInternalServerError)
		return
	}
}
