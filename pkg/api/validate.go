package api

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func fetchDictionary() io.Reader {
	file, err := os.Open("./Dictionary.txt")
	if err == nil {
		return file
	}

	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err == nil {
		return res.Body
	}

	return nil
}

func (s *Server) buildDictionary() {
	dict := fetchDictionary()
	if dict == nil {
		log.Fatal("Unable to build dictionary.")
	}

	scanner := bufio.NewScanner(dict)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		s.Dictionary.Add(word)
	}

	//err := scanner.Err()
}

func (s *Server) handleValidateWord() http.HandlerFunc {
	type request struct {
		Word string `json:"word"`
	}
	type response struct {
		Valid bool `json:"valid"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
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
		valid := s.Dictionary.Search(word)
		res := response{Valid: valid}

		s.respond(w, r, res, http.StatusOK)
	}
}
