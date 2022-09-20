package http

import (
	"github.com/labstack/echo/v4"
	"github.com/laurentino14/prismago/http/routes"
)

var E = echo.New()

func Server() {

	E.GET("/list", routes.Hello)

	E.POST("/create", routes.Create)

	E.Start(":3131")
}
