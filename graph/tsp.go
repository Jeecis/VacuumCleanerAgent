package graph

import (
	"log"
	"math"
)

func tspGraphInit(g *Graph, keyNodes []Node, startNode *Node) *Graph {
	distilledGraph := &Graph{
		AdjacencyList: make(map[Node][]Node),
		Edges:         make(map[Node]map[Node]Path),
	}

	nodes := append(keyNodes, *startNode)
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			pathNodes, cost := g.bfs(nodes[i], nodes[j])
			if cost < math.Inf(1) {
				reversedPath := reverseArray(pathNodes)

				// This is just a bug, which for now is fixed by adding this if statement
				log.Print("Nodes: ", nodes[i], nodes[j], " Cost: ", cost, " Path: ", pathNodes)
				if pathNodes[0] != nodes[i] {
					distilledGraph.addEdge(nodes[i], nodes[j], cost, reversedPath)
					distilledGraph.addEdge(nodes[j], nodes[i], cost, pathNodes)
				} else {
					distilledGraph.addEdge(nodes[i], nodes[j], cost, pathNodes)
					distilledGraph.addEdge(nodes[j], nodes[i], cost, reversedPath)
				}
			}
		}
	}

	distilledGraph.printGraphEdges()

	return distilledGraph
}

func reverseArray(arr []Node) []Node {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}

	return arr

}
