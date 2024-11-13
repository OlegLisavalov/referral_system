package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	//ReferralCode string `json:"referral_code" gorm:"unique"`
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
