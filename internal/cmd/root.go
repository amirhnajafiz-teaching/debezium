package cmd

import (
	"github.com/amirhnajafiz/sql-kafka-debezium/internal/cmd/http"
	"github.com/amirhnajafiz/sql-kafka-debezium/internal/cmd/migrate"

	"github.com/spf13/cobra"
)

func Execute() {
	cmd := cobra.Command{}

	cmd.AddCommand(
		http.GetCommand(),
		migrate.GetCommand(),
	)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
