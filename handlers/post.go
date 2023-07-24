package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/asadbek21coder/instagram/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createPost(w, r)
	case http.MethodGet:
		getPosts(w, r)
	case http.MethodPut:
		updatePost(w, r)
	case http.MethodDelete:
		deletePost(w, r)
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	var newpost models.Post
	json.NewDecoder(r.Body).Decode(&newpost)
	readposts, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(readposts, &posts)

	newpost.ID = len(posts) + 1 // logikasini tog'irlash kk
	newpost.CreatedAt = time.Now()
	// qo'shishdan oldin postname band emasligini tekshirish, band bo'lsa error qaytarish kk
	posts = append(posts, newpost)
	data, _ := json.Marshal(posts)

	os.WriteFile("db/posts.json", data, 0)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPosts(w http.ResponseWriter, r *http.Request) {

}

func updatePost(w http.ResponseWriter, r *http.Request) {

}

func deletePost(w http.ResponseWriter, r *http.Request) {

}
