package model

import "gorm.io/gorm"

type Postit struct {
	gorm.Model
	Id     int
	UserId int
	Text   string

	// User    []Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User []Users
}
