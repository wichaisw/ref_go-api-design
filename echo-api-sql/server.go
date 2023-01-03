package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"nickname"`
	Age  int    `json:"age"`
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func getUsersHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all users statement"})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all users"})
	}

	users := []User{}
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			// shouldn't log db error to client
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user: " + err.Error()})
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

type Err struct {
	Message string `json:"message"`
}

func createUsersHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", u.Name, u.Age)

	err = row.Scan(&u.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, u)
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("ELEPHANT_DB_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	g := e.Group("/api")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "apidemo" && password == "34567" {
			return true, nil
		}

		return false, nil
	}))

	g.GET("/users", getUsersHandler)
	g.POST("/users", createUsersHandler)

	log.Println("Server started at :2565")
	// log.Fatal(http.ListenAndServe(":2565", mux))
	log.Fatal(e.Start(":2565"))
	log.Println("bye")
}
