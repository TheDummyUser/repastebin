package models

import "gorm.io/gorm"

type Pastes struct {
	gorm.Model
	Title       string `json:"title"`
	Content     string `json:"content"`
	OneTimeView bool   `gorm:"column:one_time_view" json:"one_time_view"`
	IsAnon      bool   `gorm:"column:is_anon" json:"is_anon"`
	Password    string `gorm:"column:password" json:"password"`
}
