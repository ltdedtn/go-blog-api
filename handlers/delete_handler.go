package handlers

import (
	"go-blog-api/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, post := range storage.Posts {
		if post.ID == id {
			storage.Posts = append(storage.Posts[:i], storage.Posts[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}