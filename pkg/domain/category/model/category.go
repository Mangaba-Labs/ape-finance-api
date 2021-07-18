package model

import "gorm.io/gorm"

// Category model in database
type Category struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
}
