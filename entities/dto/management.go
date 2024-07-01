package dto

import (
	"letvagas/entities/models"

	"github.com/google/uuid"
)

type CreatePositionRequest struct {
	Company      string              `json:"company" binding:"required"`
	CompanyField string              `json:"company_field"`
	Title        string              `json:"title" binding:"required"`
	Level        models.Level        `json:"level" binding:"required"`
	Education    string              `json:"education"`
	Type         models.PositionType `json:"type" binding:"required"`
	Allocation   models.Allocation   `json:"allocation" binding:"required"`
	Wage         string              `json:"wage"`
	Location     string              `json:"location"`
	Description  string              `json:"description"`
	PCD          bool                `json:"pcd"`
	PCDOnly      bool                `json:"pcd_only"`
	ExternalLink string              `json:"external_link"`
}

type PositionResponse struct {
	CreatePositionRequest
	Level      string    `json:"level"`
	Type       string    `json:"type"`
	Allocation string    `json:"allocation"`
	ID         uuid.UUID `json:"id"`
}

type ListPositionResponse struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Company    string    `json:"company"`
	Location   string    `json:"location"`
	Level      string    `json:"level"`
	Type       string    `json:"type"`
	Education  string    `json:"education"`
	Allocation string    `json:"allocation"`
	Wage       string    `json:"wage"`
	PCD        bool      `json:"pcd"`
	PCDOnly    bool      `json:"pcd_only"`
	IsActive   bool      `json:"is_active"`
}

type ApllicationListResponse struct {
	ID        uuid.UUID        `json:"id"`
	ProfileID uuid.UUID        `json:"profile_id"`
	Position  PositionResponse `json:"position"`
}

type PositionSearchParams struct {
	Query string
}
