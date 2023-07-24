package models

type User struct {
	ID          int
	Username    string
	Age         int
	PhoneNumber string
	Gender      bool
}

var Users []User
