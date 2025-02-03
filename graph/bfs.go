package graph

import (
	"container/list"
	"math"
)

func (g *Graph) bfs(start, goal Node) ([]Node, float64) {
	queue := list.New()
	queue.PushBack(start)

	cameFrom := make(map[Node]Node)
	cameFrom[start] = start

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(Node)

		if current == goal {
			path := []Node{}
			for current != start {
				path = append([]Node{current}, path...)
				current = cameFrom[current]
			}
			path = append([]Node{start}, path...)
			return path, float64(len(path) - 1)
		}

		for _, neighbor := range g.getNeighbors(current) {
			if _, ok := cameFrom[neighbor]; !ok {
				queue.PushBack(neighbor)
				cameFrom[neighbor] = current
			}
		}
	}

	return nil, math.Inf(1)
}
