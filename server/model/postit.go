package model

import (
	"gorm.io/gorm"
)

type Postit struct {
	gorm.Model
	Text   string
	UserID int
	User   []Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
