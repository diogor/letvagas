package services

import (
	"errors"
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(user *dto.CreateUserRequest) (uuid.UUID, error) {
	password, err := hashPassword(user.Password)

	if err != nil {
		return uuid.Nil, err
	}

	new_user := &models.User{
		Name:         user.Name,
		SocialName:   &user.SocialName,
		Password:     password,
		Email:        user.Email,
		Role:         models.APPLICANT,
		Cpf:          user.Cpf,
		BirthDate:    user.BirthDate,
		City:         &user.City,
		State:        &user.State,
		Cep:          &user.Cep,
		Street:       &user.Street,
		Neighborhood: &user.Neighborhood,
		Number:       &user.Number,
		Complement:   &user.Complement,
		Phone1:       user.Phone1,
		AreaCode1:    user.AreaCode1,
		Phone2:       &user.Phone2,
		AreaCode2:    &user.AreaCode2,
		Phone3:       &user.Phone3,
		AreaCode3:    &user.AreaCode3,
		Linkedin:     &user.Linkedin,
		PCD:          user.PCD,
		PCDInfo:      &user.PCDInfo,
		Profile:      models.Profile{ID: uuid.New()},
	}

	result := database.DB.Create(new_user)

	return new_user.ID, result.Error
}

func GetProfile(user_id uuid.UUID) *models.Profile {
	user := models.User{ID: user_id}
	database.DB.First(&user).Association("Profile")
	profile := models.Profile{ID: user.ProfileID}
	database.DB.First(&profile)

	return &profile
}

func UpdateProfileGoal(profile_id uuid.UUID, goal string) {
	database.DB.Model(&models.Profile{}).Where("id = ?", profile_id).Update("goal", goal)
}

func GetUser(id string) *models.User {
	user := models.User{ID: uuid.MustParse(id)}
	database.DB.First(&user)
	return &user
}

func GetUserByEmail(email string) *models.User {
	user := models.User{Email: email}
	database.DB.First(&user, "email = ?", email)

	if user.ID == uuid.Nil {
		return nil
	}

	return &user
}

func CreateApplication(profile_id uuid.UUID, position_id uuid.UUID) (uid uuid.UUID, err error) {
	position := models.Position{ID: position_id}
	database.DB.First(&position)

	if position.IsActive == false {
		return uuid.Nil, errors.New("Não aceita candidaturas no momento.")
	}

	new_application := &models.Application{
		Profile:  models.Profile{ID: profile_id},
		Position: models.Position{ID: position_id},
	}
	result := database.DB.Create(new_application)

	return new_application.ID, result.Error
}

func SaveProfileFile(profile_id uuid.UUID, name string, filePath *string, url *string) {
	new_saved_file := &models.SavedFile{
		ProfileId: profile_id,
		Name:      name,
		Path:      filePath,
		Url:       url,
	}

	database.DB.Create(new_saved_file)
}

func ListProfileFiles(profile_id uuid.UUID) []dto.ProfileFile {
	saved_files := []models.SavedFile{}

	database.DB.Model(&models.SavedFile{}).Where("profile_id = ?", profile_id).Find(&saved_files)

	response := []dto.ProfileFile{}

	for _, file := range saved_files {
		link := "/media/" + *file.Path

		if file.Url != nil {
			link = *file.Url
		}

		response = append(response, dto.ProfileFile{
			Name: file.Name,
			Link: link,
		})
	}

	return response
}
