package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "golang"
)

func selectUser(db *sql.DB) {
	email := "1@ee.s"
	rows, err := db.Query("SELECT username FROM accounts WHERE email=$1", email)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s's email is %s\n", username, email)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	selectUser(db)
}
