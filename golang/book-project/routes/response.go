package routes

type AddBookResponse struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	ID     int    `json:"id"`
}
