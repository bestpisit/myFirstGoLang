package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"test/models"
)

var users = []User{
	{ID: "1", Name: "John Doe", Username: "johndoe", Active: true},
	{ID: "2", Name: "Jane Doe", Username: "janedoe", Active: true},
	{ID: "3", Name: "Michael Jordan", Username: "michaeljordan", Active: true},
	{ID: "4", Name: "John Smith", Username: "johnsmith", Active: true},
}

var posts = []Post{
	{ID: "1", Title: "Hello World", Body: "This is my first post", AuthorID: 1},
	{ID: "2", Title: "Hello World 2", Body: "This is my second post", AuthorID: 1},
	{ID: "3", Title: "Hello World 3", Body: "This is my third post", AuthorID: 2},
	{ID: "4", Title: "Hello World 4", Body: "This is my fourth post", AuthorID: 3},
}

func main() {

	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/posts", PostsHandler)
	r.HandleFunc("/posts/{id}", PostHandler)
	r.HandleFunc("/users", UsersHandler)
	r.HandleFunc("/users/{id}", UserHandler)

	log.Fatal(http.ListenAndServe(":8910", r))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}