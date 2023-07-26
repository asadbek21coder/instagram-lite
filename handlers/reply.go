package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/asadbek21coder/instagram/models"
)

func ReplyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createReply(w, r)
	case http.MethodGet:
		getReplies(w, r)
	case http.MethodPut:
		updateReply(w, r)
	case http.MethodDelete:
		deleteReply(w, r)
	}

}

func createReply(w http.ResponseWriter, r *http.Request) {
	var replies []models.Reply
	var newReply models.Reply
	json.NewDecoder(r.Body).Decode(&newReply)

	readReply, _ := os.ReadFile("db/replys.json")
	json.Unmarshal(readReply, &replies)

	if len(replies) == 0 {
		newReply.ID = 1
	} else {
		newReply.ID = len(replies) + 1
	}

	newReply.CreatedAt = time.Now()

	replies = append(replies, newReply)
	writeData, _ := json.Marshal(replies)

	os.WriteFile("db/replys.json", writeData, 0)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newReply)

}

func getReplies(w http.ResponseWriter, r *http.Request) {
	readReply, _ := os.ReadFile("db/replys.json")
	var replies []models.Reply
	var res []models.ReplyWithAuthor
	json.Unmarshal(readReply, &replies)

	for i := 0; i < len(replies); i++ {
		var newReply models.ReplyWithAuthor
		newReply.Reply = replies[i]

		var users []models.User
		readUser, _ := os.ReadFile("db/users.json")
		json.Unmarshal(readUser, &users)

		for u := 0; u < len(users); u++ {
			if replies[i].AuthorId == users[u].ID {
				newReply.Author = users[u]
			}
		}

		res = append(res, newReply)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func updateReply(w http.ResponseWriter, r *http.Request) {

}

func deleteReply(w http.ResponseWriter, r *http.Request) {

}
