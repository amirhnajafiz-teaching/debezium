package main

import (
	"github.com/amirhnajafiz/debezium/app/cmd/http"
	"github.com/amirhnajafiz/debezium/app/cmd/migrate"

	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{}

	cmd.AddCommand(
		http.GetCommand(),
		migrate.GetCommand(),
	)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
