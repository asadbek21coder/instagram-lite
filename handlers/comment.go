package handlers

import "net/http"

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

}

func getComments(w http.ResponseWriter, r *http.Request) {

}

func updateComment(w http.ResponseWriter, r *http.Request) {

}


func deleteComment(w http.ResponseWriter, r *http.Request) {

}
