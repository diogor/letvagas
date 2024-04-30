package models

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EducationType string
type SkillLevel string
type Intensity string

const (
	FUNDAMENTAL  EducationType = "fundamental"
	MEDIO        EducationType = "medio"
	SUPERIOR     EducationType = "superior"
	POS          EducationType = "pos"
	MBA          EducationType = "mba"
	OUTROS       EducationType = "outros"
	BASIC        SkillLevel    = "basic"
	INTERMEDIATE SkillLevel    = "intermediate"
	ADVANCED     SkillLevel    = "advanced"
	NONE         Intensity     = "none"
	LOW          Intensity     = "low"
	MEDIUM       Intensity     = "medium"
	HIGH         Intensity     = "high"
)

func (ct *EducationType) Scan(value interface{}) error {
	*ct = EducationType(value.(string))
	return nil
}

func (ct EducationType) Value() (driver.Value, error) {
	return string(ct), nil
}

func (ct *SkillLevel) Scan(value interface{}) error {
	*ct = SkillLevel(value.(string))
	return nil
}

func (ct SkillLevel) Value() (driver.Value, error) {
	return string(ct), nil
}

func (ct *Intensity) Scan(value interface{}) error {
	*ct = Intensity(value.(string))
	return nil
}

func (ct Intensity) Value() (driver.Value, error) {
	return string(ct), nil
}

type Profile struct {
	gorm.Model
	ID                 uuid.UUID              `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ComputingSkills    []ComputingSkillAnswer `json:"computing_skills" gorm:"foreignKey:ProfileId"`
	ComputingQuestions []ComputingAnswer      `json:"computing_questions" gorm:"foreignKey:ProfileId"`
	Educations         []Education            `json:"educations" gorm:"foreignKey:ProfileId"`
	Courses            []Course               `json:"courses" gorm:"foreignKey:ProfileId"`
	Languages          []LanguageAnswer       `json:"languages" gorm:"foreignKey:ProfileId"`
	Experiences        []Experience           `json:"experiences" gorm:"foreignKey:ProfileId"`
}

type ComputingSkill struct {
	gorm.Model
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string    `json:"name" gorm:"type:varchar(255);not null"`
}

type ComputingSkillAnswer struct {
	gorm.Model
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProfileId uuid.UUID      `json:"profile_id" gorm:"type:uuid;not null"`
	Profile   Profile        `json:"profile" gorm:"foreignKey:ProfileId"`
	Skill     ComputingSkill `json:"skill"`
	SkillId   uuid.UUID      `json:"skill_id" gorm:"type:uuid;not null"`
	Answer    SkillLevel     `json:"level" sql:"type:skill_level('basic', 'intermediate', 'advanced');not null"`
}

type ComputingQuestion struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Question string    `json:"question" gorm:"type:varchar(255);not null"`
}

type ComputingAnswer struct {
	gorm.Model
	ID         uuid.UUID         `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProfileId  uuid.UUID         `json:"profile_id" gorm:"type:uuid;not null"`
	Profile    Profile           `json:"profile" gorm:"foreignKey:ProfileId"`
	Question   ComputingQuestion `json:"question"`
	QuestionId uuid.UUID         `json:"question_id" gorm:"type:uuid;not null"`
	Answer     Intensity         `json:"level" sql:"type:intensity('none', 'low', 'medium', 'high');not null"`
}

type LanguageQuestion struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Question string    `json:"question" gorm:"type:varchar(255);not null"`
}

type LanguageAnswer struct {
	gorm.Model
	ID         uuid.UUID        `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProfileId  uuid.UUID        `json:"profile_id" gorm:"type:uuid;not null"`
	Profile    Profile          `json:"profile" gorm:"foreignKey:ProfileId"`
	Question   LanguageQuestion `json:"question"`
	QuestionId uuid.UUID        `json:"question_id" gorm:"type:uuid;not null"`
	Answer     Intensity        `json:"level" sql:"type:intensity('none', 'low', 'medium', 'high');not null"`
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

type Experience struct {
	gorm.Model
	ID                uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProfileId         uuid.UUID      `json:"profile_id" gorm:"type:uuid;not null"`
	Profile           Profile        `json:"profile" gorm:"foreignKey:ProfileId"`
	Company           string         `json:"company" gorm:"type:varchar(255);not null"`
	StartDate         datatypes.Date `json:"start_date"`
	EndDate           datatypes.Date `json:"end_date"`
	LastWage          *string        `json:"last_wage" gorm:"type:varchar(255)"`
	Role              *string        `json:"role" gorm:"type:varchar(255)"`
	Activities        *string        `json:"activities" gorm:"type:varchar(512)"`
	ReferenceContacts *string        `json:"reference_contacts" gorm:"type:varchar(512)"`
}

func (e Experience) GetStartDate() string {
	start_date, _ := e.StartDate.Value()
	return start_date.(time.Time).Format("02/01/2006")
}

func (e Experience) GetEndDate() string {
	end_date, _ := e.EndDate.Value()
	return end_date.(time.Time).Format("02/01/2006")
}
