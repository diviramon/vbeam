package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type pinpoint struct {
	RootCID   string    `json:"cid"`
	CreatedAt time.Time `json:"createdAt"`
}

func scanTopic(URL string) (*pinpoint, error) {
	resp, err := http.Get("http://" + URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	p := new(pinpoint)
	err = json.Unmarshal(body, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}
