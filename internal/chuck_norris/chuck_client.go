package chucknorris

import "net/http"

type Client interface {
	RandomJoke(category string) (*ChuckJoke, error)
	Categories() (*[]string, error)
}

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type ChuckNorrisService struct {
	httpClient HTTPClient
}

func NewChuckNorrisService(httpClient HTTPClient) *ChuckNorrisService {
	return &ChuckNorrisService{
		httpClient: httpClient,
	}
}
