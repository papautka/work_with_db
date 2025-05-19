package users

import (
	"errors"
	"fmt"
	"strings"
	"work_with_db/internal/dbs/postgres"
	"work_with_db/internal/entities"
	"work_with_db/internal/models"
)

type Repository struct {
	db postgres.Db
}

func NewRepository(db *postgres.Db) *Repository {
	return &Repository{
		db: *db,
	}
}

/*
users
| id| name           | last_login          |
|---|----------------|-------------------- |
| 1 | Алексей Петров | 2024-12-20 14:30:00 |
| 2 | Мария Иванова  | 2024-12-21 08:45:00 |
| 3 | Дмитрий Сидоров| 2024-12-20 22:12:30 |
*/

/*
posts
| id  | user_id | likes | created_at          |
|-----|---------|-------|-------------------- |
| 101 | 1       | 45    | 2024-12-15 10:00:00 |
| 102 | 1       | 123   | 2024-12-18 16:30:00 |
| 103 | 1       | 67    | 2024-12-19 09:15:00 |
| 201 | 3       | 5847  | 2024-12-10 12:00:00 |
*/
func (r *Repository) GetAllUsers() ([]entities.User, error) {
	users := make([]entities.User, 0)
	err := r.db.MySQL.Select(&users, `SELECT id, name, last_login FROM users`)
	if err != nil {
		return nil, errors.New("could not get all users" + err.Error())
	}
	return users, nil
}

func (r *Repository) GetAllUsersWithPosts() ([]models.User, error) {
	// 1. Получаем плоские данные
	const query = `
		SELECT
			users.id AS user_id, 
			users.name AS user_name,
			users.last_login AS user_last_login, 
			posts.id AS post_id,
			posts.likes AS post_likes, 
			posts.created_at AS post_created_at
		FROM users
		LEFT JOIN posts ON users.id = posts.user_id
		ORDER BY users.id, posts.created_at DESC
	`

	rawUsers := make([]entities.UserWithPosts, 0)
	err := r.db.MySQL.Select(&rawUsers, query)
	if err != nil {
		return nil, errors.New("failed to execute query: " + err.Error())
	}

	if len(rawUsers) == 0 {
		return []models.User{}, nil // Возвращаем пустой слайс, а не nil
	}
	return models.NewUser(rawUsers), nil
}

func (r *Repository) GetUserByIDs(ids []uint64) ([]models.User, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	args := make([]interface{}, 0, len(ids))
	inParams := make([]string, 0, len(ids))
	for i, id := range ids {
		args = append(args, interface{}(id))
		inParams = append(inParams, fmt.Sprintf("$%d", i+1))
	}
	query := fmt.Sprintf(`
		SELECT
			users.id AS user_id, users.name AS user_name,
            users.last_login AS user_last_login,
            posts.id AS post_id, posts.likes AS post_likes,
            posts.created_at AS post_created_at
        FROM users
		    LEFT JOIN posts ON users.id = posts.user_id
		WHERE users.id IN (%s)
		ORDER BY users.last_login DESC, posts.created_at DESC     
	`, strings.Join(inParams, ","))
	rawUsers := make([]entities.UserWithPosts, 0)
	err := r.db.MySQL.Select(&rawUsers, query, args...)
	if err != nil {
		return nil, errors.New("failed to execute query: " + err.Error())
	}
	return models.NewUser(rawUsers), nil
}
