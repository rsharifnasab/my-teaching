package routes

type AddBookRequest struct {
	Author string `json:"author" validte:"required"`
	Title  string `json:"title" validate:"required"`
}
