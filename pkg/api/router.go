package api

func (s *server) routes() {
	s.router.HandleFunc("/api/health", s.handleHealth())
	s.router.HandleFunc("/api/validate", s.handleValidateWord())
}
