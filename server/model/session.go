package model

import "gorm.io/gorm"

type SessionKey struct {
	gorm.Model
	// id         int
	UserId     int
	SessionKey string
	// created_at time.Time
	// deleted_at time.Time
	// User    []Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User []Users
}
