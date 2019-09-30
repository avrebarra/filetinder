package server

import (
	"fmt"
	"net/http"

	"github.com/shrotavre/filetinder/internal/config"
)

// Start starts dirtinder main HTTP server
func Start() {
	appconf := config.GetConfigs()

	// http.HandleFunc("/api/targets", nil)
	// http.HandleFunc("/api/perform", nil)
	// http.HandleFunc("/api/meta", nil)

	err := http.ListenAndServe(fmt.Sprintf(":%d", appconf.Port), nil)
	if err != nil {
		panic(err)
	} else {
	}

}
