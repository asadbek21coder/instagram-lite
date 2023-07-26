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

	posts = append(posts, newpost)
	data, _ := json.Marshal(posts)

	os.WriteFile("db/posts.json", data, 0)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPosts(w http.ResponseWriter, r *http.Request) {

	var res []models.GetAllPosts
	var posts []models.Post
	readPost, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(readPost, &posts)
	for i := 0; i < len(posts); i++ {
		var newPost models.GetAllPosts
		newPost.ID = posts[i].ID
		newPost.Body = posts[i].Body
		newPost.Title = posts[i].Title
		newPost.ImageUrl = posts[i].ImageUrl
		newPost.LikeCount = posts[i].LikeCount
		newPost.CreatedAt = posts[i].CreatedAt

		var users []models.User
		readUser, _ := os.ReadFile("db/users.json")
		json.Unmarshal(readUser, &users)

		for i2 := 0; i2 < len(users); i2++ {
			if users[i2].ID == posts[i].AuthorId {
				newPost.Author = users[i2]
			}
		}

		var comments []models.Comment
		readComment, _ := os.ReadFile("db/comments.json")
		json.Unmarshal(readComment, &comments)

		for c := 0; c < len(comments); c++ {
			if comments[c].CommentedPostId == posts[i].ID {
				var newComment models.CommentWithreply

				newComment.Comment.Body = comments[c].Body
				newComment.Comment.LikeCount = comments[c].LikeCount
				newComment.Comment.AuthorId = comments[c].AuthorId
				newComment.Comment.CreatedAt = comments[c].CreatedAt
				newComment.Comment.ID = comments[c].ID

				var replies []models.Reply
				// var newReply models.Reply
				readReply, _ := os.ReadFile("db/replys.json")
				json.Unmarshal(readReply, &replies)

				for r := 0; r < len(replies); r++ {
					if replies[r].CommentId == comments[c].ID {
						newComment.Replies = append(newComment.Replies, replies[r])
					}
				}
				newPost.Comments = append(newPost.Comments, newComment)
			}

		}

		res = append(res, newPost)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func updatePost(w http.ResponseWriter, r *http.Request) {

}

func deletePost(w http.ResponseWriter, r *http.Request) {

}
