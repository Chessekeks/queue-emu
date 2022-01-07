package queuesim

import (
	"fmt"
)

type Simulator struct {
	Queues          []*Queue
	Duration        int
	Clients         []*Client
	dropClientCount int
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
	s.Clients = append(s.Clients, cl)

	for idx, q := range s.Queues {
		for name, val := range cl.rules {
			if v, ok := q.rules[name]; ok && v.Include(val) {
				s.Queues[idx].AddClient(cl)
				return
			}
		}
	}

	s.dropClientCount++
	//s.Queues[rand.Intn(len(s.Queues))].AddClient(cl)
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
		avrWaitTime := q.AvgWaitTime()
		totalServeTime += avrServeTime
		totalWaitTime += avrWaitTime

		q.PrintStats()
	}
	totalServeTime /= float32(len(s.Queues))
	totalWaitTime /= float32(len(s.Queues))

	fmt.Printf("Drop clients count is %d\n", s.dropClientCount)
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
		s.Clients = append(s.Clients, MakeClient())
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
