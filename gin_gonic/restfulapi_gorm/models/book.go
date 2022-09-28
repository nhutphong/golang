package models

import (
	"github.com/jinzhu/gorm"
)


type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}
