package cmd

import (
	"database/sql"
	"os"

	"github.com/amirhnajafiz/debezium/app/internal/database"

	"github.com/spf13/cobra"
)

type Migrate struct {
	Connection *sql.DB
}

func (m *Migrate) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { m.main() }
	return &cobra.Command{Use: "migrate", Run: run}
}

func (m *Migrate) main() {
	// defining the migration files
	source := "./internal/database/migrate/migrate.sql"

	// reading query
	data, err := os.ReadFile(source)
	if err != nil {
		panic(err)
	}

	// opening connection to postgresQL
	conn, err := database.NewConnection()
	if err != nil {
		panic(err)
	}

	// execute migration query
	if _, err := conn.Exec(string(data)); err != nil {
		panic(err)
	}
}
