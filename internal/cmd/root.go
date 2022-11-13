package cmd

import (
	"github.com/amirhnajafiz/henry/internal/cmd/http"
	"github.com/amirhnajafiz/henry/internal/cmd/kafka"
	"github.com/amirhnajafiz/henry/internal/cmd/migrate"

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
