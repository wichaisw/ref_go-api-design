package user

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("ELEPHANT_DB_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	createTb := `CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT, age INT);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table", err)
	}

	log.Println("okay")
}
