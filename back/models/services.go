package models

import "github.com/google/uuid"

type Service struct {
	ID          uuid.UUID    `gorm:"type:uuid;primary_key"`
	Name        string       `gorm:"type:varchar(255);not null"`
	Description string       `gorm:"type:text"`
	SalonID     *uuid.UUID   `gorm:"type:uuid"`
	Prestation  []Prestation `gorm:"foreignKey:ServiceID"`
}
