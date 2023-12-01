package handlers

import (
    "encoding/json"
	"fmt"
    "net/http"
    "github.com/gorilla/mux"
    "test/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Posts)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var post models.Post

	for _, p := range models.Posts {
		if p.ID == id {
			post = p
			break
		}
	}
	json.NewEncoder(w).Encode(post)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Users)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User

	for _, u := range models.Users {
		if u.ID == id {
			user = u
			break
		}
	}

	json.NewEncoder(w).Encode(user)
}