package category

import (
	"gfly/app/dto"
	"gfly/app/http"
	"gfly/app/http/response"
	"gfly/app/http/transformers"
	"gfly/app/services"

	"github.com/gflydev/core"
	"github.com/gflydev/core/log"
)

// ====================================================================
// ======================== Controller Creation =======================
// ====================================================================

type ListArticlesApi struct {
	core.Api
}

func NewListArticlesApi() *ListArticlesApi {
	return &ListArticlesApi{}
}

// ====================================================================
// ======================== Request Validation ========================
// ====================================================================

// Validate validates the query parameters for category listing
func (h *ListArticlesApi) Validate(c *core.Ctx) error {
	// Create filter from query parameters and validate it
	filter := dto.CategoryFilter{}

	// Get base filter parameters (page, per_page, keyword, order_by)
	filter.Page, _ = c.QueryInt("page")
	filter.PerPage, _ = c.QueryInt("per_page")
	filter.Keyword = c.QueryStr("keyword")
	filter.OrderBy = c.QueryStr("order_by")

	// Set default values if not provided
	if filter.Page < 1 {
		filter.Page = 1
	}

	if filter.PerPage < 1 {
		filter.PerPage = 10
	}

	// Validate filter
	if errData := http.Validate(filter); errData != nil {
		return c.Error(errData)
	}

	// Store filter in context
	c.SetData("filter", filter)

	return nil
}

// ====================================================================
// ========================= Request Handling =========================
// ====================================================================

// Handle function gets a list of articles based on filter criteria
// @Description Returns a paginated list of articles that can be filtered and sorted
// @Summary List articles with pagination and filtering
// @Tags Articles
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10)"
// @Param keyword query string false "Search keyword in title, slug, excerpt, and content"
// @Param order_by query string false "Field to order by (prefix with '-' for descending, e.g. '-created_at')"
// @Param status query string false "Filter by article status (draft, published, archived)"
// @Success 200 {object} response.PaginatedResponse
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Unauthorized
// @Security ApiKeyAuth
// @Router /articles [get]
func (h *ListArticlesApi) Handle(c *core.Ctx) error {
	// Get filter from context
	filter := c.GetData("filter").(dto.CategoryFilter)

	// Get articles from service
	categories, total, err := services.FindCategories(filter)
	if err != nil {
		log.Errorf("Error while fetching articles: %v", err)

		return c.Error(response.Error{
			Code:    core.StatusInternalServerError,
			Message: "Error occurred while fetching articles",
		}, core.StatusInternalServerError)
	}

	// Transform articles to response format
	categoriesResponse := transformers.ToCategoryListResponse(categories)

	// Create paginated response
	return c.Success(response.PaginatedResponse{
		Data:       categoriesResponse,
		Pagination: createPagination(filter.Page, filter.PerPage, total),
	})
}

// Helper function to create pagination metadata
func createPagination(page, perPage, total int) response.Pagination {
	totalPages := (total + perPage - 1) / perPage // Ceiling division

	return response.Pagination{
		CurrentPage: page,
		PerPage:     perPage,
		Total:       total,
		TotalPages:  totalPages,
		HasMore:     page < totalPages,
	}
}
