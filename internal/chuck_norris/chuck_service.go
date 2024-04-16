package chucknorris

type ChuckResponse interface {
}

type ChuckService struct {
}

type Chuck struct {
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
