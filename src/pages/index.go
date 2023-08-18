package pages

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ServeIndexPage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
