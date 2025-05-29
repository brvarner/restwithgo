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

func GetUserByID(id int) (*User, error) {
	query := "SELECT id, name, email, created_at FROM users WHERE id = $1"

	var user User
	err := database.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func UpdateUser(id int, user User) (*User, error) {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email, created_at"

	var updatedUser User
	err := database.DB.QueryRow(query, user.Name, user.Email, id).Scan(
		&updatedUser.ID, &updatedUser.Name, &updatedUser.Email, &updatedUser.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
