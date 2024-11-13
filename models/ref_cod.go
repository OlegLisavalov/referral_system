package models

import (
	"gorm.io/gorm"
	"time"
)

type RefCode struct {
	gorm.Model
	Id     int       `gorm:"primaryKey;autoIncrement"`
	UserId int       `json:"user_id"`
	Code   string    `json:"code" gorm:"unique"`
	Expiry time.Time `json:"expiry"`
	Active bool      `json:"active"`
}
