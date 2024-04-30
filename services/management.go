package services

import (
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"
)

func CreateQuestion(question string, options []string, question_type models.QuestionType) error {

	new_question := models.Question{
		Question: question,
		Options:  options,
		Type:     question_type,
	}

	result := database.DB.Create(&new_question)

	return result.Error
}

func ListAllQuestions() []dto.Question {
	questions := []models.Question{}
	result := []dto.Question{}

	database.DB.Find(&questions)

	for _, question := range questions {
		result = append(result, dto.Question{
			QuestionId: question.ID,
			Question:   question.Question,
			Options:    question.Options,
		})
	}

	return result
}
