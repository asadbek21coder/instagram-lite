package models

import "time"

type Post struct {
	ID        int
	Title     string
	Body      string
	ImageUrl  string
	AuthorId  int
	LikeCount int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetAllPosts struct {
	ID        int
	Title     string
	Body      string
	ImageUrl  string
	Author    User
	Comments  []CommentWithreply
	LikeCount int
	CreatedAt time.Time
	UpdatedAt time.Time
}
