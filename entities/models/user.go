package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name     string    `json:"name" gorm:"type:varchar(255);not null"`
	Password string    `json:"-" gorm:"type:varchar(255);not null"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null"`
}
