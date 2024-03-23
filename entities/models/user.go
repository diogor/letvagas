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
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Cpf        string    `json:"cpf" gorm:"type:varchar(15);not null"`
	Name       string    `json:"name" gorm:"type:varchar(255);not null"`
	BirthDate  string    `json:"birth_date" gorm:"type:date;not null"`
	Password   string    `json:"-" gorm:"type:varchar(255);not null"`
	Email      string    `json:"email" gorm:"type:varchar(255);not null"`
	Phone      string    `json:"phone" gorm:"type:varchar(255);not null"`
	Role       UserRole  `json:"role" sql:"type:user_role('admin', 'client', 'applicant');not null;default:'applicant'"`
	SocialName *string   `json:"social_name" gorm:"type:varchar(255)"`
	Linkedin   *string   `json:"linkedin" gorm:"type:varchar(255)"`
	City       *string   `json:"city" gorm:"type:varchar(255)"`
	State      *string   `json:"state" gorm:"type:varchar(255)"`
	Address    *string   `json:"address" gorm:"type:varchar(255)"`
	Cep        *string   `json:"cep" gorm:"type:varchar(255)"`
}
