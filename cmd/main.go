package main

import (
	"fmt"
	"net/http"

	"github.com/asadbek21coder/instagram/handlers"
)

const PORT = "9000"

func main() {
	// fmt.Println("Hello world")

	http.HandleFunc("/", handlers.GetHomePage)

	http.HandleFunc("/users", handlers.UserHandler)

	http.HandleFunc("/posts", handlers.PostHandler)

	http.HandleFunc("/comment", handlers.CommentHandler)

	http.HandleFunc("/reply", handlers.ReplyHandler)

	fmt.Println("Server is working on: " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
