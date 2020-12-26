package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func SqlConnection() *sql.DB {
	user, pass, dbname := "postgres", "1234", "test"

	url := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, dbname)

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Printf("Error an open db %s\n", fmt.Sprint(err))
	}

	return db
}
