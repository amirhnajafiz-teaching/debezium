package cmd

import (
	"github.com/amirhnajafiz/debezium/app/internal/cmd/http"
	"github.com/amirhnajafiz/debezium/app/internal/cmd/kafka"
	"github.com/amirhnajafiz/debezium/app/internal/cmd/migrate"

	"github.com/spf13/cobra"
)

func Execute() {
	cmd := cobra.Command{}

	cmd.AddCommand(
		http.GetCommand(),
		migrate.GetCommand(),
		kafka.GetCommand(),
	)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
