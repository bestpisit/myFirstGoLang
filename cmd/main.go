package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"test/handlers"
)

func main() {

	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/posts", handlers.PostsHandler)
	r.HandleFunc("/posts/{id}", handlers.PostHandler)
	r.HandleFunc("/users", handlers.UsersHandler)
	r.HandleFunc("/users/{id}", handlers.UserHandler)

	log.Fatal(http.ListenAndServe(":8910", r))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}