package db

import (
	"database/sql"
	"fmt"
	"go-api/config"

	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	pgsqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open("postgres", pgsqlInfo)
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
