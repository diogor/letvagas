package dto

import "letvagas/entities/models"

type CreateUserRequest struct {
	Name       string          `json:"name" binding:"required"`
	SocialName string          `json:"social_name"`
	Password   string          `json:"password" binding:"required"`
	Email      string          `json:"email" binding:"required"`
	Cpf        string          `json:"cpf" binding:"required"`
	Role       models.UserRole `json:"role" binding:"required"`
	BirthDate  string          `json:"birth_date" binding:"required"`
	Phone      string          `json:"phone" binding:"required"`
	Linkedin   string          `json:"linkedin"`
	City       string          `json:"city"`
	State      string          `json:"state"`
	Address    string          `json:"address"`
	Cep        string          `json:"cep"`
}
