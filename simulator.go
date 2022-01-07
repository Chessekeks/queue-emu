package queuesim

import (
	"fmt"
	"math/rand"
)

type Simulator struct {
	Queues   []*Queue
	Duration int
	Clients  []*Client
}

func NewSimulator(queues ...*Queue) *Simulator {
	return &Simulator{
		Queues: queues,
	}
}

func (s Simulator) AllQueuesFinished() bool {
	for _, q := range s.Queues {
		if !q.Finished() {
			return false
		}
	}

	return true
}

func (s *Simulator) AddClient(cl *Client) {
	if cl == nil {
		return
	}

	s.Queues[rand.Intn(len(s.Queues))].AddClient(cl)
	s.Clients = append(s.Clients, cl)
}

func (s *Simulator) Tick() {
	for qIdx := range s.Queues {
		s.Queues[qIdx].Tick()
	}
}

func (s Simulator) PrintResults() {
	fmt.Printf("Simulation time: %d seconds\n", s.Duration)
	fmt.Printf("Simulation client counts: %d\n", len(s.Clients))

	var totalServeTime, totalWaitTime float32
	for _, q := range s.Queues {
		avrServeTime := q.AvgServeTime()
		q.PrintAvgServeTime()
		avrWaitTime := q.AvgWaitTime()
		q.PrintAvgWaitTime()
		q.PrintServeCount()
		q.PrintWaitersCount()
		totalServeTime += avrServeTime
		totalWaitTime += avrWaitTime
	}
	totalServeTime /= float32(len(s.Queues))
	totalWaitTime /= float32(len(s.Queues))

	fmt.Printf("Total avarage serve time is %.2f seconds\n", totalServeTime)
	fmt.Printf("Total avarage wait time is %.2f seconds\n", totalWaitTime)
}

func (s *Simulator) SimulateByDuration(dur int) {
	s.Duration = dur

	for idx := 0; idx < s.Duration; idx++ {
		cl := ProduceClient()
		s.AddClient(cl)
		s.Tick()
	}
}

func (s *Simulator) SimulateByClients(clientCount int) {
	for idx := 0; idx < clientCount; idx++ {
		s.Clients = append(s.Clients, MustProduceClient())
	}

	dur := 0
	for _, cl := range s.Clients {
		s.AddClient(cl)
		s.Tick()
		dur++
	}

	for !s.AllQueuesFinished() {
		s.Tick()
		dur++
	}

	s.Duration = dur
}
