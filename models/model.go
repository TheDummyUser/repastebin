package models

import "gorm.io/gorm"

type Pastes struct {
	gorm.Model
	Title       string `json:"title"`
	Content     string `json:"content"`
	OneTimeView bool   `gorm:"column:one_time_view" json:"one_time_view"`
	IsAnon      bool   `gorm:"column:is_anon" json:"is_anon"`
	Password    string `gorm:"column:password" json:"-"`
}

type PasteResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	OneTimeView bool   `json:"one_time_view"`
	IsAnon      bool   `json:"is_anon"`
}

type PasteRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	// optional
	OneTimeView bool   `json:"one_time_view"`
	IsAnon      bool   `json:"is_anon"`
	Password    string `json:"password"`
}

type PastePasswordRequest struct {
	ID       uint   `json:"id"`
	Password string `json:"password"`
}
