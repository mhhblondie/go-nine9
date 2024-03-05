package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Slot struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"type:uuid;primary_key"`
	SlotTime            time.Time `gorm:"type:timestamp;not null"`
	HairdressingStaffID uuid.UUID `gorm:"type:uuid;not null"`
	HairdressingStaff   User      `gorm:"foreignKey:HairdressingStaffID"`
	SalonID             uuid.UUID `gorm:"type:uuid;not null"`
}
