package models

import (
	"testing"
	"user-management-api/database"
)

func TestCreateUser(t *testing.T) {
	database.InitDB()

	user := User{
		Name: "Test User",
		Email: "test@example.com",
	}

	createdUser, err := CreateUser(user) 
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if createdUser.ID == 0 {
		t.Error("Expected user ID to be set")
	}

	if createdUser.Name != user.Name {
		t.Errorf("Expected name %s, got %s", user.Name, createdUser.Name)
	}
}