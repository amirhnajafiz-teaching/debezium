package http

type response struct {
	id    int    `json:"id"`
	name  string `json:"name"`
	email string `json:"email"`
}
