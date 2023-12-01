package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Active   bool   `json:"active"`
}

type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID int    `json:"author_id"`
}