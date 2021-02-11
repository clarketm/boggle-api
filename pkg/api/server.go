package api

import (
	"net/http"

	"github.com/clarketm/boggle-api/pkg/trie"
)

type server struct {
	//db     *someDatabase
	dictionary *trie.Trie
	router     *http.ServeMux
}

func NewServer() *server {
	s := &server{
		router:     http.NewServeMux(),
		dictionary: trie.NewTrie(),
	}
	s.buildDictionary()
	s.routes()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
