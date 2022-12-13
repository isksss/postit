package database

import "gorm.io/gorm"

func HasTable(db *gorm.DB, table interface{}) bool {
	if db.Migrator().HasTable(table) {
		return true
	} else {
		return false
	}
}
