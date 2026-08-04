package models

type CreateProjectRequest struct {
	ClientName string `json:"client_name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Company    string `json:"company"`

	ProjectName string `json:"project_name" binding:"required"`

	Service string `json:"service" binding:"required"`

	Budget string `json:"budget"`

	Description string `json:"description"`
}
