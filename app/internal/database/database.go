package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// parameters for connecting
// to postgresql cluster.
const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "super-secret-password"
	dbname   = "debezium"
)

// generateUrl
// is used for creating the url for connecting
// to postgresql based on parameters that we give.
func generateUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
}

// NewConnection
// opens a connection to postgresql cluster.
func NewConnection() (*sql.DB, error) {
	return sql.Open("postgres", generateUrl())
}
