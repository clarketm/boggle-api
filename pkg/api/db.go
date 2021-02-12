package api

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"

	"github.com/clarketm/boggle-api/pkg/db"
)

func (s *service) getAllWords() (io.Reader, error) {
	var row []byte
	buf := bytes.Buffer{}

	rows, err := s.db.Query("SELECT word from db.dictionary;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&row); err != nil {
			return nil, err
		}
		buf.Write(row)
		buf.WriteByte('\n')
	}

	return &buf, nil
}

func (s *service) initDB() error {
	// Create database `db`.
	createDbQuery := `CREATE DATABASE IF NOT EXISTS db;`
	_, err := s.db.Execute(createDbQuery, nil)
	if err != nil {
		return err
	}

	// Drop table `dictionary`.
	//dropDbQuery := `DROP TABLE db.dictionary;`
	//_, err = s.db.Execute(dropDbQuery, nil)
	//if err != nil {
	//	return err
	//}

	// Create table `dictionary`.
	createTableQuery := `CREATE TABLE IF NOT EXISTS db.dictionary(word VARCHAR(50) NOT NULL PRIMARY KEY);`
	_, err = s.db.Execute(createTableQuery, nil)
	if err != nil {
		return err
	}

	file, err := os.Open("./dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	insertWordTmpl := `INSERT IGNORE INTO db.dictionary VALUES `
	placeholders := []string{}
	words := []interface{}{}

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		placeholders = append(placeholders, "(?)")
		words = append(words, word)

		if len(words) == db.BatchSize {
			_, err = s.db.Execute(insertWordTmpl+strings.Join(placeholders, ","), words)
			if err != nil {
				return err
			}

			placeholders = placeholders[:0]
			words = words[:0]
		}
	}

	_, err = s.db.Execute(insertWordTmpl+strings.Join(placeholders, ","), words)
	if err != nil {
		return err
	}

	return nil
}
