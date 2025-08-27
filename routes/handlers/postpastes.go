package handlers

import (
	"net/http"
	"pastebin/db"
	"pastebin/models"

	"github.com/labstack/echo/v4"
)

func PostPastes(c echo.Context) error {
	post := new(models.Pastes)
	if err := c.Bind(post); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{
				"message": "something went wrong",
				"status":  http.StatusInternalServerError,
			},
		)
	}
	res := db.DB.Save(post)
	if res.Error != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{
				"message": "something went wrong while adding data to the db",
				"status":  http.StatusInternalServerError,
				"err":     res.Error.Error(),
			},
		)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "added post successfully",
		"staus":   http.StatusAccepted,
		"id":      post.ID,
	})
}
