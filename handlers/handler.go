package handlers

import (
    "encoding/json"
	"fmt"
    "net/http"
    "github.com/gorilla/mux"
    "test/models"
	"test/cmd"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var post models.Post

	for _, p := range posts {
		if p.ID == id {
			post = p
			break
		}
	}
	json.NewEncoder(w).Encode(post)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user User

	for _, u := range users {
		if u.ID == id {
			user = u
			break
		}
	}

	json.NewEncoder(w).Encode(user)
}