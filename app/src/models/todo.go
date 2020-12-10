package models

// Todo model
type Todo struct {
	ID   int    `gorm:"primaryKey"`
	UUID string `gorm:"type:varchar(36);unique;not null"`
	Items []Item
}
