package handlers

import "net/http"

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

}

func getReplies(w http.ResponseWriter, r *http.Request) {

}

func updateReply(w http.ResponseWriter, r *http.Request) {

}

func deleteReply(w http.ResponseWriter, r *http.Request) {

}
