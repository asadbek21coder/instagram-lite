package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

	query := r.URL.Query().Get("id")
	// fmt.Println(id == "")
	id, _ := strconv.Atoi(query)

	var res []models.GetAllPosts
	var posts []models.Post
	readPost, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(readPost, &posts)

	if query != "" {

		for i := 0; i < len(posts); i++ {

			if posts[i].ID == id {
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
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(newPost)
				return

			}

		}
	}

	for i := 0; i < len(posts); i++ {
		var newPost models.GetAllPosts
		newPost.ID = posts[i].ID
		newPost.Body = posts[i].Body
		newPost.Title = posts[i].Title
		newPost.ImageUrl = posts[i].ImageUrl
		newPost.LikeCount = posts[i].LikeCount
		newPost.CreatedAt = posts[i].CreatedAt
		newPost.UpdatedAt = posts[i].UpdatedAt

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

	// 1. requestda yangi update body
	// fmt.Println(r.Body)
	// 2. uni parse qilamiz(go tushunadigan formatga o'gramiz)

	var req models.Post
	json.NewDecoder(r.Body).Decode(&req)

	// 3. posts.json filedan ma`lumotlarni o'qib parse qilamiz(arrayga)
	var posts []models.Post
	readPosts, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(readPosts, &posts)

	// 4. postlar joylangan shu array bo'yicha yurib chiqib, bodyda kelgan id bilan bir xil id`li elementni indexini olamiz
	index := -1
	for p := 0; p < len(posts); p++ {
		if req.ID == posts[p].ID {
			index = p
		}
	}

	if index == -1 {
		fmt.Println("Bunday id li element yo'q")
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "Bunday id li element yo'q")
		return
	}

	// 5. o'sha elementni arraydan o'chirib, yangi elementni qo'shamiz

	posts = append(posts[:index], posts[index+1:]...)
	req.UpdatedAt = time.Now()
	posts = append(posts, req)

	// 6. hosil bo'lgan arrayni marshall qilib faylga yozamiz
	postsJson, _ := json.Marshal(posts)
	os.WriteFile("db/posts.json", postsJson, 0)
	// 7. response`da jo'natamiz
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}

func deletePost(w http.ResponseWriter, r *http.Request) {

}
