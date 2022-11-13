package http

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	insertQuery = "INSERT INTO `users` (`name`, `email`) VALUES (?, ?)"
)

type Handler struct {
	Connection *sql.DB
}

func (h *Handler) HandleRequests(ctx *fiber.Ctx) error {
	var userReq request

	if err := ctx.BodyParser(&userReq); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if _, err := h.Connection.Query(insertQuery, userReq.name, userReq.email); err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendString("user added")
}
