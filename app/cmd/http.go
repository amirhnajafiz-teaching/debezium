package cmd

import (
	"database/sql"

	"github.com/amirhnajafiz/debezium/app/internal/database"
	"github.com/amirhnajafiz/debezium/app/internal/port/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type HTTP struct {
	Connection *sql.Conn
}

func (h *HTTP) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { h.main() }
	return &cobra.Command{Use: "http", Run: run}
}

func (h *HTTP) main() {
	// opening connection to postgresQL
	conn, err := database.NewConnection()
	if err != nil {
		panic(err)
	}

	// creating a new handler
	hd := http.Handler{
		Connection: conn,
	}

	// creating a new fiber app
	app := fiber.New()

	// defining routes
	app.Get("/api", hd.HandleGetRequests)
	app.Post("/api", hd.HandlePostRequests)

	// running app
	if err := app.Listen(":7490"); err != nil {
		panic(err)
	}
}
