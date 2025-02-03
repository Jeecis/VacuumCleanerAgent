package main

import (
	"log"

	"github.com/Jeecis/VacuumCleanerAgent/graph"
	"github.com/Jeecis/VacuumCleanerAgent/solver"
)

func main() {
	vacuumInfo, destilledGraph, err := graph.TaskInit("input.csv")
	if err != nil {
		log.Fatalf("Error initializing task: %v", err)
	}
	totalDirt, squaresCleared, resultPath, remainingBattery := solver.FindOrienteeringPath(*destilledGraph,
		graph.Node{XCoordinate: vacuumInfo.X, YCoordinate: vacuumInfo.Y, IsWall: false, DirtCount: vacuumInfo.InitialDirt},
		vacuumInfo.Battery,
		vacuumInfo.MovementCost,
		vacuumInfo.VacuumingCost)

	log.Printf("Total dirt: %d, Squares cleared: %d, Remaining battery: %d, Path: %v", totalDirt, squaresCleared, remainingBattery, resultPath)
}
