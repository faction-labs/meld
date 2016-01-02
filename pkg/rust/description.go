package rust

import (
	"fmt"
	"strings"
)

func (r *RustServer) Description() (string, error) {
	host, err := r.RunCommand("server.description")
	if err != nil {
		return "", err
	}

	parts := strings.Split(host, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("unable to determine description")
	}

	h := parts[1]

	h = strings.Replace(parts[1], "\"", "", -1)
	h = strings.TrimSpace(h)

	return h, nil
}
