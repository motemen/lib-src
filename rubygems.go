package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type rubygemInfo struct {
	SourceCodeURI string `json:"source_code_uri"`
	HomepageURI   string `json:"homepage_uri"`
	BugTrackerURI string `json:"bug_tracker_uri,"`
}

func fetchRubyGems(libName string) (*libInfo, error) {
	url := fmt.Sprintf("https://rubygems.org/api/v1/gems/%s.json", libName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	info := &rubygemInfo{}
	err = json.NewDecoder(resp.Body).Decode(info)
	if err != nil {
		return nil, err
	}

	return &libInfo{
		source:     info.SourceCodeURI,
		homepage:   info.HomepageURI,
		bugtracker: info.BugTrackerURI,
	}, nil
}
