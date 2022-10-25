package main

import (
	"github.com/lazyboson/assignmentproblem/graph"
	"testing"
)

func getGraph(agents, tasks int) *graph.Graph {
	g := graph.InitGraph(agents, tasks)
	for i := 1; i <= agents; i++ {
		for j := 1; j <=tasks; j++ {
			g.AddEdges(i, j)
		}
	}
	return g
}

var table = []struct {
	g *graph.Graph
} {
	{g :getGraph(1000, 2000)}, 
	{g: getGraph(2000, 3000)},
	{g: getGraph(3000, 4000)}, 
	{g: getGraph(5000, 5000)}, 
	{g: getGraph(6000, 5000)}, 
	{g: getGraph(7000, 4000)}, 
	{g: getGraph(5000, 8000)}, 
	{g: getGraph(5000, 9000)}, 
	{g : getGraph(10000, 10000)},
}


/*
Benchmarking on the worst possible dense graph with 15000*15000 ~ 22,50,00,000 ~ 225 Million edges
Running Time of Algorithm ~ 2.36 Second
*/
func BenchmarkHopcroftKarp(b *testing.B) {
	for _, val := range table {
		for n := 0; n < b.N; n++ {
			val.g.HopcroftKarp()
		}
	}	
}
