package routes

type AddBookResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	ID     int    `json:"id"`
}
