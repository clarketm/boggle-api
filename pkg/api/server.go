package api

import (
	"net/http"

	"github.com/clarketm/boggle-api/pkg/trie"
	"github.com/gorilla/mux"
)

type Server struct {
	//db     *someDatabase
	Dictionary *trie.Trie
	Router     *mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router:     mux.NewRouter(),
		Dictionary: trie.NewTrie(),
	}
	s.buildDictionary()
	s.routes()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
