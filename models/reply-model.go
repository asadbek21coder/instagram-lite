package models

import "time"

type Reply struct {
	ID        int
	CommentId int
	PostId    int
	AuthorId  int
	Body      string
	LikeCount int
	CreatedAt time.Time
	UpdatedAt time.Time
}
