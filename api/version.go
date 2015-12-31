package api

import (
	"fmt"
	"net/http"

	"github.com/factionlabs/meld/version"
)

func (a *API) version(w http.ResponseWriter, r *http.Request) {
	info := fmt.Sprintf("%s %s\n", version.FullName(), version.FullVersion())
	w.Write([]byte(info))
}
