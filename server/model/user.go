package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// id         int
	Name     string
	Email    string
	Password string
	// created_at time.Time
	// Users    []Users
	Postit     []Postit     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SessionKey []SessionKey `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
