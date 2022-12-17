package model

import (
	"crypto/sha256"
	"encoding/hex"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// id         int
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// created_at time.Time
	// Users    []Users
	Postit     []Postit     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SessionKey []SessionKey `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Passwordをハッシュ化する
func (u *User) ToHash() {
	// r := sha256.Sum256([]byte(s))
	// return r[:]
	nohash := u.Password
	sha256 := sha256.Sum256([]byte(nohash))
	hashed := hex.EncodeToString(sha256[:])
	u.Password = hashed
}
