package api

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/factionlabs/meld/utils/steam"
	"github.com/gorilla/mux"
)

func (a *API) installSteamCmd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	steamDir := vars["path"]

	// check for existing
	_, err := steam.GetSteamCmdPath(steamDir)
	if err != nil {
		if os.IsNotExist(err) {
			// install
			if err := steam.InstallSteamCmd(steamDir); err != nil {
				log.Error(err)
				http.Error(w, fmt.Sprintf("error installing steam cmd: %s", err), http.StatusInternalServerError)
				return
			}

			if _, err := steam.GetSteamCmdPath(steamDir); err != nil {
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

	w.Write([]byte("steam cmd installed successfully"))
}

func (a *API) installRust(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	steamDir := vars["steampath"]
	rustDir := vars["path"]

	// check for existing
	steamCmdPath, err := steam.GetSteamCmdPath(steamDir)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error finding steam cmd: %s; perhaps try to install it", err), http.StatusInternalServerError)
		return
	}

	if err := steam.InstallRust(steamCmdPath, rustDir, false); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error installing rust: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("rust installed successfully"))
}

func (a *API) updateRust(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	steamDir := vars["steampath"]
	rustDir := vars["path"]

	// check for existing
	steamCmdPath, err := steam.GetSteamCmdPath(steamDir)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error finding steam cmd: %s; perhaps try to install it", err), http.StatusInternalServerError)
		return
	}

	if err := steam.InstallRust(steamCmdPath, rustDir, true); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error updating rust: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("rust updated successfully"))
}

func (a *API) installOxide(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rustDir := vars["path"]

	if err := steam.InstallOxideMod(rustDir); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error installing oxide: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("oxide installed successfully"))
}
