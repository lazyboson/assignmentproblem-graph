package graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	agents int
	tasks  int
	adj    [][]int
	pairU  []int
	pairV  []int
	dist   []int
	curr   int
}

var (
	baseIndex = 0
	INF       = 1<<63 - 1
)

func InitGraph(agentCount, taskCount int) *Graph {
	adjList := make([][]int, agentCount+1)
	for i := 0; i < agentCount+1; i++ {
		adjList[i] = make([]int, 0)
	}

	return &Graph{
		agents: agentCount,
		tasks:  taskCount,
		adj:    adjList,
	}
}

// AddEdges  Adding directed edges
func (g *Graph) AddEdges(src, dest int) {
	g.adj[src] = append(g.adj[src], dest)
}

func (g *Graph) dfs(u int) bool {
	if u != 0 {
		for _, val := range g.adj[u] {
			if g.dist[g.pairV[val]] == g.dist[u]+1 {
				if g.dfs(g.pairV[val]) {
					g.pairV[val] = u
					g.pairU[u] = val
					g.curr = val
					return true
				}
			}
		}
		// there is no augmenting path from u so setting distance to infinite
		g.dist[u] = INF

		return false
	}

	return true
}

func (g *Graph) bfs() bool {
	queue := list.New()

	for i := 0; i < g.agents; i++ {
		if g.pairU[i] == 0 {
			g.dist[i] = 0
			queue.PushBack(i)
		} else {
			g.dist[i] = INF
		}
	}

	g.dist[baseIndex] = INF

	for queue.Len() > 0 {
		head := queue.Front()
		queue.Remove(head)
		if g.dist[head.Value.(int)] < g.dist[baseIndex] {
			for _, val := range g.adj[head.Value.(int)] {
				if g.dist[g.pairV[val]] == INF {
					g.dist[g.pairV[val]] = g.dist[head.Value.(int)] + 1
					queue.PushBack(g.pairV[val])
				}
			}
		}
	}

	return g.dist[baseIndex] != INF
}

// HopcroftKarp This is Printing Maximum bipartite matching -- exact assignment can be also printed
// Running complexity is O(edges*sqrt(vertices)) -- which is quit fast
func (g *Graph) HopcroftKarp() map[int]int {
	g.pairU = make([]int, g.agents+1)
	g.pairV = make([]int, g.tasks+1)
	g.dist = make([]int, g.agents+1)
	for index, _ := range g.pairU {
		g.pairU[index] = 0
	}
	for index, _ := range g.pairV {
		g.pairV[index] = 0
	}

	matchingSet := make(map[int]int)
	result := 0
	for g.bfs() {
		for i := 1; i <= g.agents; i++ {
			if g.pairU[i] == 0 && g.dfs(i) {
				matchingSet[i] = g.curr
				result += 1
			}
		}
	}

	fmt.Printf("Maximum Cardinality Matching for given graph :%d\n", result)

	return matchingSet
}
