package main

import (
	"github.com/lazyboson/assignmentproblem/graph"
	"testing"
)

var (
	maxAgents = 15000
	maxTasks  = 1500
)

/*
Benchmarking on the worst possible dense graph with 15000*15000 ~ 22,50,00,000 ~ 225 Million edges
Running Time of Algorithm ~ 2.36 Second
*/
func BenchmarkHopcroftKarp(b *testing.B) {
	g := graph.InitGraph(maxAgents, maxTasks)
	for i := 1; i <= maxAgents; i++ {
		for j := 1; j <= maxTasks; j++ {
			g.AddEdges(i, j)
		}
	}

	g.HopcroftKarp()
}
