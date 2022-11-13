package http

import (
	"github.com/amirhnajafiz/henry/internal/database"
	"github.com/amirhnajafiz/henry/internal/port/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use: "http",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {
	// opening connection to postgresQL
	conn, err := database.NewConnection()
	if err != nil {
		panic(err)
	}

	// creating a new handler
	h := http.Handler{
		Connection: conn,
	}

	// creating a new fiber app
	app := fiber.New()

	// defining routes
	app.Get("/api", h.HandleGetRequests)
	app.Post("/api", h.HandlePostRequests)

	// running app
	if err := app.Listen(":7490"); err != nil {
		panic(err)
	}
}
