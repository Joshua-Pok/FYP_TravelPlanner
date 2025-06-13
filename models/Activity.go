package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	ItineraryId uint   `gorm:"not null"`
	Description string "gorm"
}
