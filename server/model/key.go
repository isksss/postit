package model

import "gorm.io/gorm"

type Key struct {
	gorm.Model
	Key    string
	UserID int
	User   []Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
