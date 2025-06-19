package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CZnavody19/graph-algorithms/algorithm"
	"github.com/CZnavody19/graph-algorithms/database"
	"github.com/CZnavody19/graph-algorithms/utils"
)

func makeSampleGraph() ([][]uint16, uint16) {
	return [][]uint16{
		{0, 10, 0, 0, 3, 0},
		{0, 0, 2, 0, 0, 8},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 4, 0, 0, 0},
		{0, 0, 0, 0, 0, 3},
		{0, 0, 4, 1, 0, 0},
	}, 6
}

func getGraphFromDB(connectionString string) ([][]uint16, uint16) {
	dbConn, err := database.NewDatabaseConnection(connectionString)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		os.Exit(1)
	}

	graph, size, err := dbConn.GetGraphFromDB()
	if err != nil {
		fmt.Printf("Failed to get graph from the database: %v\n", err)
		os.Exit(1)
	}

	return graph, size
}

func validateVertices(v1, v2, numberOfVertices uint16) error {
	if v1 == v2 {
		return fmt.Errorf("vertices must be different")
	}
	if v1 > numberOfVertices || v2 > numberOfVertices {
		return fmt.Errorf("vertices must be in the range 1 to %d", numberOfVertices)
	}
	return nil
}

func interactive(distances, pathVertex *[][]uint16, numberOfVertices uint16) {
	for {
		var input string
		fmt.Print("Enter a pair of vertices like v1,v2 or q to exit: ")

		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Failed to read input:", err)
			continue
		}

		if input == "q" {
			return
		}

		var v1, v2 uint16
		_, err = fmt.Sscanf(input, "%d,%d", &v1, &v2)
		if err != nil {
			fmt.Println("Invalid input format. Please use 'v1,v2' format.")
			continue
		}
		err = validateVertices(v1, v2, numberOfVertices)
		if err != nil {
			fmt.Println(err)
			continue
		}

		distance := (*distances)[v1][v2]
		if distance == ^uint16(0) {
			fmt.Printf("There is no path between %d and %d.\n", v1, v2)
			continue
		}
		fmt.Printf("Distance between %d and %d: %d\n", v1, v2, distance)

		fmt.Print("Path: ")
		path := algorithm.GetPath(pathVertex, v1, v2)
		utils.PrintPath(path)

		fmt.Println()
	}
}

func main() {
	info := false

	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Usage: ./graph-algorithms <\"example\" | postgres connection string> <number of threads> <optional \"info\">")
		os.Exit(1)
	}

	numThreads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid number of threads: %s\n", os.Args[2])
		os.Exit(1)
	}

	if len(os.Args) == 4 && os.Args[3] == "info" {
		info = true
	}

	var graph [][]uint16
	var size uint16

	if os.Args[1] == "example" {
		graph, size = makeSampleGraph()
	} else {
		graph, size = getGraphFromDB(os.Args[1])
	}

	if info {
		utils.PrintMatrix(graph, "Input graph")
	}

	fmt.Println("Initializing matrices...")
	distances, pathVertex := algorithm.InitializeMatrices(&graph, size)

	if info {
		utils.PrintMatrix(*distances, "Initialized distances")
		utils.PrintMatrix(*pathVertex, "Initialized path vertex")
	}

	fmt.Println("Running Floyd-Warshall algorithm...")
	if numThreads > 1 {
		algorithm.FloydWarshallParallel(distances, pathVertex, size, numThreads)
	} else {
		algorithm.FloydWarshall(distances, pathVertex, size)
	}

	if info {
		utils.PrintMatrix(*distances, "Post algo distances")
		utils.PrintMatrix(*pathVertex, "Post algo path vertex")
	}

	interactive(distances, pathVertex, size)
}
