package models

import (
	"time"
	"work_with_db/internal/entities"
)

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

func NewUser(rawUsers []entities.UserWithPosts) []User {
	var users []User
	userIndexMap := make(map[uint64]int) // Карта ID пользователя -> индекс в слайсе
	for _, user := range rawUsers {
		userIndex, exists := userIndexMap[user.ID]

		if !exists {
			// создаем нового пользователя
			newUser := User{
				ID:        user.ID,
				Name:      user.Name,
				LastLogin: user.LastLogin,
				Posts:     make([]Post, 0),
			}

			users = append(users, newUser)
			userIndexMap[newUser.ID] = len(users) - 1
			userIndex = len(users) - 1
		}

		// добавляем пост
		if user.PostID != nil {
			post := Post{
				ID:        *user.PostID,
				UserID:    user.ID,
				Likes:     *user.PostLikes,
				CreatedAt: *user.PostCreatedAt,
			}
			users[userIndex].Posts = append(users[userIndex].Posts, post)
		}
	}
	return users
}
