package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	config := LoadConfig()
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + config.DBName)
	return db, nil
}