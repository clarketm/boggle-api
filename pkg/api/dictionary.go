package api

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func (s *service) fetchDictionary() (io.Reader, error) {
	words, err := s.getAllWords()
	if err == nil {
		return words, nil
	}

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
	dict, err := s.fetchDictionary()
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
