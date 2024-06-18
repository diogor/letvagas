package models

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PositionType string
type Allocation string
type ContractType string
type Level string

const (
	TEMPORARY PositionType = "temporary"
	CONTRACT  PositionType = "contract"
	LONG_TERM PositionType = "long_term"
)

const (
	REMOTE  Allocation = "remote"
	ON_SITE Allocation = "on_site"
	HYBRID  Allocation = "hybrid"
)

const (
	CLT     ContractType = "clt"
	PJ      ContractType = "pj"
	SERVICE ContractType = "service"
)

const (
	INTERNSHIP Level = "internship"
	JUNIOR     Level = "junior"
	MID        Level = "mid"
	SENIOR     Level = "senior"
	SPECIALIST Level = "specialist"
)

func (ct *PositionType) Scan(value interface{}) error {
	*ct = PositionType(value.(string))
	return nil
}

func (ct PositionType) Value() (driver.Value, error) {
	return string(ct), nil
}

func (ct *Allocation) Scan(value interface{}) error {
	*ct = Allocation(value.(string))
	return nil
}

func (ct Allocation) Value() (driver.Value, error) {
	return string(ct), nil
}

func (ct *ContractType) Scan(value interface{}) error {
	*ct = ContractType(value.(string))
	return nil
}

func (ct ContractType) Value() (driver.Value, error) {
	return string(ct), nil
}

func (ct *Level) Scan(value interface{}) error {
	*ct = Level(value.(string))
	return nil
}

func (ct Level) Value() (driver.Value, error) {
	return string(ct), nil
}

type Position struct {
	gorm.Model
	ID          uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Company     string       `json:"company" gorm:"type:varchar(255);not null"`
	Title       string       `json:"title" gorm:"type:varchar(255);not null"`
	Level       Level        `json:"level" sql:"type:level('internship', 'junior', 'mid', 'senior', 'specialist');not null"`
	Type        PositionType `json:"type" sql:"type:position_type('temporary', 'contract', 'long_term');not null"`
	Allocation  Allocation   `json:"allocation" sql:"type:allocation('remote', 'on_site', 'hybrid');not null"`
	Wage        *string      `json:"wage" gorm:"type:varchar(255)"`
	Contract    ContractType `json:"contract" sql:"type:contract_type('clt', 'pj', 'service');not null"`
	Location    string       `json:"location" gorm:"type:varchar(255)"`
	Description string       `json:"description" gorm:"type:text"`
	PCD         bool         `json:"pcd" gorm:"not null;default:false"`
	CreatedBy   Profile      `json:"created_by" gorm:"foreignKey:CreatedByID"`
	CreatedByID uuid.UUID    `json:"created_by_id" gorm:"type:uuid;not null"`
}

type Application struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Position   Position  `json:"position" gorm:"foreignKey:PositionId"`
	PositionId uuid.UUID `json:"position_id" gorm:"type:uuid;not null;uniqueIndex:application_idx"`
	Profile    Profile   `json:"profile" gorm:"foreignKey:ProfileId"`
	ProfileId  uuid.UUID `json:"profile_id" gorm:"type:uuid;not null;uniqueIndex:application_idx"`
}
