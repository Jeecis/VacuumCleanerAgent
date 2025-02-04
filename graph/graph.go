package graph

import "log"

type Node struct {
	IsWall      bool
	XCoordinate int
	YCoordinate int
	DirtCount   int
}

type Path struct {
	Weight float64
	Nodes  []Node
}

type Graph struct {
	AdjacencyList map[Node][]Node // we initialize adjacency list for easier neighbour lookup
	Edges         map[Node]map[Node]Path
}

func (g *Graph) addNode(elementNode Node) {
	_, ok := g.AdjacencyList[elementNode]
	if !ok {
		g.AdjacencyList[elementNode] = []Node{}
	}
}

func (g *Graph) addEdge(fromNode Node, toNode Node, weight float64, pathNodes []Node) {
	g.addNode(fromNode)
	g.addNode(toNode)
	g.AdjacencyList[fromNode] = append(g.AdjacencyList[fromNode], toNode)
	if g.Edges == nil {
		g.Edges = make(map[Node]map[Node]Path)
	}
	if g.Edges[fromNode] == nil {
		g.Edges[fromNode] = make(map[Node]Path)
	}
	g.Edges[fromNode][toNode] = Path{Weight: weight, Nodes: pathNodes}
}

func (g *Graph) getNeighbors(elementNode Node) []Node {
	neighbors, ok := g.AdjacencyList[elementNode]
	if !ok {
		return []Node{}
	}
	return neighbors
}

func (g *Graph) printGraphEdges() {
	for fromNode, toNodes := range g.Edges {
		for toNode, path := range toNodes {
			log.Print(fromNode, " -> ", toNode, " Cost: ", path.Weight, " Path: ", path.Nodes)
		}
	}
}
