package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("ELEPHANT_DB_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	// .Exec will only know if it failled or not
	// .QueryRow will return 1 record // .Query will get resultSet
	// Go need RETURNING
	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", "Jack", "19")

	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't insert data", err)
	}

	fmt.Print("insert user success id: ", id)
}
