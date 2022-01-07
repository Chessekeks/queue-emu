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

	duration := rand.Intn(totalDurMax)
	q1 := queuesim.NewQueue("q1")
	q1.AddRule(queuesim.WeightRuleName, queuesim.NewWeightRule(1, 2, 3, 4))
	q2 := queuesim.NewQueue("q2")
	q2.AddRule(queuesim.WeightRuleName, queuesim.NewWeightRule(5, 6, 7, 8))

	simulator := queuesim.NewSimulator(q1, q2)
	simulator.SimulateByDuration(duration)
	simulator.PrintResults()
}
