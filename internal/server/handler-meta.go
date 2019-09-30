package server

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
	appconf := config.GetConfigs()

	appmeta.Port = appconf.Port
	appmeta.PID = os.Getpid()
}

func metaHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		j, _ := json.Marshal(appmeta)
		w.Write(j)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
