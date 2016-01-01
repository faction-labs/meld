package api

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/factionlabs/meld/version"
)

func (a *API) info(w http.ResponseWriter, r *http.Request) {
	steamPath, err := a.db.Get(dbBucketConfig, dbKeySteamCmdPath)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error getting configuration: %s", err), http.StatusInternalServerError)
		return
	}

	if steamPath == nil {
		steamPath = []byte("not installed")
	}

	rustPath, err := a.db.Get(dbBucketConfig, dbKeyRustPath)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error getting configuration: %s", err), http.StatusInternalServerError)
		return
	}

	if rustPath == nil {
		rustPath = []byte("not installed")
	}

	info := fmt.Sprintf(`
%s %s
SteamCMD Path: %s
Rust Path: %s
`, version.FullName(), version.FullVersion(), steamPath, rustPath)

	w.Write([]byte(info))
}
