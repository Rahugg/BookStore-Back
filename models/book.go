package models

type Book struct {
	ID          uint   ` gorm:"primary_key" json:id`
	Title       string `json:title`
	Description string `json:description`
	Cost        string `json:cost`
}
