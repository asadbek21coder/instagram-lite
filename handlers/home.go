package handlers

import (
	"fmt"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Welcome to hom page")
}
