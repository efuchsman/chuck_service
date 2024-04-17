package chucknorris

type TestClient struct {
	RandomJokeData *ChuckJoke
	RandomJokeErr  error

	CategoriesData *[]string
	CategoriesErr  error
}

func (c TestClient) RandomJoke(category string) (*ChuckJoke, error) {
	return c.RandomJokeData, c.RandomJokeErr
}

func (c TestClient) Categories() (*[]string, error) {
	return c.CategoriesData, c.CategoriesErr
}
