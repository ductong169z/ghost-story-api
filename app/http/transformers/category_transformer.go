package transformers

import (
	"gfly/app/domain/models"
	"gfly/app/http/response"
)

// ToCategoryResponse transforms an Category model to an Category response
func ToCategoryResponse(category models.Category) response.Category {
	return response.Category{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description.String,
		IsActive:    category.IsActive,
		SortOrder:   category.SortOrder,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt.Time,
		DeletedAt:   category.DeletedAt.Time,
	}
}

// ToCategoryListResponse transforms a slice of Category models to a slice of Category responses
func ToCategoryListResponse(categories []models.Category) []response.Category {
	result := make([]response.Category, len(categories))
	for i, category := range categories {
		result[i] = ToCategoryResponse(category)
	}
	return result
}

// ToCategoryForGuestResponse transforms an Category model to an Category response
func ToCategoryForGuestResponse(category models.Category) response.Category {
	return response.Category{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description.String,
		IsActive:    category.IsActive,
		SortOrder:   category.SortOrder,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt.Time,
	}
}
