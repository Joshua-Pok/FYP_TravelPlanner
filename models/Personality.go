package models

import (
	"gorm.io/gorm"
)

type Personality struct {
	gorm.Model
	UserID            uint    `gorm:"unique; not null; index"`
	Openness          float32 `gorm:"type: numeric(4,2); default:0.0;"`
	Conscientiousness float32 `gorm:"type:numeric(4,2); default:0.0"`
	Extraversion      float32 `gorm:"type: numeric(4,2); default:0.0"`
	Agreeableness     float32 `gorm:"type: numeric(4,2); default:0.0"`
	Neuroticism       float32 `gorm:"type: numeric(4,2); default:0.0"`
}
