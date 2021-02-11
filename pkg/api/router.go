package api

func (s *Server) routes() {
	s.Router.HandleFunc("/api/validate", s.handleValidateWord())
}
