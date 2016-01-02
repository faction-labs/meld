package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func (a *API) rustStatus(w http.ResponseWriter, r *http.Request) {
	status, err := a.rustServer.Status()
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error getting server status: %s", err), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(status); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error encoding: %s", err), http.StatusInternalServerError)
		return
	}
}
