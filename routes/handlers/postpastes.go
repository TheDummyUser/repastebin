package handlers

import (
	"net/http"
	"pastebin/db"
	"pastebin/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(post.Password), cost)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "something went wrong while hashing password",
			"error":   err.Error(),
		})
	}
	diffSave := &models.Pastes{
		Title:       post.Title,
		Content:     post.Content,
		Password:    string(hashedPassword),
		IsAnon:      post.IsAnon,
		OneTimeView: post.OneTimeView,
	}
	res := db.DB.Save(diffSave)
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
