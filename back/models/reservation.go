package models

import "github.com/google/uuid"

type Reservation struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null"`
	SlotID     uuid.UUID `gorm:"type:uuid;not null"`
	Slot       Slot      `gorm:"foreignKey:SlotID"`
	User       User      `gorm:"foreignKey:CustomerID"`
}
