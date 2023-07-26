package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/asadbek21coder/instagram/models"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createComment(w, r)
	case http.MethodGet:
		getComments(w, r)
	case http.MethodPut:
		updateComment(w, r)
	case http.MethodDelete:
		deleteComment(w, r)
	}

}

func createComment(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comment
	json.NewDecoder(r.Body).Decode(&newComment)

	var comments []models.Comment

	read, _ := os.ReadFile("db/comments.json")
	json.Unmarshal(read, &comments)

	if len(comments) == 0 {
		newComment.ID = 1
	} else {
		newComment.ID = len(comments) + 1
	}
	newComment.CreatedAt = time.Now()

	comments = append(comments, newComment)

	writeData, _ := json.Marshal(comments)
	os.WriteFile("db/comments.json", writeData, 0)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newComment)

}

func getComments(w http.ResponseWriter, r *http.Request) {
	var comments []models.Comment
	readComments, _ := os.ReadFile("db/comments.json")
	json.Unmarshal(readComments, &comments)
	var res []models.CommentWithreply

	for i := 0; i < len(comments); i++ {
		var newComment models.CommentWithreply
		// newComment.Comment.ID = comments[i].ID
		// newComment.Comment.Body = comments[i].Body
		// newComment.Comment.CommentedPostId = comments[i].CommentedPostId
		// newComment.Comment.AuthorId = comments[i].AuthorId
		// newComment.Comment.CreatedAt = comments[i].CreatedAt
		// newComment.Comment.LikeCount = comments[i].LikeCount
		newComment.Comment = comments[i]

		var replies []models.Reply
		readReply, _ := os.ReadFile("db/replys.json")
		json.Unmarshal(readReply, &replies)

		for r := 0; r < len(replies); r++ {
			if replies[r].CommentId == comments[i].ID {
				newComment.Replies = append(newComment.Replies, replies[r])
			}
		}
		res = append(res, newComment)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func updateComment(w http.ResponseWriter, r *http.Request) {

}

func deleteComment(w http.ResponseWriter, r *http.Request) {

}
