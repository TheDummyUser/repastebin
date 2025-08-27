package routes

import (
	"pastebin/routes/handlers"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return handlers.GetPastes(c)
	})
	e.POST("/add", func(c echo.Context) error {
		return handlers.PostPastes(c)
	})

	e.GET("/:id", func(c echo.Context) error {
		return handlers.GetSinglePaste(c)
	})
}
