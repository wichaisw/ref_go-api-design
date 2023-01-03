package user

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all users statement" + err.Error()})
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

func GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, name, age FROM users WHERE id =$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query user by id statement" + err.Error()})
	}

	row := stmt.QueryRow(id)
	u := User{}
	err = row.Scan(&u.Id, &u.Name, &u.Age)
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, u)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user: " + err.Error()})
	}

	return c.JSON(http.StatusOK, u)

}
