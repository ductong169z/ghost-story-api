package article

import (
	"gfly/app/constants"
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

type GetArticleByIdApi struct {
	core.Api
}

func NewGetArticleByIdApi() *GetArticleByIdApi {
	return &GetArticleByIdApi{}
}

// ====================================================================
// ======================== Request Validation ========================
// ====================================================================

func (h *GetArticleByIdApi) Validate(c *core.Ctx) error {
	return http.ProcessPathID(c)
}

// ====================================================================
// ========================= Request Handling =========================
// ====================================================================

// Handle function gets article by given id. If article doesn't exist, returns not found status.
// @Description Function gets article by given id. If article doesn't exist, returns not found status.
// @Summary Get article by given id
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param count_view query bool false "Whether to count this view"
// @Success 200 {object} response.Article
// @Failure 401 {object} response.Unauthorized
// @Failure 404 {object} response.Error
// @Security ApiKeyAuth
// @Router /articles/{id} [get]
func (h *GetArticleByIdApi) Handle(c *core.Ctx) error {
	articleID := c.GetData(constants.Data).(int)

	article, err := services.GetArticleByID(articleID)
	if err != nil {
		log.Error(err)

		return c.Error(response.Error{
			Code:    core.StatusNotFound,
			Message: "Article not found",
		}, core.StatusNotFound)
	}

	// Check if we should count this view
	// This allows admins to preview articles without incrementing the view count
	// unless explicitly requested
	countView := c.QueryStr("count_view") == "true"
	if countView {
		if err := services.IncrementArticleViewCount(articleID); err != nil {
			log.Warnf("Failed to increment view count for article %d: %v", articleID, err)
			// Continue even if view count update fails
		}
	}

	// Transform to response data
	articleResponse := transformers.ToArticleResponse(*article)

	return c.Success(articleResponse)
}
