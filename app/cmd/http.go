package cmd

import (
	"database/sql"
	"fmt"

	"github.com/amirhnajafiz/debezium/app/internal/port/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type HTTP struct {
	Connection *sql.DB
	Port       int
}

func (h HTTP) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { h.main() }
	return &cobra.Command{Use: "http", Run: run}
}

func (h HTTP) main() {
	// creating a new handler
	hd := http.Handler{
		Connection: h.Connection,
	}

	// creating a new fiber app
	app := fiber.New()

	// defining routes
	app.Get("/api", hd.HandleGetRequests)
	app.Post("/api", hd.HandlePostRequests)

	// running app
	if err := app.Listen(fmt.Sprintf(":%d", h.Port)); err != nil {
		panic(err)
	}
}
