package main

import (
	"github.com/WeaveIndex/backend/src"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	src.Serve(e)
}
