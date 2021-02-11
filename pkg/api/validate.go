package api

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetchDictionary() (io.Reader, error) {
	// TODO: fetch from DB (if available).

	file, err := os.Open("./dictionary.txt")
	if err == nil {
		return file, nil
	}

	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err == nil {
		return res.Body, nil
	}

	return nil, fmt.Errorf("unable to fetch dictionary: %v", err)
}

func (s *service) buildDictionary() error {
	dict, err := fetchDictionary()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(dict)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		s.dictionary.Add(word)
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}

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
