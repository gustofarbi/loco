package main

import (
	"github.com/labstack/echo/v4"
	"loco/api"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error { return c.String(200, "Hello, World!") })
	e.POST("/api/v1/translations", api.CreateTranslationKeys)
	e.GET("/api/v1/translations", api.GetTranslationKey)
	e.DELETE("/api/v1/translations", api.DeleteTranslationKeys)

	e.Logger.Fatal(e.Start(":8080"))
}
