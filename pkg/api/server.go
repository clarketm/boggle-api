package api

import (
	"net/http"

	"github.com/clarketm/boggle-api/pkg/trie"
)

type Server struct {
	//db     *someDatabase
	Dictionary *trie.Trie
	Router     *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		Router:     http.NewServeMux(),
		Dictionary: trie.NewTrie(),
	}
	s.buildDictionary()
	s.routes()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
