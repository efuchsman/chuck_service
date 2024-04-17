package chucknorris

import "net/http"

type HTTPTestClient struct {
	GetData *http.Response
	GetErr  error
}

func (h *HTTPTestClient) Get(url string) (*http.Response, error) {
	return h.GetData, h.GetErr
}
