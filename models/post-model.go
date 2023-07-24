package models

import "time"

type Post struct {
	ID        int
	Title     string
	Body      string
	ImageUrl  string
	LikeCount int
	AuthorId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

var posts []Post
