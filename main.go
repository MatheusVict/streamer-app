package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's working baby")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
