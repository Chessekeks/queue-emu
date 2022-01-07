package main

import (
	"math/rand"
	"queuesim"
	"time"
)

const (
	totalDurMax = 5000
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Добавить реализацию правил(фильтры) для очереди - клиент должен подходить по всем нужным параметрам, иначе он уходит

	duration := rand.Intn(totalDurMax)
	q1 := queuesim.NewQueue("q1")
	q2 := queuesim.NewQueue("q2")

	simulator := queuesim.NewSimulator(q1, q2)

	simulator.SimulateByDuration(duration)

	simulator.PrintResults()
}
