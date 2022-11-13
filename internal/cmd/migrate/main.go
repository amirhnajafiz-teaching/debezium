package migrate

import (
	"os"

	"github.com/amirhnajafiz/sql-kafka-debezium/internal/database"

	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {
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
