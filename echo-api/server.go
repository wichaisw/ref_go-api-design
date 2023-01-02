package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"nickname"`
	Age  int    `json:"age"`
}

var users = []User{
	{Id: 1, Name: "Kradas", Age: 12},
	{Id: 2, Name: "Noom", Age: 132},
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func getUsersHandler(c echo.Context) error {
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

	users = append(users, u)

	return c.JSON(http.StatusCreated, users)
}

func main() {
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
