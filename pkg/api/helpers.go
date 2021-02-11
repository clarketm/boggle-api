package api

import (
	"encoding/json"
	"net/http"
)

func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		//err := json.NewEncoder(w).Encode(data)
		json.NewEncoder(w).Encode(data)
	}

}
