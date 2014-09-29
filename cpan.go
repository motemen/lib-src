package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type metacpanModuleInfo struct {
	Release struct {
		Source struct {
			Resources struct {
				Repository struct {
					Web string
				}
				Homepage   string
				Bugtracker struct {
					Web string
				}
			}
		} `json:"_source"`
	}
}

func fetchCPAN(libName string) (*libInfo, error) {
	url := fmt.Sprintf("http://api.metacpan.org/v0/module/%s?join=release", libName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	info := &metacpanModuleInfo{}
	err = json.NewDecoder(resp.Body).Decode(info)
	if err != nil {
		return nil, err
	}

	return &libInfo{
		source:     info.Release.Source.Resources.Repository.Web,
		homepage:   info.Release.Source.Resources.Homepage,
		bugtracker: info.Release.Source.Resources.Bugtracker.Web,
	}, nil
}
