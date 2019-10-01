package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/shrotavre/filetinder/internal/config"
)

type meta struct {
	Port int `json:"port"`
	PID  int `json:"pid"`
}

var appmeta meta

func init() {
	appmeta.Port = config.DefaultPort
	appmeta.PID = os.Getpid()
}

// HandleAPIMeta handler function for '/api/meta' endpoints
func HandleAPIMeta(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		j, _ := json.Marshal(appmeta)
		w.Write(j)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
