package models

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	ADMIN     UserRole = "admin"
	CLIENT    UserRole = "client"
	APPLICANT UserRole = "applicant"
)

func (ct *UserRole) Scan(value interface{}) error {
	*ct = UserRole(value.(string))
	return nil
}

func (ct UserRole) Value() (driver.Value, error) {
	return string(ct), nil
}

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name     string    `json:"name" gorm:"type:varchar(255);not null"`
	Password string    `json:"-" gorm:"type:varchar(255);not null"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null"`
	Role     UserRole  `json:"role" sql:"type:user_role('admin', 'client', 'applicant');not null;default:'applicant'"`
}
