package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          uint   ` gorm:"primary_key" json:id`
	Title       string `json:Title`
	Description string `json:Description`
	Cost        string `json:Cost`
}
