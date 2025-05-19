package entities

import "time"

type User struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	LastLogin time.Time `db:"last_login"`
}
