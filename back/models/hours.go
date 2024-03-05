package models

import (
	"time"

	"github.com/google/uuid"
)

type Hours struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key"`
	SalonID     *uuid.UUID `gorm:"type:uuid"`
	DayOfWeek   string     `gorm:"type:varchar(255);not null"`
	OpeningTime time.Time  `gorm:"type:timestamp;not null"`
	ClosingTime time.Time  `gorm:"type:timestamp;not null"`
}
