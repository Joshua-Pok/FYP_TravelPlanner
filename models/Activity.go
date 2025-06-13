package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	ItineraryId  uint   `gorm:"not null"`
	Name         string `gorm:"type: varchar(255); not null"`
	Description  string `gorm:"type:text; not null"`
	ActivityType string `gorm:"type: varchar(255); not null"`
	Duration     int    `gorm:"type: default:60"`
	Location     string `gorm:"type:varchar(255); not null"`
}
