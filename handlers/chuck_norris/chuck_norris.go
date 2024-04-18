package chucknorris

import (
	"chuck_service/handlers/responses"
	chucknorris "chuck_service/internal/chuck_norris"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ChuckNorrisHandler struct {
	chuckClient chucknorris.Client
}

func NewChuckNorrisHandler(c chucknorris.Client) *ChuckNorrisHandler {
	return &ChuckNorrisHandler{
		chuckClient: c,
	}
}

func (h *ChuckNorrisHandler) GetRandomJoke(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling /random")

	category := r.URL.Query().Get("category")
	if category != "" {
		log.Println("Category query parameter is present:", category)
	} else {
		log.Println("Category query parameter is not present")
	}

	fields := log.Fields{"Category": category}

	joke, err := h.chuckClient.RandomJoke(category)
	if err != nil {
		log.WithFields(fields).Errorf("%+v", err)
		responses.NotFound404(w, "Jokes")
		return
	}

	responses.OK200(w, joke)
}
