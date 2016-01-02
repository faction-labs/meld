package rust

import (
	"fmt"
	"strings"
)

func (r *RustServer) Hostname() (string, error) {
	host, err := r.RunCommand("server.hostname")
	if err != nil {
		return "", err
	}

	parts := strings.Split(host, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("unable to determine hostname")
	}

	h := parts[1]

	h = strings.Replace(parts[1], "\"", "", -1)
	h = strings.TrimSpace(h)

	return h, nil
}
