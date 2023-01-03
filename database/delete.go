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

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		log.Fatal("can't prepare query all users statement", err)
	}

	// use .Exec we just want to know if it success or not, don't want anything to be returned
	if _, err := stmt.Exec(1); err != nil {
		log.Fatal("error execute delete: ", err)
	}

	fmt.Println("delete success")
}
