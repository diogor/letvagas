package models

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRole string

const (
	ADMIN     userRole = "admin"
	CLIENT    userRole = "client"
	APPLICANT userRole = "applicant"
)

func (ct *userRole) Scan(value interface{}) error {
	*ct = userRole(value.([]byte))
	return nil
}

func (ct userRole) Value() (driver.Value, error) {
	return string(ct), nil
}

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name     string    `json:"name" gorm:"type:varchar(255);not null"`
	Password string    `json:"-" gorm:"type:varchar(255);not null"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null"`
	Role     userRole  `json:"role" gorm:"type:user_role('admin', 'client', 'applicant');not null;default:'applicant'"`
}
