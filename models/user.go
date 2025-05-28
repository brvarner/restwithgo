package models

import (
	"database/sql"
	"user-management-api/database"
)

type User struct {
	ID		  int	 `json:"id"`
	Name	  string `json:"name"`
	Email 	  string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func GetAllUsers() ([]User, error) {
	query := "SELECT id, name, email, created_at FROM users ORDER BY id"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(user User) (*User, error) {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at"

	err := database.DB.QueryRow(query, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}