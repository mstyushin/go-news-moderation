package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ModerationRequest struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

func (api *API) check(w http.ResponseWriter, r *http.Request) {
	var mr ModerationRequest
	err := json.NewDecoder(r.Body).Decode(&mr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, badWord := range api.cfg.BadWords {
		if strings.Contains(mr.Author, badWord) || strings.Contains(mr.Text, badWord) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
