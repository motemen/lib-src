package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type npmInfo struct {
	Repository struct {
		URL string
	}
	Homepage string
	Bugs     struct {
		URL string
	}
}

func fetchNPM(libName string) (*libInfo, error) {
	url := fmt.Sprintf("https://registry.npmjs.org/%s", libName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	info := &npmInfo{}
	err = json.NewDecoder(resp.Body).Decode(info)
	if err != nil {
		return nil, err
	}

	return &libInfo{
		source:     info.Repository.URL,
		homepage:   info.Homepage,
		bugtracker: info.Bugs.URL,
	}, nil
}
