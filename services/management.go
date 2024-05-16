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

func ListAllQuestions() []dto.QuestionList {
	questions := []models.Question{}
	result := []dto.QuestionList{}

	database.DB.Find(&questions)

	for _, question := range questions {
		result = append(result, dto.QuestionList{
			QuestionId: question.ID,
			Question:   question.Question,
			Options:    question.Options,
		})
	}

	return result
}

func ListUserStates() []string {
	states := []string{}

	database.DB.Model(&models.User{}).Distinct().Pluck("state", &states)

	return states
}

func ListUserCities(state string) []string {
	cities := []string{}

	database.DB.Model(&models.User{}).Where("state = ?", state).Distinct().Pluck("city", &cities)

	return cities
}

func SearchProfiles(q string, city string, state string) []models.User {
	var profiles []models.User

	query := database.DB

	if city != "" {
		query = query.Where("city = ?", city)
	}

	if state != "" {
		query = query.Where("state = ?", state)
	}

	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}

	query.Find(&profiles)

	return profiles
}
