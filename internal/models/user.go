package models

import (
	"context"

	"github.com/ismail-hafiane/goinfra/internal/database"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers() ([]User, error) {
	rows, err := database.Pool.Query(context.Background(),
		"SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func CreateUser(name, email string) (User, error) {
	var u User
	err := database.Pool.QueryRow(context.Background(),
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email",
		name, email,
	).Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}
