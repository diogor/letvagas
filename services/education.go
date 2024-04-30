package services

import (
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"

	"github.com/google/uuid"
)

func CreateEducation(profile_id uuid.UUID, education *dto.CreateEducationRequest) error {
	profile := models.Profile{ID: profile_id}
	database.DB.First(&profile)

	new_ed := models.Education{
		Type:        education.EducationType,
		Institution: education.Institution,
		Graduation:  &education.Graduation,
		StartYear:   &education.StartYear,
		EndYear:     &education.EndYear,
		StartMonth:  &education.StartMonth,
		EndMonth:    &education.EndMonth,
		Description: &education.Description,
		Ongoing:     education.Ongoing,
	}

	result := database.DB.Model(&profile).
		Association("Educations").
		Append(&new_ed)

	return result
}

func ListEducations(profile_id uuid.UUID) []models.Education {
	educations := []models.Education{}

	profile := models.Profile{ID: profile_id}
	database.DB.First(&profile)

	database.DB.Model(&profile).Association("Educations").Find(&educations)

	return educations
}

func DeleteEducation(education_id uuid.UUID) error {
	education := models.Education{ID: education_id}

	result := database.DB.Delete(&education)

	return result.Error
}

func CreateAnswer(profile_id uuid.UUID, answer *dto.CreateAnswerRequest) error {
	profile := models.Profile{ID: profile_id}
	database.DB.First(&profile)

	new_answer := models.Answer{
		QuestionId: answer.QuestionId,
		Answer:     answer.Answer,
	}

	result := database.DB.Model(&profile).
		Association("Answers").
		Append(&new_answer)

	return result
}

func ListQuestions(question_type models.QuestionType) []dto.Question {
	questions := []models.Question{}
	result := []dto.Question{}

	database.DB.Find(&questions).Where("type = ?", question_type)

	for _, question := range questions {
		result = append(result, dto.Question{
			QuestionId: question.ID,
			Question:   question.Question,
			Options:    question.Options,
		})
	}

	return result
}
