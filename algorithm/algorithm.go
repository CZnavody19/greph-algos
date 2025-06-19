package algorithm

import (
	"sync"
)

func innerLoop(i, k, vertexCount uint16, distances, pathVertex *[][]uint16) {
	for j := uint16(0); j < vertexCount; j++ {
		if (*distances)[i][k] != ^uint16(0) && (*distances)[k][j] != ^uint16(0) && (*distances)[i][j] > (*distances)[i][k]+(*distances)[k][j] {
			(*distances)[i][j] = (*distances)[i][k] + (*distances)[k][j]
			(*pathVertex)[i][j] = (*pathVertex)[k][j]
		}
	}
}

func FloydWarshall(distances, pathVertex *[][]uint16, vertexCount uint16) {
	for k := uint16(0); k < vertexCount; k++ {
		for i := uint16(0); i < vertexCount; i++ {
			innerLoop(i, k, vertexCount, distances, pathVertex)
		}
	}
}

func FloydWarshallParallel(distances, pathVertex *[][]uint16, vertexCount uint16, numThreads int) {
	for k := uint16(0); k < vertexCount; k++ {
		var wg sync.WaitGroup
		sem := make(chan struct{}, numThreads)

		for i := uint16(0); i < vertexCount; i++ {
			sem <- struct{}{}
			wg.Add(1)
			go func(i uint16) {
				defer wg.Done()
				innerLoop(i, k, vertexCount, distances, pathVertex)
				<-sem
			}(i)
		}
		wg.Wait()
	}
}

func InitializeMatrices(graph *[][]uint16, vertexCount uint16) (*[][]uint16, *[][]uint16) {
	var distances = make([][]uint16, vertexCount)
	var pathVertex = make([][]uint16, vertexCount)

	for i := uint16(0); i < vertexCount; i++ {
		distances[i] = make([]uint16, vertexCount)
		pathVertex[i] = make([]uint16, vertexCount)

		for j := uint16(0); j < vertexCount; j++ {
			if i == j {
				distances[i][j] = 0
				pathVertex[i][j] = i
			} else if (*graph)[i][j] != 0 {
				distances[i][j] = (*graph)[i][j]
				pathVertex[i][j] = i
			} else {
				distances[i][j] = ^uint16(0)
				pathVertex[i][j] = ^uint16(0)
			}
		}
	}

	return &distances, &pathVertex
}

func GetPath(pathVertex *[][]uint16, start, end uint16) []uint16 {
	if (*pathVertex)[start][end] == ^uint16(0) {
		return nil
	}

	path := []uint16{end}

	for start != end {
		end = (*pathVertex)[start][end]
		path = append([]uint16{end}, path...) //prepend
	}

	return path
}
