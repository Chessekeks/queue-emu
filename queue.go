package queuesim

import "fmt"

type Queue struct {
	name string

	curClient              *Client
	curClientServeDuration int
	waiters                []*Client

	serveDuration float32
	serveCount    int
	waitDuration  float32
	waitCount     int
}

func NewQueue(name string) *Queue {
	return &Queue{
		name: name,
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

func (q *Queue) Tick() {
	if len(q.waiters) == 0 {
		return
	}

	curClient := q.curClient

	if curClient == nil {
		curClient = q.waiters[0]
		q.waiters = q.waiters[1:]
	}

	q.curClientServeDuration++
	for idx := range q.waiters {
		q.waiters[idx].waitDuration++
	}
	if curClient.serveDuration == q.curClientServeDuration {
		q.serveDuration += float32(q.curClientServeDuration)
		q.serveCount++
		q.waitDuration += float32(q.curClient.waitDuration)
		q.waitCount++
		q.curClientServeDuration = 0
		curClient = nil
	}
	q.curClient = curClient
}

func (q Queue) AvgServeTime() float32 {
	return q.serveDuration / float32(q.serveCount)
}

func (q Queue) PrintAvgServeTime() {
	fmt.Printf("Queue %s avarage serve time is %.2f seconds\n", q.name, q.AvgServeTime())
}

func (q Queue) AvgWaitTime() float32 {
	return q.waitDuration / float32(q.waitCount)
}

func (q Queue) PrintAvgWaitTime() {
	fmt.Printf("Queue %s avarage wait time is %.2f seconds\n", q.name, q.AvgWaitTime())
}
