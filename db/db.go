package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DbCompany struct {
	uuid        string
	companyName string
	rating      float64
	lat         float64
	lon         float64
	companyType string
	createdAt   time.Time
	updatedAt   time.Time
}

func createCompany(conn *pgx.Conn, c DbCompany) {
	conn.Exec(
		context.Background(),
		"INSERT INTO company VALUES ($1, $2, $3, $4, $5, $6, $7, $8);",
		c.uuid,
		c.companyName,
		c.rating,
		c.lat,
		c.lon,
		c.companyType,
		c.createdAt,
		c.updatedAt,
	)
}

func clearCompanyTable(conn *pgx.Conn) pgconn.CommandTag {
	resp, err := conn.Exec(
		context.Background(),
		"DELETE FROM company;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}

	return resp
}
