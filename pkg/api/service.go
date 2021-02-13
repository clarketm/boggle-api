package api

import (
	"net/http"

	"github.com/clarketm/boggle-api/pkg/db"
	"github.com/clarketm/boggle-api/pkg/trie"
	"github.com/clarketm/boggle-api/pkg/util"
)

type service struct {
	db         *db.DB
	dictionary *trie.Trie
	router     *http.ServeMux
	log        *util.Logger
}

func NewService(cfg *util.Config, logger *util.Logger) (*service, error) {
	dbConnect, err := db.NewDB(cfg.DbEnable, cfg.DbUser, cfg.DbPassword, cfg.DbHost)
	if err != nil {
		return nil, err
	}

	s := &service{
		db:         dbConnect,
		dictionary: trie.NewTrie(),
		router:     http.NewServeMux(),
		log:        logger,
	}

	if err = s.Configure(cfg); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *service) Configure(cfg *util.Config) error {
	if err := s.initDB(cfg); err != nil {
		return err
	}
	if err := s.buildDictionary(cfg); err != nil {
		return err
	}
	s.routes()

	return nil
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
