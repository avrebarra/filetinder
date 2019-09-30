package server

import (
	"fmt"
	"net/http"

	"github.com/shrotavre/filetinder/internal/config"
)

// Start starts dirtinder main HTTP server
func Start() error {
	appconf := config.GetConfigs()

	// http.HandleFunc("/api/targets", nil)
	// http.HandleFunc("/api/perform", nil)
	http.HandleFunc("/api/meta", metaHandler)

	return http.ListenAndServe(fmt.Sprintf(":%d", appconf.Port), nil)
}
