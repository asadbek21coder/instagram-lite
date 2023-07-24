package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/asadbek21coder/instagram/models"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createUser(w, r)
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	var newuser models.User
	json.NewDecoder(r.Body).Decode(&newuser)
	readUsers, _ := os.ReadFile("db/users.json")
	json.Unmarshal(readUsers, &users)

	if len(users) == 0 {
		newuser.ID = 1
	} else {
		var max int = users[0].ID
		for i := 0; i < len(users); i++ {
			if users[i].ID > max {
				max = users[i].ID
			}
		}

		newuser.ID = max + 1
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == newuser.Username {
			fmt.Println("This username is already taken")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "error: This username is already taken")
			return
		}
	}
	users = append(users, newuser)
	data, _ := json.Marshal(users)

	os.WriteFile("db/users.json", data, 0)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newuser)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	read, _ := os.ReadFile("db/users.json")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(read))
}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}
