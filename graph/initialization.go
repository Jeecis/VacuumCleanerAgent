package graph

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

type vacuumInfo struct {
	X             int
	Y             int
	InitialDirt   int
	Battery       int
	MovementCost  int
	VacuumingCost int
}

func TaskInit(csvPath string) (*vacuumInfo, *Graph, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Could not read CSV file: %v", err)
	}

	vacuum, err := vacuumInit(records)
	if err != nil {
		return nil, nil, err
	}

	graph, keyNodes, startNode, err := graphInit(records, vacuum.X, vacuum.Y)
	if err != nil {
		return nil, nil, err
	}

	tspGraph := tspGraphInit(graph, keyNodes, startNode)
	log.Print(vacuum)

	return vacuum, tspGraph, nil

}

func vacuumInit(records [][]string) (*vacuumInfo, error) {
	startX, err := strconv.Atoi(records[0][0])
	if err != nil {
		return nil, err
	}

	startY, err := strconv.Atoi(records[1][0])
	if err != nil {
		return nil, err
	}

	startBattery, err := strconv.Atoi(records[2][0])
	if err != nil {
		return nil, err
	}

	movementCost, err := strconv.Atoi(records[3][0])
	if err != nil {
		return nil, err
	}

	vacumingCost, err := strconv.Atoi(records[4][0])
	if err != nil {
		return nil, err
	}

	initialDirt, err := strconv.Atoi(records[startY+5][startX])
	if err != nil {
		return nil, err
	}

	vacuum := vacuumInfo{
		X:             startX,
		Y:             startY,
		InitialDirt:   initialDirt,
		Battery:       startBattery,
		MovementCost:  movementCost,
		VacuumingCost: vacumingCost,
	}

	return &vacuum, nil
}

func graphInit(records [][]string, startX int, startY int) (*Graph, []Node, *Node, error) {
	distilledRecords := records[5:]
	var startNode Node
	nodeArray := make([][]Node, len(distilledRecords))
	for i := range nodeArray {
		nodeArray[i] = make([]Node, len(distilledRecords)) // Adjust size as needed
	}

	var dirtNodes []Node
	for i, row := range distilledRecords {
		for j, cell := range row {

			cell = strings.ReplaceAll(cell, " ", "")
			value, err := strconv.Atoi(cell)
			if err != nil {
				return nil, nil, nil, err
			}

			isWall := false
			if value == 9001 {
				isWall = true
			}

			n := Node{
				IsWall:      isWall,
				XCoordinate: j,
				YCoordinate: i,
				DirtCount:   value,
			}

			if j == startX && i == startY {
				startNode = n
				nodeArray[i][j] = n // avoid adding start node into dirt Nodes
				continue
			}

			if value > 0 && !isWall {
				dirtNodes = append(dirtNodes, n)
			}

			nodeArray[i][j] = n
		}
	}

	g := connectNodes(nodeArray)

	return g, dirtNodes, &startNode, nil

}

func connectNodes(nodeGrid [][]Node) *Graph {
	rows := len(nodeGrid)

	// Initialize the graph.
	g := &Graph{
		AdjacencyList: make(map[Node][]Node),
		Edges:         make(map[Node]map[Node]Path),
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			currentNode := nodeGrid[i][j]
			if currentNode.IsWall {
				continue
			}

			// Check the neighbor below.
			if i+1 < rows {
				neighbor := nodeGrid[i+1][j]
				if !neighbor.IsWall {
					g.addEdge(currentNode, neighbor, 1, []Node{currentNode, neighbor})
					g.addEdge(neighbor, currentNode, 1, []Node{neighbor, currentNode})
				}
			}

			// Check the neighbor to the right.
			if j+1 < rows {
				neighbor := nodeGrid[i][j+1]
				if !neighbor.IsWall {
					g.addEdge(currentNode, neighbor, 1, []Node{currentNode, neighbor})
					g.addEdge(neighbor, currentNode, 1, []Node{neighbor, currentNode})
				}
			}
		}
	}

	return g

}
