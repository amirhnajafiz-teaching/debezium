package main

import (
	"github.com/amirhnajafiz/debezium/app/cmd"
	"github.com/amirhnajafiz/debezium/app/internal/database"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{}

	conn, err := database.NewConnection()
	if err != nil {
		panic(err)
	}

	root.AddCommand(
		cmd.Migrate{Connection: conn, Source: "./internal/database/migrate/migrate.sql"}.Command(),
		cmd.HTTP{Connection: conn, Port: 7041}.Command(),
	)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
