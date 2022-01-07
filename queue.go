package queuesim

import "fmt"

type Queue struct {
	name string

	curClient              *Client
	curClientServeDuration int
	waiters                []*Client
	rules                  map[string]Rule

	serveDuration float32
	waitDuration  float32
	clientCount   int
}

func NewQueue(name string) *Queue {
	return &Queue{
		name:  name,
		rules: map[string]Rule{},
	}
}

func (q Queue) Finished() bool {
	return q.curClient == nil && len(q.waiters) == 0
}

func (q *Queue) AddClient(cl *Client) {
	if cl == nil {
		return
	}

	q.waiters = append(q.waiters, cl)
}

func (q *Queue) AddRule(name string, rule Rule) {
	q.rules[name] = rule
}

func (q *Queue) Tick() {
	if q.Finished() {
		return
	}

	curClient := q.curClient

	if curClient == nil && len(q.waiters) > 0 {
		curClient = q.waiters[0]
		q.waiters = q.waiters[1:]
	}

	q.curClientServeDuration++
	for idx := range q.waiters {
		q.waiters[idx].waitSeconds++
	}
	if curClient.serveSeconds == q.curClientServeDuration {
		q.serveDuration += float32(q.curClientServeDuration)
		q.waitDuration += float32(q.curClient.waitSeconds)
		q.clientCount++
		q.curClientServeDuration = 0
		curClient = nil
	}
	q.curClient = curClient
}

// Statistic methods

func (q Queue) AvgServeTime() float32 {
	return q.serveDuration / float32(q.clientCount)
}

func (q Queue) AvgWaitTime() float32 {
	return q.waitDuration / float32(q.clientCount)
}

func (q Queue) PrintStats() {
	fmt.Printf("Queue %s avarage serve time is %.2f seconds\n", q.name, q.AvgServeTime())
	fmt.Printf("Queue %s avarage wait time is %.2f seconds\n", q.name, q.AvgWaitTime())
	fmt.Printf("Queue %s served %d clients\n", q.name, q.clientCount)
	if len(q.waiters) > 0 {
		fmt.Printf("Queue %s has %d waiting clients\n", q.name, len(q.waiters))
	}
}
