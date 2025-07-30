package handlers

import (
	"encoding/json"
	"go-blog-api/models"
	"go-blog-api/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(storage.Posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for _, post := range storage.Posts {
        if post.ID == id {
            json.NewEncoder(w).Encode(post)
            return
        }
    }
    http.Error(w, "Post not found", http.StatusNotFound)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var newPost models.Post
    json.NewDecoder(r.Body).Decode(&newPost)
    newPost.ID = len(storage.Posts) + 1
    storage.Posts = append(storage.Posts, newPost)
    json.NewEncoder(w).Encode(newPost)
}
