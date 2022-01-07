package queuesim

import "math/rand"

const (
	serveMinSeconds int = 7
	serveMaxSeconds int = 15
	clProduceMin    int = 7
	clProduceMax    int = 10
)

type Client struct {
	serveDuration int
	waitDuration  int
}

func ProduceClient() *Client {
	if rand.Intn(clProduceMax+1) > clProduceMin {
		return MustProduceClient()
	}

	return nil
}

func MustProduceClient() *Client {
	return &Client{
		serveDuration: serveMaxSeconds - rand.Intn((serveMaxSeconds-serveMinSeconds)+1),
	}
}
