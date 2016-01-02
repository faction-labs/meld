package api

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func (a *API) rustExec(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query().Get("cmd")

	if cmd == "" {
		http.Error(w, "you must specify a command", http.StatusBadRequest)
		return
	}

	resp, err := a.rustServer.Exec(cmd)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error running rcon command: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(resp))
}
