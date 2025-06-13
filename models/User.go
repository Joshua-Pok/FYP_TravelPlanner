package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string      `gorm:"type: varchar(255); unique; not null"`
	Email       string      `gorm:"type: varchar(255); unique; not null"`
	Password    string      `gorm:"type: varchar(255); not null"`
	Personality Personality `gorm:"foreignKey:UserID;"`
}
