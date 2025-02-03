package solver

import "github.com/Jeecis/VacuumCleanerAgent/graph"

type candidate struct {
	target graph.Node
	// cost in energy units needed to move (not including vacuum cost)
	movementCost int
	// total cost including vacuuming at the target.
	totalCost int
	score     float64
	// path from the current node to the candidate.
	path []graph.Node
}

func FindOrienteeringPath(g graph.Graph, start graph.Node, battery int, moveCost, vacuumCost int) (totalDirt int, squaresCleared int, resultPath graph.Path, remainingBattery int) {
	// Set up state
	visited := make(map[graph.Node]bool)
	current := start
	remainingBattery = battery
	resultPath.Nodes = []graph.Node{start}

	// If the starting square has dirt, “vacuum” it (if enough battery remains)
	if !start.IsWall && start.DirtCount > 0 && remainingBattery >= vacuumCost {
		totalDirt += start.DirtCount
		remainingBattery -= vacuumCost
		visited[start] = true
		squaresCleared++
	}

	// Continue until no reachable candidate remains
	for {
		var bestCand *candidate
		// Look over all nodes in the graph’s adjacency list.
		// (In a full solution you might iterate over a list of dirt targets.)
		for target := range g.AdjacencyList {
			// Skip walls and already visited squares.
			if target.IsWall || visited[target] {
				continue
			}
			// Skip targets with no dirt – if you want to maximize squares cleared,
			// you might also consider empty nodes. Here we prioritize dirt.
			if target.DirtCount <= 0 {
				continue
			}
			// Compute the shortest path from current to target.
			path := g.Edges[current][target]

			// Energy needed to move along the path.
			costForMovement := int(path.Weight) * moveCost
			// Total cost includes vacuuming at the target.
			totalCost := costForMovement + vacuumCost
			// If we do not have enough battery, skip this candidate.
			if totalCost > remainingBattery {
				continue
			}
			// Compute a score: here we use dirt per energy unit.
			score := float64(target.DirtCount) / float64(totalCost)
			// Update best candidate if this one is better.
			if bestCand == nil || score > bestCand.score {
				bestCand = &candidate{
					target:       target,
					movementCost: costForMovement,
					totalCost:    totalCost,
					score:        score,
					path:         path.Nodes,
				}
			}
		}
		// If no candidate found, exit the loop.
		if bestCand == nil {
			break
		}

		// “Travel” along the computed path:
		// Note: In this example, we assume that moving along the intermediate nodes does not vacuum.
		// Only at the candidate target do we vacuum.
		// Also, we subtract the movement energy cost.
		remainingBattery -= bestCand.movementCost
		// Append intermediate nodes (skipping the first which is current)
		for i := 1; i < len(bestCand.path); i++ {
			resultPath.Nodes = append(resultPath.Nodes, bestCand.path[i])
		}

		// Vacuum at the target (if possible)
		if remainingBattery >= vacuumCost {
			remainingBattery -= vacuumCost
			totalDirt += bestCand.target.DirtCount
			squaresCleared++
		} else {
			// Not enough battery to vacuum (should not happen due to candidate filtering)
			break
		}

		// Mark the target as visited and update the current position.
		visited[bestCand.target] = true
		current = bestCand.target
	}

	// Compute the total weight as the battery used (or you can define it otherwise)
	usedBattery := battery - remainingBattery
	resultPath.Weight = float64(usedBattery)

	return totalDirt, squaresCleared, resultPath, remainingBattery
}
