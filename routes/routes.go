package routes

import (
	"go-blog-api/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
    r.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
    r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
    return r
}
