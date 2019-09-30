package server

import (
	"fmt"
	"net/http"

	"github.com/shrotavre/filetinder/internal/config"
	"github.com/shrotavre/filetinder/internal/server/handlers"
)

// Start starts FileTinder main HTTP server
func Start() error {
	appconf := config.GetConfigs()

	http.HandleFunc("/api/targets", handlers.HandleAPITarget)
	// http.HandleFunc("/api/perform", nil)
	http.HandleFunc("/api/meta", handlers.HandleAPIMeta)

	return http.ListenAndServe(fmt.Sprintf(":%d", appconf.Port), nil)
}
