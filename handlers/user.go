package handlers

import (
	"encoding/json"
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
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	var newuser models.User
	json.NewDecoder(r.Body).Decode(&newuser)
	readUsers, _ := os.ReadFile("db/users.json")
	json.Unmarshal(readUsers, &users)

	newuser.ID = len(users) + 1 // logikasini tog'irlash kk

	// qo'shishdan oldin username band emasligini tekshirish, band bo'lsa error qaytarish kk
	users = append(users, newuser)
	data, _ := json.Marshal(users)

	os.WriteFile("db/users.json", data, 0)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

}
