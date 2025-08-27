package main

import (
	"fmt"
	"pastebin/db"
	"pastebin/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := db.Initdb()
	if err != nil {
		fmt.Println("something is fishy", err)
	}
	e := echo.New()
	e.Use(middleware.CORS())
	routes.SetUpRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
