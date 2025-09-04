package handlers

import (
	"net/http"
	"pastebin/db"
	"pastebin/models"

	"github.com/labstack/echo/v4"
)

func GetPastes(c echo.Context) error {
	var pastes []models.Pastes
	result := db.DB.
		Where("one_time_view = ? AND is_anon = ? AND password = ?", false, false, "").
		Find(&pastes)

	if result.Error != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{
				"message": "something went wrong while fetching data",
				"status":  http.StatusInternalServerError,
			},
		)
	}

	if len(pastes) == 0 {
		return c.JSON(
			http.StatusOK,
			map[string]interface{}{
				"message": "no data found",
				"status":  http.StatusOK,
				"data":    []string{},
			},
		)
	}

	var pastes_with_out_pass []models.PasteResponse

	for _, P := range pastes {
		pastes_with_out_pass = append(pastes_with_out_pass, models.PasteResponse{
			ID:          P.ID,
			Title:       P.Title,
			Content:     P.Content,
			OneTimeView: P.OneTimeView,
			IsAnon:      P.IsAnon,
		})
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"message": "fetched all the pastes",
			"status":  http.StatusOK,
			"data":    pastes_with_out_pass,
		},
	)
}

func GetSinglePaste(c echo.Context) error {
	var paste models.Pastes
	id := c.Param("id")
	data := db.DB.First(&paste, "id = ?", id)
	if data.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "something went wrong",
			"error":   data.Error.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	if paste.OneTimeView == true {
		res := db.DB.Delete(&paste)
		if res.Error != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "something went wrong",
				"error":   res.Error.Error(),
				"status":  http.StatusBadRequest,
			})
		}
	}

	if paste.Password != "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "sorry but to view this paste you need password",
			"status":  http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data fetched successfully",
		"paste":   paste,
	})
}
