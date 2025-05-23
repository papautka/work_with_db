package entities

import "time"

// 1. SELECT
type User struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	LastLogin time.Time `db:"last_login"`
}

// 2. SELECT with JOIN
type UserWithPosts struct {
	ID            uint64     `db:"user_id"`
	Name          string     `db:"user_name"`
	LastLogin     time.Time  `db:"user_last_login"`
	PostID        *uint64    `db:"post_id"`
	PostLikes     *int64     `db:"post_likes"`
	PostCreatedAt *time.Time `db:"post_created_at"`
}
