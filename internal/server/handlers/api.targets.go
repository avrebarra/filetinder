package handlers

import (
	"encoding/json"
	"net/http"
)

type target struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
	Tag string `json:"tag"`
}

type targets []target

var targetStore targets
var targetLastIndex int64

// HandleAPITarget handler function for '/api/targets' endpoints
func HandleAPITarget(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t := target{}

		// send response
		j, _ := json.Marshal(t)
		w.Write(j)
		return

	case "POST":
		var t target

		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		t.ID = targetLastIndex
		t.Tag = ""

		targetStore = append(targetStore, t)
		targetLastIndex++

		// send response
		j, _ := json.Marshal(targetStore)
		w.Write(j)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
