/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package data

import (
	"aoc2021/secret"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	Year     = 2022
	URL      = "https://adventofcode.com/%d/day/%d"
	Filename = "data/day%d"
)

type Configuration struct {
	SessionCookie string
	Output        string
	Year          int
	Day           int
	Level         int
	Answer        int
}

func Post(day, level, result int) error {
	config := &Configuration{
		SessionCookie: secret.SessionID,
		Year:          Year,
		Day:           day,
		Level:         level,
		Answer:        result,
	}
	_, err := query(http.MethodPost, config)
	if err != nil {
		return err
	}
	return nil
}

func Get(day int, filename string) error {
	config := &Configuration{
		SessionCookie: secret.SessionID,
		Output:        filename,
		Year:          Year,
		Day:           day,
	}

	file, err := os.Open(config.Output)
	if err == nil {
		file.Close()
		return nil
	}

	resp, err := query(http.MethodGet, config)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	file, err = os.Create(config.Output)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	log.Printf("read data from aoc and wrote %d bytes to disk\n", bytes)
	return nil
}

func query(method string, config *Configuration) (*http.Response, error) {
	client := new(http.Client)

	path := "/input"
	var body io.Reader
	if method == http.MethodPost {
		path = "/answer"
		form := url.Values{}
		form.Add("level", strconv.Itoa(config.Level))
		form.Add("answer", strconv.Itoa(config.Answer))
		body = strings.NewReader(form.Encode())
	}

	req, err := http.NewRequest(method, fmt.Sprintf(URL+path, config.Year, config.Day), body)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"User-Agent": []string{"Advent of code solver"},
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: config.SessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return resp, nil
}
