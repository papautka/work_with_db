package users

import (
	"errors"
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

	/*
		| user_id | user_name       | last_login          | post_id | post_likes | post_created_at     |
		|---------|----------------|-------------------- |---------|------------|-------------------- |
		| 1       | Алексей Петров | 2024-12-20 14:30:00 | 101     | 45         | 2024-12-15 10:00:00 |
		| 1       | Алексей Петров | 2024-12-20 14:30:00 | 102     | 123        | 2024-12-18 16:30:00 |
		| 1       | Алексей Петров | 2024-12-20 14:30:00 | 103     | 67         | 2024-12-19 09:15:00 |
		| 2       | Мария Иванова  | 2024-12-21 08:45:00 | NULL    | NULL       | NULL                |
		| 3       | Дмитрий Сидоров| 2024-12-20 22:12:30 | 201     | 5847       | 2024-12-10 12:00:00 |
	*/
	// 2. Нашей задачей будет собрать данные из плоских файлов
	/*
				go[]User{
			    {ID: 1, Name: "Алексей Петров", Posts: [Post{101}, Post{102}, Post{103}]},
			    {ID: 2, Name: "Мария Иванова", Posts: []},
			    {ID: 3, Name: "Дмитрий Сидоров", Posts: [Post{201}]},
			}
		[1, 101; 1, 102; 1, 103; 2, NULL; 3, 201]
	*/
	var users []models.User
	userIndexMap := make(map[uint64]int) // Карта ID пользователя -> индекс в слайсе
	for _, user := range rawUsers {
		userIndex, exists := userIndexMap[user.ID]

		if !exists {
			// создаем нового пользователя
			newUser := models.User{
				ID:        user.ID,
				Name:      user.Name,
				LastLogin: user.LastLogin,
				Posts:     make([]models.Post, 0),
			}

			users = append(users, newUser)
			userIndexMap[newUser.ID] = len(users) - 1
			userIndex = len(users) - 1
		}

		// добавляем пост
		if user.PostID != nil {
			post := models.Post{
				ID:        *user.PostID,
				UserID:    user.ID,
				Likes:     *user.PostLikes,
				CreatedAt: *user.PostCreatedAt,
			}
			users[userIndex].Posts = append(users[userIndex].Posts, post)
		}
	}
	return users, nil
}
