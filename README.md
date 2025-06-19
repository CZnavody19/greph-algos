# Graph algorithms

## Graph theory
A graph is a set of nodes (vertices) connected by edges (links)

### Types of graphs
- Undirected - Points are interconnected by simple paths
- Directed - Points are interconnected by one way paths
- Weighted - Points are interconnected by paths each with its own numerical weight

### Representations
- Adjacency matrix - Edges and their weight are stored in a matrix which is indexed by two vertices
- Incidence matrix - A boolean matrix whose rows represent vertices and columns represent edges
- Adjacency list - An unordered list of vertex neighbours
- Edge list - An array of pairs of vertices

### Applications
- Navigation systems
- Computer and social networks
- Representation of chemical molecules

### The shortest path
When trying to find the shortest path we are looking for a path with the least combined cost in case of weighted graphs, or the path with the least hops between verticies in case of unweighted graphs

### Negative edges
The edge weight can be any number, but when its negative we can run into an issue where some algorithms keep taking the negative edge forever to keep the total cost as low as possible

### Algorithms
#### Dijkstra’s Algorithm
Uses a greedy approach and a priority queue to always choose the closest unvisited node

Doesnt work correctly with negative edges

Time complexity: O((V + E) log V)

#### Bellman-Ford Algorithm
Repeatedly relaxes edges (tries to improve shortest paths) up to V−1 times

Works even with negative edges

Time complexity: O(V × E)

#### Floyd-Warshall Algorithm
Uses dynamic programming to find shortest paths between all pairs of nodes

Best when you need all-pairs shortest paths

Time complexity: O(V³)

## Floyd-Warshall implementation in Go
### Building:
- Have Go 1.24
- `go build`

### Usage: 
`./graph-algorithms <"example" | postgres connection string> <number of threads> <optional "info">`

There are two ways to get a graph:
- The example graph
- Load the graph from [Michael's Postgres database](https://github.com/lopataa/delta-alg-graphs)

There is a single and a multithreaded version of the algorithm, when you set the number of threads to 1 it automatically runs the single threaded one

You can also print the matrices at different stages, just pass the info parameter - WARNING: really slow on large graphs

NOTE: The search is 0 based