package main

import "github.com/lazyboson/assignmentproblem/graph"

func main() {
	g := graph.InitGraph(4, 4)
	g.AddEdges(1, 2)
	g.AddEdges(1, 3)
	g.AddEdges(2, 1)
	g.AddEdges(3, 2)
	g.AddEdges(4, 2)
	g.AddEdges(4, 4)
	g.HopcroftKart()
}
