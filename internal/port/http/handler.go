package http

import "database/sql"

type Handler struct {
	Connection *sql.DB
}

func (h *Handler) HandleRequests() {

}
