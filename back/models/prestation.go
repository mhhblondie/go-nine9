package models

import "github.com/google/uuid"

type Prestation struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Price     float64    `gorm:"type:double;not null"`
	ServiceID *uuid.UUID `gorm:"type:uuid"`
}
