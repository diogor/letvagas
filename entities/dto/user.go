package dto

import (
	"github.com/google/uuid"
	"letvagas/entities/models"
)

type CreateUserRequest struct {
	Name         string          `json:"name" binding:"required"`
	SocialName   string          `json:"social_name"`
	Password     string          `json:"password" binding:"required"`
	Email        string          `json:"email" binding:"required"`
	Cpf          string          `json:"cpf" binding:"required"`
	Role         models.UserRole `json:"role" binding:"required"`
	BirthDate    string          `json:"birth_date" binding:"required"`
	Phone1       string          `json:"phone1" binding:"required"`
	AreaCode1    string          `json:"area_code1" binding:"required"`
	Phone2       string          `json:"phone2"`
	AreaCode2    string          `json:"area_code2"`
	Phone3       string          `json:"phone3"`
	AreaCode3    string          `json:"area_code3"`
	Linkedin     string          `json:"linkedin"`
	City         string          `json:"city"`
	Neighborhood string          `json:"neighborhood"`
	State        string          `json:"state"`
	Cep          string          `json:"cep"`
	Street       string          `json:"street"`
	Number       string          `json:"number"`
	Complement   string          `json:"complement"`
	PCD          bool            `json:"pcd"`
	PCDInfo      string          `json:"pcd_info"`
}

type ProfileFile struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Link string    `json:"link"`
}
