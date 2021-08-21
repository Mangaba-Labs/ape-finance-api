package model

import "gorm.io/gorm"

// Category model in database
type Category struct {
	gorm.Model
	IDUser int    `gorm:"not null" json:"id_user"`
	Name   string `gorm:"not null" json:"name"`
	Type   string `gorm:"not null" json:"type"`
}

// CategoryResponse for API Requests
type CategoryResponse struct {
	ID     int    `json:"id"`
	IDUser int    `json:"id_user"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

// ParseFromDatabase format category for response
func (c *CategoryResponse) ParseFromDatabase(category Category) {
	c.ID = int(category.ID)
	c.IDUser = int(category.IDUser)
	c.Name = category.Name
	c.Type = category.Type
}
