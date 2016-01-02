package api

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/factionlabs/meld/utils"
)

func (a *API) installSteamCmd(w http.ResponseWriter, r *http.Request) {
	steamDir := r.URL.Query().Get("steampath")

	// default dirs if none specified
	if steamDir == "" {
		steamDir = utils.SteamDefaultDir
	}

	// check for existing
	_, err := utils.GetSteamCmdPath(steamDir)
	if err != nil {
		if os.IsNotExist(err) {
			// install
			if err := utils.InstallSteamCmd(steamDir); err != nil {
				log.Error(err)
				http.Error(w, fmt.Sprintf("error installing steam cmd: %s", err), http.StatusInternalServerError)
				return
			}

			if _, err := utils.GetSteamCmdPath(steamDir); err != nil {
				log.Error(err)
				http.Error(w, fmt.Sprintf("error finding steam cmd after install: %s", err), http.StatusInternalServerError)
				return
			}
		} else {
			log.Error(err)
			http.Error(w, fmt.Sprintf("error locating steam cmd: %s", err), http.StatusInternalServerError)
			return
		}
	}

	// update db
	if err := a.db.Set(dbBucketConfig, dbKeySteamCmdPath, []byte(steamDir)); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error updating configuration: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("steam cmd installed successfully"))
}

func (a *API) installRust(w http.ResponseWriter, r *http.Request) {
	steamDir := r.URL.Query().Get("steampath")
	rustDir := r.URL.Query().Get("rustpath")

	// default dirs if none specified
	if steamDir == "" {
		steamDir = utils.SteamDefaultDir
	}

	if rustDir == "" {
		rustDir = utils.RustDefaultDir
	}

	// check for existing
	steamCmdPath, err := utils.GetSteamCmdPath(steamDir)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error finding steam cmd: %s; perhaps try to install it", err), http.StatusInternalServerError)
		return
	}

	if err := utils.InstallRust(steamCmdPath, rustDir, false); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error installing rust: %s", err), http.StatusInternalServerError)
		return
	}

	// update db
	if err := a.db.Set(dbBucketConfig, dbKeyRustPath, []byte(rustDir)); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error updating configuration: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("rust installed successfully"))
}

func (a *API) updateRust(w http.ResponseWriter, r *http.Request) {
	steamDir := r.URL.Query().Get("steampath")
	rustDir := r.URL.Query().Get("rustpath")

	// default dirs if none specified
	if steamDir == "" {
		steamDir = utils.SteamDefaultDir
	}

	if rustDir == "" {
		rustDir = utils.RustDefaultDir
	}

	// check for existing
	steamCmdPath, err := utils.GetSteamCmdPath(steamDir)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error finding steam cmd: %s; perhaps try to install it", err), http.StatusInternalServerError)
		return
	}

	if err := utils.InstallRust(steamCmdPath, rustDir, true); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error updating rust: %s", err), http.StatusInternalServerError)
		return
	}

	// update db
	if err := a.db.Set(dbBucketConfig, dbKeyRustPath, []byte(rustDir)); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error updating configuration: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("rust updated successfully"))
}

func (a *API) installOxide(w http.ResponseWriter, r *http.Request) {
	rustDir := r.URL.Query().Get("rustpath")

	// default dirs if none specified
	if rustDir == "" {
		rustDir = utils.RustDefaultDir
	}

	if err := utils.InstallOxideMod(rustDir); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error installing oxide: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("oxide installed successfully"))
}
