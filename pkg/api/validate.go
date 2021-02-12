package api

import (
	"net/http"
	"strings"
)

func (s *service) handleValidateWord() http.HandlerFunc {
	type request struct {
		Word string `json:"word"`
	}
	type response struct {
		Valid bool `json:"valid"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			s.respond(w, r, nil, http.StatusMethodNotAllowed)
			return
		}

		req := request{}
		err := s.decode(w, r, &req)

		if err != nil {
			s.respond(w, r, nil, http.StatusUnprocessableEntity)
			return
		}

		word := strings.ToLower(req.Word)
		valid := s.dictionary.Search(word)
		res := response{Valid: valid}

		s.respond(w, r, res, http.StatusOK)
	}
}
