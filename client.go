package queuesim

import "math/rand"

const (
	serveMinSeconds int = 7
	serveMaxSeconds int = 15
	clProduceMin    int = 7
	clProduceMax    int = 10
)

type Client struct {
	serveSeconds int
	waitSeconds  int

	rules map[string]int
}

func ProduceClient() *Client {
	if rand.Intn(clProduceMax+1) > clProduceMin {
		return MakeClient()
	}

	return nil
}

func MakeClient() *Client {
	rules := make(map[string]int)
	rules[WeightRuleName] = rand.Intn(9)

	return &Client{
		serveSeconds: serveMaxSeconds - rand.Intn((serveMaxSeconds-serveMinSeconds)+1),
		rules:        rules,
	}
}
