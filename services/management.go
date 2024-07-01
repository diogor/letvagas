package services

import (
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"

	"github.com/google/uuid"
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

func DeleteQuestion(question_id uuid.UUID) error {
	question := models.Question{ID: question_id}

	result := database.DB.Delete(&question)

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
			Type:       question.Type,
			Options:    question.Options,
		})
	}

	return result
}

func CreatePosition(position dto.CreatePositionRequest, user_id uuid.UUID) error {

	profile := GetProfile(user_id)

	new_position := models.Position{
		Company:      &position.Company,
		CompanyField: &position.CompanyField,
		Title:        &position.Title,
		Level:        position.Level,
		Education:    &position.Education,
		Type:         position.Type,
		Allocation:   position.Allocation,
		Wage:         &position.Wage,
		Location:     &position.Location,
		Description:  &position.Description,
		PCD:          position.PCD,
		PCDOnly:      position.PCDOnly,
		IsActive:     true,
		CreatedByID:  profile.ID,
		ExternalLink: &position.ExternalLink,
	}

	result := database.DB.Create(&new_position)

	return result.Error
}

func GetPositionByID(position_id uuid.UUID) *models.Position {
	position := models.Position{ID: position_id}
	database.DB.First(&position)
	return &position
}

func GetPositionSummaryByID(position_id uuid.UUID) *dto.ListPositionResponse {
	position := GetPositionByID(position_id)
	return &dto.ListPositionResponse{
		ID:         position.ID,
		Title:      *position.Title,
		Company:    *position.Company,
		Level:      position.GetLevel(),
		Education:  *position.Education,
		Type:       position.GetType(),
		Allocation: position.GetAllocation(),
		Wage:       *position.Wage,
		Location:   *position.Location,
		IsActive:   position.IsActive,
	}
}

func ListPositions(page, pageSize int, searchParams dto.PositionSearchParams) ([]dto.ListPositionResponse, int) {
	var total int64
	var positions []models.Position
	var response []dto.ListPositionResponse

	query := database.DB

	if searchParams.Query != "" {
		query = query.Where(
			"title LIKE ?", "%"+searchParams.Query+"%").Or(
			"company LIKE ?", "%"+searchParams.Query+"%").Or(
			"location LIKE ?", "%"+searchParams.Query+"%").Or(
			"description LIKE ?", "%"+searchParams.Query+"%")
	}

	query.Find(&positions)

	query.Model(&models.Position{}).Count(&total)
	query.Scopes(Paginate(page, pageSize)).Find(&positions)

	for i := 0; i < len(positions); i++ {
		response = append(response, dto.ListPositionResponse{
			ID:         positions[i].ID,
			Title:      *positions[i].Title,
			Company:    *positions[i].Company,
			Location:   *positions[i].Location,
			Level:      positions[i].GetLevel(),
			Type:       positions[i].GetType(),
			Allocation: positions[i].GetAllocation(),
			Wage:       *positions[i].Wage,
			PCD:        positions[i].PCD,
			IsActive:   positions[i].IsActive,
		})
	}

	return response, int(total)
}

func DeletePosition(position_id uuid.UUID) error {
	position := models.Position{ID: position_id}
	result := database.DB.Delete(&position)

	return result.Error
}

func SetPositionStatus(position_id uuid.UUID, active bool) error {
	position := models.Position{ID: position_id}
	result := database.DB.Model(&position).Update("is_active", active)

	return result.Error
}

func FindApplicationByProfileAndPosition(profile_id, position_id uuid.UUID) *models.Application {
	application := models.Application{}

	database.DB.First(&application, "profile_id = ? AND position_id = ?", profile_id, position_id)
	return &application
}

func FindApplicationsByProfile(profile_id uuid.UUID) []dto.ApllicationListResponse {
	application := []models.Application{}
	response := []dto.ApllicationListResponse{}
	database.DB.Where("profile_id = ?", profile_id).Find(&application)

	for _, app := range application {
		position := GetPositionByID(app.PositionId)

		position_dto := dto.PositionResponse{
			ID: app.PositionId,
		}

		position_dto.Title = *position.Title
		position_dto.Company = *position.Company
		position_dto.Location = *position.Location
		position_dto.Level = position.GetLevel()
		position_dto.Type = position.GetType()
		position_dto.Allocation = position.GetAllocation()
		position_dto.Wage = *position.Wage
		position_dto.PCD = position.PCD

		response = append(response, dto.ApllicationListResponse{
			ID:        app.ID,
			Position:  position_dto,
			ProfileID: app.ProfileId,
		})
	}

	return response

}

func FindUsersForPosition(position_id uuid.UUID) []models.User {
	applications := []models.Application{}

	database.DB.Where("position_id = ?", position_id).Find(&applications)
	users := []models.User{}

	for _, app := range applications {
		user := models.User{}
		database.DB.First(&user, "profile_id = ?", app.ProfileId)
		users = append(users, user)
	}

	return users
}

func ListAllUsers(page, pageSize int) ([]models.User, int) {
	users := []models.User{}
	var total int64

	database.DB.Model(&models.User{}).Count(&total)
	database.DB.Scopes(Paginate(page, pageSize)).Find(&users)

	return users, int(total)

}

func ChangeUserRole(user_id uuid.UUID, role models.UserRole) {
	database.DB.Model(&models.User{}).Where("id = ?", user_id).Update("role", role)
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

func SearchProfiles(page int, pageSize int, searchParams dto.SearchParams) ([]models.User, int) {
	var profiles []models.User
	var total int64

	query := database.DB

	if searchParams.Query != "" {
		query = query.Joins(
			"left join profiles p on users.profile_id = p.id").Joins(
			"left join experiences e on p.id = e.profile_id").Joins(
			"left join educations ed on e.profile_id = ed.profile_id").Joins(
			"left join courses c on ed.profile_id = c.profile_id").Where(
			"users.name LIKE ?", "%"+searchParams.Query+"%").Or(
			"p.goal LIKE ?", "%"+searchParams.Query+"%").Or(
			"users.social_name LIKE ?", "%"+searchParams.Query+"%").Or(
			"e.activities LIKE ?", "%"+searchParams.Query+"%").Or(
			"ed.institution LIKE ?", "%"+searchParams.Query+"%").Or(
			"ed.graduation LIKE ?", "%"+searchParams.Query+"%").Or(
			"ed.description LIKE ?", "%"+searchParams.Query+"%").Or(
			"c.name LIKE ?", "%"+searchParams.Query+"%").Or(
			"c.description LIKE ?", "%"+searchParams.Query+"%")
	}

	if searchParams.City != "" {
		query = query.Where("city = ?", searchParams.City)
	}

	if searchParams.State != "" {
		query = query.Where("state = ?", searchParams.State)
	}

	if searchParams.EducationType != "" {
		query = query.Where("education_type = ?", searchParams.EducationType)
	}

	if searchParams.Neighborhood != "" {
		query = query.Where("neighborhood = ?", searchParams.Neighborhood)
	}

	query.Find(&profiles)

	query.Model(&models.User{}).Count(&total)
	query.Scopes(Paginate(page, pageSize)).Find(&profiles)
	return profiles, int(total)
}
