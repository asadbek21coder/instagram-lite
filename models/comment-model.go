package models

import "time"

type Comment struct {
	ID              int
	CommentedPostId int
	AuthorId        int
	Body            string
	LikeCount       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CommentWithreply struct {
	Comment Comment
	Replies []Reply
}
