package dto

import (
	"letvagas/entities/models"

	"github.com/google/uuid"
)

type CreateEducationRequest struct {
	EducationType models.EducationType `json:"education_type" binding:"required"`
	Institution   string               `json:"institution" binding:"required"`
	Graduation    string               `json:"graduation"`
	StartYear     string               `json:"start_year"`
	EndYear       string               `json:"end_year"`
	StartMonth    string               `json:"start_month"`
	EndMonth      string               `json:"end_month"`
	Description   string               `json:"description"`
	Ongoing       bool                 `json:"ongoing"`
}

type CreateCourseRequest struct {
	Name        string `json:"name" binding:"required"`
	Year        string
	Month       string
	Description string
	Ongoing     bool
}

type CreateExperienceRequest struct {
	Company    string `json:"company" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date"`
	LastWage   string `json:"last_wage"`
	Role       string `json:"role"`
	Activities string `json:"activities"`
	Reference  string `json:"reference_contacts"`
}

type CreateAnswerRequest struct {
	QuestionId uuid.UUID `json:"question_id" binding:"required"`
	Answer     string    `json:"answer" binding:"required"`
}

type Options struct {
	Option  string `json:"option" binding:"required"`
	Checked bool   `json:"checked" binding:"required"`
}

type Question struct {
	QuestionId uuid.UUID `json:"question_id" binding:"required"`
	Question   string    `json:"question" binding:"required"`
	Options    []Options `json:"options" binding:"required"`
}

type QuestionList struct {
	QuestionId uuid.UUID `json:"question_id" binding:"required"`
	Question   string    `json:"question" binding:"required"`
	Options    []string  `json:"options" binding:"required"`
}
