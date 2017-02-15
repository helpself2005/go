package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Run(standard.New(":1323"))
}
