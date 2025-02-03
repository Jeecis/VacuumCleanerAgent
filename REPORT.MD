# Analysis
### Problem description (optional to read)
In short, the task entails that an intelligent vacuum  navigates a room represented as an n*n grid. Each square contains a value indicating either empty space, dirt (1-9000), or an impassable wall (9001). The agent moves in four directions (north, south, east, or west), with movement costing E energy units and cleaning dirt costing V energy units. Given a starting position and an initial battery capacity B, the agent must maximize dirt cleaned and squares visited while managing its energy efficiently. The input is a CSV file representing the room, and the output includes the total dirt cleaned, squares cleared, path taken, and remaining battery.

### Prelude
While carefully, thinking about the task it is clear that an optimal solution would require a either a genius algorithm or very complex one. That's when I luckily encountered an Instagram reel advertisement about Traveling Salesman Problem (TSP), which seemed to be very similar to the provided one. However, after doing a little bit more research the problem described in the task is a version of TSP known as Orienteering TSP (OP). This is important, because TSP is NP-hard solution meaning that for a large dataset the complexity grows exponentially and to somewhat optimally solve it, one must use an approximation algorithm such as Christofide's algorithm.

### Initial ideas
However, it is important to mention some initial thoughts on solving this problem. The most obvious one I had was to calculate for each dirt square (node) some specific rating value and then move to the best one and then repeat the algorithm until vacuum runs out of energy. One such estimate would be using greedy search with `dirt_amount/(shortest_path_cost + vacuum_cost)` formula, however, this algorithm fails to recognize a more dense spots dirt squares that would be more beneficial to move to at start.

### Preparation

To ensure I had prepared the most optimal solution, I chose to at first model the data for the TSP problem. This means that I need to create a graph where every dirt node is connected with each other and with store the most efficient path of it to trace it back in the end. Space complexity wise this is an inefficient part of my code, however, it is to increase the time complexity later. 

For this I developed a class graph which consists of an adjacency list and edges that store start node, end node and path of the 

### Results for different test grids
### Challenges