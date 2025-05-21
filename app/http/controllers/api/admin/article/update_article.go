package article

import (
	"gfly/app/constants"
	"gfly/app/dto"
	"gfly/app/http"
	"gfly/app/http/request"
	"gfly/app/http/response"
	"gfly/app/http/transformers"
	"gfly/app/services"

	"github.com/gflydev/core"
	"github.com/gflydev/core/log"
)

// ====================================================================
// ======================== Controller Creation =======================
// ====================================================================

type UpdateArticleApi struct {
	core.Api
}

func NewUpdateArticleApi() *UpdateArticleApi {
	return &UpdateArticleApi{}
}

// ====================================================================
// ======================== Request Validation ========================
// ====================================================================

func (h *UpdateArticleApi) Validate(c *core.Ctx) error {
	return http.ProcessUpdateRequest[request.UpdateArticle, dto.UpdateArticle](c)
}

// ====================================================================
// ========================= Request Handling =========================
// ====================================================================

// Handle function allows users to update an existing article
// @Description Function allows users to update an existing article
// @Summary Update an existing article
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param data body request.UpdateArticle true "UpdateArticle payload"
// @Success 200 {object} response.Article
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Unauthorized
// @Failure 404 {object} response.Error
// @Security ApiKeyAuth
// @Router /admin/articles/{id} [put]
func (h *UpdateArticleApi) Handle(c *core.Ctx) error {
	updateArticleDto := c.GetData(constants.Data).(dto.UpdateArticle)

	article, err := services.UpdateArticle(updateArticleDto)
	if err != nil {
		log.Error(err)

		// Check if it's a not found error
		if err.Error() == "Article not found" {
			return c.Error(response.Error{
				Code:    core.StatusNotFound,
				Message: err.Error(),
			}, core.StatusNotFound)
		}

		return c.Error(response.Error{
			Code:    core.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Transform to response data
	articleResponse := transformers.ToArticleResponse(*article)

	return c.Success(articleResponse)
}
