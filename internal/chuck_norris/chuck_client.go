package chucknorris

type ChuckNorrisClient struct {
	ChuckService ChuckService
}

func NewChuckNorrisClient(cs ChuckService) *ChuckNorrisClient {
	return &ChuckNorrisClient{
		ChuckService: cs,
	}
}
