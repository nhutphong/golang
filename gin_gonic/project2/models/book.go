package models

import "github.com/jinzhu/gorm"

type Book struct {
	// ID     uint   `json:"id" gorm:"primary_key"`
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}
