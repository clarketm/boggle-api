package api

import (
	"net/http"
)

func (s *server) handleHealth() http.HandlerFunc {
	type response struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Status  Status `json:"status"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			s.respond(w, r, nil, http.StatusMethodNotAllowed)
			return
		}

		res := response{
			Name:    Name,
			Version: Version,
			Status:  Up,
		}

		s.respond(w, r, res, http.StatusOK)
	}
}
