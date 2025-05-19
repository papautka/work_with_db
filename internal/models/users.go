package models

import "time"

type User struct {
	ID        uint64
	Name      string
	LastLogin time.Time
	Posts     []Post
}

type Post struct {
	ID        uint64
	UserID    uint64
	Likes     int64
	CreatedAt time.Time
}
