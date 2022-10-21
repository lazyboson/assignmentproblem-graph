package main

import "github.com/lazyboson/assignmentproblem/graph"

func main() {
	g := graph.InitGraph(6, 6)
	g.AddEdges(1, 2)
	g.AddEdges(1, 3)
	g.AddEdges(3, 1)
	g.AddEdges(3, 4)
	g.AddEdges(4, 3)
	g.AddEdges(5, 3)
	g.AddEdges(5, 4)
	g.AddEdges(6, 6)
	g.HopcroftKarp()
}
