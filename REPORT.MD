# Analysis

## Introduction
### Problem description (optional to read)
In short, the task entails that an intelligent vacuum  navigates a room represented as an n*n grid. Each square contains a value indicating either empty space, dirt (1-9000), or an impassable wall (9001). The agent moves in four directions (north, south, east, or west), with movement costing E energy units and cleaning dirt costing V energy units. Given a starting position and an initial battery capacity B, the agent must maximize dirt cleaned and squares visited while managing its energy efficiently. The input is a CSV file representing the room, and the output includes the total dirt cleaned, squares cleared, path taken, and remaining battery.

### Prelude to the approach
While carefully, thinking about the task it is clear that an optimal solution would require a either a genius algorithm or very complex one. That's when I luckily encountered an Instagram reel advertisement about Traveling Salesman Problem (TSP), which seemed to be very similar to the provided one. However, after doing a little bit more research the problem described in the task is a version of TSP known as Orienteering TSP (OP). This is important, because TSP is NP-hard solution meaning that for a large dataset the complexity grows exponentially and to somewhat optimally solve it, one must use an approximation algorithm such as Christofide's algorithm.

### Thinking process
However, it is important to mention some initial thoughts on solving this problem. The most obvious one I had was to calculate for each dirt square (node) some specific rating value and then move to the best one and then repeat the algorithm until vacuum runs out of energy. One such estimate would be using greedy search with `dirt_amount/(shortest_path_cost + vacuum_cost)` formula, however, this algorithm fails to recognize a more dense spots dirt squares that would be more beneficial to move to at start. 
So I went exploring ways to solve OP. Turns out it can be solved using integer linear programming, if properly defined. However, since I didn't understand the mathematics behind it after hours of research I decided to drop the idea and stick with the initial idea of development.

## How the algorithm works
### Preparation
To ensure I had prepared the most optimal solution, I chose to at first model the data for the TSP problem. This means that I need to create a graph where every dirt node is connected with each other and to store the most efficient path of it to trace it back in the end. Space complexity wise this is an inefficient part of my code, however, it is to increase the time complexity later. 

For this I developed a class graph which consists of an adjacency list and edges map that stores start node, end node and path of between the nodes. 

So at the start I read in each square as a node of a graph and then created a distilled version of the graph using bfs to store only dirt nodes and paths.

Thus implementing a solver `dirt_amount/(shortest_path_cost + vacuum_cost)` was an easier part since all it took was to iterate through neighbours, find the best candidate that has the highest dirt_amount/total_cost coefficient and then move to it and repeat the algorithm.

### Results for different test grids
input.csv result:
`Total dirt: 9190, Squares cleared: 6, Remaining battery: 3, Path: {47 [{false 0 0 0} {false 1 0 0} {false 2 0 0} {false 2 1 20} {false 2 2 50} {false 2 3 40} {false 2 4 50} {false 3 4 0} {false 4 4 9000} {false 3 4 0} {false 2 4 50} {false 2 3 40} {false 2 2 50} {false 2 2 50} {false 1 3 30} {false 2 3 40} {false 2 2 50} {false 2 1 20}]}`
It is important to note that the Vacuum cleaner attempted to clean the last square containing 10 dirt thats why it continued to move. 

input2.cs result:
`Total dirt: 27, Squares cleared: 2, Remaining battery: 1, Path: {9 [{false 3 3 0} {false 4 3 1} {false 5 3 1} {false 5 4 20} {false 4 4 1} {false 3 4 1} {false 2 4 5} {false 2 3 7}]}`
Here the algorithm failed to recognize that it would have made more sense to move to an area where there are more nodes with average higher values than that distant node of value 20.

input3.csv
` Total dirt: 500, Squares cleared: 10, Remaining battery: 5, Path: {195 [{false 5 5 0} {false 4 5 50} {false 3 5 0} {false 3 4 50} {false 3 3 0} {false 3 2 50} {false 3 3 0} {false 3 2 50} {false 3 3 0} {false 2 3 50} {false 2 4 0} {false 2 3 50} {false 1 5 0} {false 1 4 50} {false 1 3 0} {false 1 2 50} {false 0 2 0} {false 0 1 50} {false 0 2 0} {false 0 1 50}]}`
Here the vacuum chose the most optimal path by taking the shortests paths to the nodes.

### Challenges
A big challenge was to prepare the graph since I didn't know exactly what kind of data I will be storing. thus I chose to store almost everything including a path from one node to another.

Furthermore, the biggest challenge (which also accounts as a failure for me) was the attempt to use linear programming algorithms to solve the problem. I tried to use google OR tools, which had a sketchy package for golang and it didn't work. Thus I just didn't have enough brain power to optimize the existing algorithm. But in the end at least I feel a bit more knowledgeable in TSP problem solving.

And last challenge was fixing all the graph initialization bugs that occurred, which I successfully finished.

## Conclusion
I would rate my algorithm as an alright for solving the problem. It still has several issues as not choosing the best node at the start and not recognizing dense areas and overvaluing nodes with high dirt amount.