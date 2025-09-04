package handlers

import (
	"net/http"
	"pastebin/db"
	"pastebin/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func PostPastes(c echo.Context) error {
	post := new(models.PasteRequest)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request body",
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
	}

	if post.Title == "" || post.Content == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "title and content are required",
			"status":  http.StatusBadRequest,
		})
	}

	var hashedPassword string
	if post.Password != "" {
		hp, err := bcrypt.GenerateFromPassword([]byte(post.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed to hash password",
				"error":   err.Error(),
			})
		}
		hashedPassword = string(hp)
	}

	diffSave := &models.Pastes{
		Title:       post.Title,
		Content:     post.Content,
		Password:    hashedPassword,
		IsAnon:      post.IsAnon,
		OneTimeView: post.OneTimeView,
	}
	res := db.DB.Save(diffSave)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to save paste",
			"status":  http.StatusInternalServerError,
			"error":   res.Error.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "paste added successfully",
		"status":  http.StatusCreated,
		"id":      diffSave.ID,
	})
}
