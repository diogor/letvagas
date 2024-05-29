package models

import (
	"database/sql/driver"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EducationType string
type QuestionType string

const (
	FUNDAMENTAL EducationType = "fundamental"
	MEDIO       EducationType = "medio"
	SUPERIOR    EducationType = "superior"
	POS         EducationType = "pos"
	MBA         EducationType = "mba"
	OUTROS      EducationType = "outros"
	COMPUTING   QuestionType  = "computing"
	LANGUAGE    QuestionType  = "language"
)

func (ct *EducationType) Scan(value interface{}) error {
	*ct = EducationType(value.(string))
	return nil
}

func (ct EducationType) Value() (driver.Value, error) {
	return string(ct), nil
}

func (ct *QuestionType) Scan(value interface{}) error {
	*ct = QuestionType(value.(string))
	return nil
}

func (ct QuestionType) Value() (driver.Value, error) {
	return string(ct), nil
}

type QuestionOptions []string

func (o *QuestionOptions) Scan(value interface{}) error {
	if value == nil {
		*o = []string{}
		return nil
	}
	*o = strings.Split(value.(string), ",")
	return nil
}

func (o QuestionOptions) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}
	return strings.Join(o, ","), nil
}

type Profile struct {
	gorm.Model
	ID          uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Answers     []Answer     `json:"answers" gorm:"foreignKey:ProfileId"`
	Educations  []Education  `json:"educations" gorm:"foreignKey:ProfileId"`
	Courses     []Course     `json:"courses" gorm:"foreignKey:ProfileId"`
	Experiences []Experience `json:"experiences" gorm:"foreignKey:ProfileId"`
	Goal        *string      `json:"goal" gorm:"type:varchar(255)"`
}

type Answer struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProfileId  uuid.UUID `json:"profile_id" gorm:"type:uuid;not null;uniqueIndex:answer_question_id"`
	Profile    Profile   `json:"profile" gorm:"foreignKey:ProfileId"`
	Question   Question  `json:"question"`
	QuestionId uuid.UUID `json:"question_id" gorm:"type:uuid;not null;uniqueIndex:answer_question_id"`
	Answer     string    `json:"answer"`
}

type Question struct {
	gorm.Model
	ID       uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Question string          `json:"question" gorm:"type:varchar(255);not null"`
	Type     QuestionType    `json:"type" sql:"type:question_type('computing', 'language');not null"`
	Options  QuestionOptions `json:"options" gorm:"type:varchar(255)"`
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
