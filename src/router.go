package src

import (
	"backend/src/pages"
	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {
	e.GET("/", pages.ServeIndexPage)

	e.Logger.Fatal(e.Start(":3031"))
}
