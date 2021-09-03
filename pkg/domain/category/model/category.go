package model

import "gorm.io/gorm"

// Category model in database
type Category struct {
	gorm.Model
	IDUser uint64 `gorm:"not null" json:"id_user"`
	Name   string `gorm:"not null" json:"name"`
	Type   string `gorm:"not null" json:"type"`
}

// CategoryResponse for API Requests
type CategoryResponse struct {
	ID     uint64 `json:"id"`
	IDUser uint64 `json:"id_user"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

// ParseFromDatabase format category for response
func (c *CategoryResponse) ParseFromDatabase(category Category) {
	c.ID = uint64(category.ID)
	c.IDUser = uint64(category.IDUser)
	c.Name = category.Name
	c.Type = category.Type
}
