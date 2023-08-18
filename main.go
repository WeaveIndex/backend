package main

import (
	"backend/src"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	src.Serve(e)
}
