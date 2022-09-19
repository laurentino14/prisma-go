package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/laurentino14/teste/teste"
)

func Hello(c echo.Context) error {
	client := teste.NewClient()

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
		return c.JSON(400, err)
	}

	return c.JSON(200, users)

}

func Create(c echo.Context) error {
	client := teste.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	created, err := client.User.CreateOne(
		teste.User.FirstName.Set("first_name"),
		teste.User.LastName.Set("last_name"),
		teste.User.Email.Set("email"),
		teste.User.Age.Set("age"),
	).Exec(c.Request().Context())
	if err != nil {
		return c.JSON(400, created)
	}
	return c.JSON(200, created)

}
