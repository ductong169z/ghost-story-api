package dto

// CreateCategory struct to describe the request body to create a new category.
// @Description Request payload for creating a new category.
// @Tags Categories

type CreateCategory struct {
	Name        string `json:"name" example:"Technology" validate:"required,max=255" doc:"Category name (required, max length 255)"`
	Slug        string `json:"slug" example:"technology" validate:"required,max=255" doc:"URL-friendly slug (required, max length 255)"`
	Description string `json:"description" example:"Latest news and updates in the technology industry" validate:"omitempty" doc:"Description of the category (optional)"`
	ParentID    int    `json:"parent_id" example:"1" validate:"omitempty,gte=1" doc:"ID of the parent category (optional, greater than or equal to 1)"`
}

// UpdateCategory struct to describe the request body to update an existing category.
// @Description Request payload for updating an existing category.
// @Tags Categories
type UpdateCategory struct {
	ID          int    `json:"-" validate:"omitempty,gte=1" doc:"Category ID (greater than or equal to 1)"`
	Name        string `json:"name" example:"Technology" validate:"omitempty,max=255" doc:"Category name (optional, max length 255)"`
	Slug        string `json:"slug" example:"technology" validate:"omitempty,max=255" doc:"URL-friendly slug (optional, max length 255)"`
	Description string `json:"description" example:"Latest news and updates in the technology industry" validate:"omitempty" doc:"Description of the category (optional)"`
	ParentID    int    `json:"parent_id" example:"1" validate:"omitempty,gte=1" doc:"ID of the parent category (optional, greater than or equal to 1)"`
}
