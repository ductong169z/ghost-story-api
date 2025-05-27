package services

import (
	"gfly/app/domain/models"
	"gfly/app/dto"

	mb "github.com/gflydev/db"
	qb "github.com/jivegroup/fluentsql"
)

// ====================================================================
// ========================= Main functions ===========================
// ====================================================================

// FindCategories retrieves a list of categories from the database based on the provided filter criteria.
//
// Parameters:
//   - filterDto (dto.Filter): The filter containing search criteria, order by field, page, and per-page details.
//
// Returns:
//
//	([]models.Category, int, error): A list of category models, the total number of categories, and any error encountered.
func FindCategories(filterDto dto.CategoryFilter) ([]models.Category, int, error) {
	// DB Model instance
	dbInstance := mb.Instance()
	// Error variable
	var err error

	// Define Category variable.
	var categories []models.Category
	var total int
	var offset = 0

	if filterDto.Page > 0 {
		offset = (filterDto.Page - 1) * filterDto.PerPage
	}

	builder := dbInstance.Select("*").
		Where(models.TableCategory+".deleted_at", qb.Null, nil).
		When(filterDto.Keyword != "", func(query qb.WhereBuilder) *qb.WhereBuilder {
			query.WhereGroup(func(queryGroup qb.WhereBuilder) *qb.WhereBuilder {
				queryGroup.Where(models.TableCategory+".name", qb.Like, "%"+filterDto.Keyword+"%").
					WhereOr(models.TableCategory+".slug", qb.Like, "%"+filterDto.Keyword+"%")

				return &queryGroup
			})

			return &query
		}).
		Limit(filterDto.PerPage, offset)

	// Query data
	total, err = builder.Find(&categories)

	// Return query result.
	return categories, total, err
}
