package database

//go:generate go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgres://alg:q5b6qT2NvzyCas@10.0.1.99:5432/alg -path=./gen

import (
	"database/sql"

	"github.com/CZnavody19/graph-algorithms/database/gen/alg/public/model"
	"github.com/CZnavody19/graph-algorithms/database/gen/alg/public/table"
	"github.com/XSAM/otelsql"
	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	DB *sql.DB
}

func NewDatabaseConnection(connectionString string) (*DatabaseConnection, error) {
	dbConn, err := otelsql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{
		DB: dbConn,
	}, nil
}

func (dc *DatabaseConnection) GetGraphFromDB() ([][]uint16, uint16, error) {
	vertifcesStmt := table.Vertices.SELECT(table.Vertices.AllColumns)

	var vertices []model.Vertices
	err := vertifcesStmt.Query(dc.DB, &vertices)
	if err != nil {
		return nil, 0, err
	}

	edgesStmt := table.Edges.SELECT(table.Edges.AllColumns)

	var edges []model.Edges
	err = edgesStmt.Query(dc.DB, &edges)
	if err != nil {
		return nil, 0, err
	}

	matrix, size := constructMatrix(vertices, edges)

	return matrix, size, nil
}

func constructMatrix(vertices []model.Vertices, edges []model.Edges) ([][]uint16, uint16) {
	vertexCount := uint16(len(vertices))
	matrix := make([][]uint16, vertexCount)

	for i := range matrix {
		matrix[i] = make([]uint16, vertexCount)
		for j := range matrix[i] {
			matrix[i][j] = ^uint16(0)
		}
	}

	for _, edge := range edges {
		matrix[*edge.Origin-1][*edge.Target-1] = uint16(*edge.Cost)
	}

	return matrix, vertexCount
}
