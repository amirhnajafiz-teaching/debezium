package cmd

import (
	"database/sql"
	"os"

	"github.com/spf13/cobra"
)

type Migrate struct {
	Connection *sql.DB
	Source     string
}

func (m Migrate) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { m.main() }
	return &cobra.Command{Use: "migrate", Run: run}
}

func (m Migrate) main() {
	// reading query
	data, err := os.ReadFile(m.Source)
	if err != nil {
		panic(err)
	}

	// execute migration query
	if _, err := m.Connection.Exec(string(data)); err != nil {
		panic(err)
	}
}
