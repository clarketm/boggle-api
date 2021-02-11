package api

func (s *Server) routes() {
	s.Router.HandleFunc("/api/health", s.handleHealth())
	s.Router.HandleFunc("/api/validate", s.handleValidateWord())
}
