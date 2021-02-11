package api

import (
	"net/http"

	"github.com/clarketm/boggle-api/pkg/trie"
	"github.com/clarketm/boggle-api/pkg/util"
)

type service struct {
	//db     *someDatabase
	dictionary *trie.Trie
	router     *http.ServeMux
	Log        *util.Logger
}

func NewService() *service {
	s := &service{
		dictionary: trie.NewTrie(),
		router:     http.NewServeMux(),
		Log:        util.NewLogger(),
	}
	return s
}

func (s *service) Configure() error {
	if err := s.buildDictionary(); err != nil {
		return err
	}
	s.routes()

	return nil
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
