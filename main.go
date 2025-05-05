package main

import (
	db2 "github.com/MatheusVict/streamer-app/config/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"io"
	"log"
	"net/http"
)

func main() {
	_, err := db2.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's working baby")
	})

	e.POST("/auth", func(c echo.Context) error {
		log.Default().Println("POST /auth")
		body := c.Request().Body
		defer body.Close()

		fields, _ := io.ReadAll(body)
		color.Println(string(fields))
		return c.String(http.StatusOK, "It's working baby")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
