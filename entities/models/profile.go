package models

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EducationType string

const (
	FUNDAMENTAL EducationType = "fundamental"
	MEDIO       EducationType = "medio"
	SUPERIOR    EducationType = "superior"
	POS         EducationType = "pos"
	MBA         EducationType = "mba"
	OUTROS      EducationType = "outros"
)

func (ct *EducationType) Scan(value interface{}) error {
	*ct = EducationType(value.(string))
	return nil
}

func (ct EducationType) Value() (driver.Value, error) {
	return string(ct), nil
}

type Profile struct {
	gorm.Model
	ID         uuid.UUID   `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserId     uuid.UUID   `json:"user_id" gorm:"type:uuid;not null"`
	User       User        `json:"user" gorm:"foreignKey:UserId"`
	PCD        bool        `json:"pcd" gorm:"not null;default:false"`
	Educations []Education `json:"educations" gorm:"foreignKey:ProfileId"`
	Courses    []Course    `json:"courses" gorm:"foreignKey:ProfileId"`
}

type Education struct {
	gorm.Model
	ID          uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProfileId   uuid.UUID     `json:"profile_id" gorm:"type:uuid;not null"`
	Profile     Profile       `json:"profile" gorm:"foreignKey:ProfileId"`
	Type        EducationType `json:"type" sql:"type:education_type('fundamental', 'medio', 'superior', 'pos', 'mba', 'outros');not null"`
	Institution string        `json:"institution" gorm:"type:varchar(255);not null"`
	Graduation  *string       `json:"graduation" gorm:"type:varchar(255)"`
	StartYear   *string       `json:"start_year" gorm:"type:varchar(255)"`
	EndYear     *string       `json:"end_year" gorm:"type:varchar(255)"`
	StartMonth  *string       `json:"start_month" gorm:"type:varchar(255)"`
	EndMonth    *string       `json:"end_month" gorm:"type:varchar(255)"`
	Description *string       `json:"description" gorm:"type:varchar(255)"`
	Ongoing     bool          `json:"ongoing" gorm:"not null;default:false"`
}

type Course struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProfileId   uuid.UUID `json:"profile_id" gorm:"type:uuid;not null"`
	Profile     Profile   `json:"profile" gorm:"foreignKey:ProfileId"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Year        *string   `json:"year" gorm:"type:varchar(255)"`
	Month       *string   `json:"month" gorm:"type:varchar(255)"`
	Description *string   `json:"description" gorm:"type:varchar(255)"`
	Ongoing     bool      `json:"ongoing" gorm:"not null;default:false"`
}
