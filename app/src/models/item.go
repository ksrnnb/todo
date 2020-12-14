package models

// Item model
type Item struct {
	ID     int    `gorm:"primaryKey"`
	TodoID int    `gorm:"not null"`
	Todo   Todo
	Name   string `gorm:"not null"`
	Done   bool   `gorm:"default:false;not null"`
}

func (item Item) Delete() {
	Db.Delete(&item)
}
