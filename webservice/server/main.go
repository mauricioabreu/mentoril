package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/*", func(c echo.Context) error {
		path := c.Request().URL.Path
		return c.String(http.StatusOK, fmt.Sprintf("Requested path: %s", path))
	})
	e.Logger.Fatal(e.Start(":8080"))
}
