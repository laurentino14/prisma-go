package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/laurentino14/prismago/prisma"
)

func Hello(c echo.Context) error {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	users, err := client.User.FindMany().Exec(c.Request().Context())
	if err != nil {
		return c.JSONPretty(400, err, " ")
	}

	return c.JSONPretty(200, users, " ")

}

type User struct {
	FirstName string `json:"firstname"`
	LarstName string `json:"lastname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

// c echo.Context
func Create(c echo.Context) error {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	user := User{}
	c.Bind(&user)

	created, err := client.User.CreateOne(
		prisma.User.FirstName.Set(user.FirstName),
		prisma.User.LastName.Set(user.LarstName),
		prisma.User.Email.Set(user.Email),
		prisma.User.Age.Set(user.Age),
	).Exec(c.Request().Context())
	if err != nil {
		return c.JSONPretty(400, err.Error(), " ")
	}
	return c.JSONPretty(200, created, " ")

}
