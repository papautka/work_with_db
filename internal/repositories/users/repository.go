package users

import (
	"errors"
	"work_with_db/internal/dbs/postgres"
	"work_with_db/internal/entities"
)

type Repository struct {
	db postgres.Db
}

func NewRepository(db *postgres.Db) *Repository {
	return &Repository{
		db: *db,
	}
}

func (r *Repository) GetAllUsers() ([]entities.User, error) {
	users := make([]entities.User, 0)
	err := r.db.MySQL.Select(&users, "SELECT id, name, last_login FROM users")
	if err != nil {
		return nil, errors.New("could not get all users" + err.Error())
	}
	return users, nil
}
