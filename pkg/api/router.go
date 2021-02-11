package api

import "net/http"

func (s *Server) routes() {
	api := s.Router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/validate", s.postValidateWord()).
		Methods(http.MethodPost)
}
