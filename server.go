package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
}

func postCreateUserHandler(context echo.Context) error {
	user := new(User)
	err := context.Bind(user)
	if err != nil {
		return err
	}

	return context.JSON(
		http.StatusOK,
		user,
	)
}

func main() {
	e := echo.New()
	e.POST("/users/create", postCreateUserHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
