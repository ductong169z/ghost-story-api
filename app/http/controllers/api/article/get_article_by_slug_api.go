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

type GetArticleBySlugApi struct {
	core.Api
}

func NewGetArticleBySlugApi() *GetArticleBySlugApi {
	return &GetArticleBySlugApi{}
}

// ====================================================================
// ======================== Request Validation ========================
// ====================================================================

// Validate validates the slug parameter
func (h *GetArticleBySlugApi) Validate(c *core.Ctx) error {
	return http.ProcessPathParam(c, "slug")
}

// ====================================================================
// ========================= Request Handling =========================
// ====================================================================

// Handle function gets article by slug. If article doesn't exist, returns not found status.
// @Description Function gets article by slug. If article doesn't exist, returns not found status.
// @Summary Get article by slug
// @Tags Articles
// @Accept json
// @Produce json
// @Param slug path string true "Article Slug"
// @Success 200 {object} response.Article
// @Failure 401 {object} response.Unauthorized
// @Failure 404 {object} response.Error
// @Security ApiKeyAuth
// @Router /articles/slug/{slug} [get]
func (h *GetArticleBySlugApi) Handle(c *core.Ctx) error {
	slug := c.GetData(constants.Data).(string)

	article, err := services.GetArticleBySlug(slug)
	if err != nil {
		log.Error(err)

		return c.Error(response.Error{
			Code:    core.StatusNotFound,
			Message: "Article not found",
		}, core.StatusNotFound)
	}

	// Transform to response data
	articleResponse := transformers.ToArticleForGuestResponse(*article)

	return c.Success(articleResponse)
}
