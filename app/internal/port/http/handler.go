package http

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	insertQuery = "INSERT INTO `users` (`name`, `email`) VALUES (?, ?)"
	selectQuery = "SELECT * FROM `users`"
)

type Handler struct {
	Connection *sql.DB
}

func (h *Handler) HandlePostRequests(ctx *fiber.Ctx) error {
	var userReq request

	if err := ctx.BodyParser(&userReq); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if _, err := h.Connection.Query(insertQuery, userReq.Name, userReq.Email); err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendString("user added")
}

func (h *Handler) HandleGetRequests(ctx *fiber.Ctx) error {
	var (
		users []response
		user  response
	)

	rows, err := h.Connection.Query(selectQuery)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	for rows.Next() {
		if er := rows.Scan(&user); er != nil {
			return ctx.SendStatus(http.StatusLoopDetected)
		}

		users = append(users, user)
	}

	return ctx.JSON(users)
}
