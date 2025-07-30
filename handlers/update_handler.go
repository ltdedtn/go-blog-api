package handlers

import (
	"encoding/json"
	"go-blog-api/models"
	"go-blog-api/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, post := range storage.Posts {
		if post.ID == id {
			var updatedPost models.Update
			json.NewDecoder(r.Body).Decode(&updatedPost)
			storage.Posts[i].Title = updatedPost.Title
			storage.Posts[i].Content = updatedPost.Content
			json.NewEncoder(w).Encode(storage.Posts[i])
			return
		}
	}
}