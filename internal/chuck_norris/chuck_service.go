package chucknorris

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ChuckFunc interface {
	RandomJoke(category string) (*ChuckJoke, error)
}

type ChuckService struct {
}

type ChuckJoke struct {
	Categories []string `json:"categroies"`
	CreatedAt  string   `json:"created_at"`
	IconUrl    string   `json:"icon_url"`
	Id         string   `json:"id"`
	Url        string   `json:"url"`
	Value      string   `json:"value"`
}

var (
	root = "https://api.chucknorris.io/jokes/"
)

func (c *ChuckService) RandomJoke(category string) (*ChuckJoke, error) {
	var endpoint string
	var fields log.Fields
	if category == "" {
		endpoint = root + "random"
	} else {
		fields = log.Fields{"Category": category}
		endpoint = fmt.Sprintf("%srandom?category=%s", root, category)
	}

	response, err := http.Get(endpoint)
	if err != nil {
		log.WithFields(fields).Errorf("%+v", err)
		return nil, err
	}
	defer response.Body.Close()

	var chuckJoke ChuckJoke
	if err := json.NewDecoder(response.Body).Decode(&chuckJoke); err != nil {
		return nil, err
	}

	return &chuckJoke, nil
}

func (c *ChuckService) Categories() (*[]string, error) {
	endpoint := root + "categories"

	response, err := http.Get(endpoint)
	if err != nil {
		log.Errorf("Failed to fetch categories: %+v", err)
		return nil, err
	}
	defer response.Body.Close()

	var categories []string
	if err := json.NewDecoder(response.Body).Decode(&categories); err != nil {
		return nil, err
	}

	return &categories, nil
}
