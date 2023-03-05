package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Name     string `gorm:"not null" json:"name"`
	Surname  string `gorm:"not null" json:"surname"`
	Password string `gorm:"not null" json:"password"`
	Mail     string `gorm:"unique;not null" json:"mail"`
	Tel      string `gorm:"unique;not null" json:"tel"`
	Address  string `json:"address"`
}
