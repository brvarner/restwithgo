package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	connStr := "user=" + os.Getenv("DB_USER") +
			" dbname=" + os.Getenv("DB_NAME") +
			" sslmode=" + os.Getenv("DB_SSLMODE") +
			" password=" + os.Getenv("DB_PASSWORD") +
			" host=" + os.Getenv("DB_HOST") +
			" port=" + os.Getenv("DB_PORT")


	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("Database connected successfully")
	createTables()
}

func createTables() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}