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

type UpdateArticleStatusApi struct {
	core.Api
}

func NewUpdateArticleStatusApi() *UpdateArticleStatusApi {
	return &UpdateArticleStatusApi{}
}

// ====================================================================
// ======================== Request Validation ========================
// ====================================================================

func (h *UpdateArticleStatusApi) Validate(c *core.Ctx) error {
	return http.ProcessUpdateRequest[request.UpdateArticleStatus, dto.UpdateArticleStatus](c)
}

// ====================================================================
// ========================= Request Handling =========================
// ====================================================================

// Handle function updates an article's status
// @Description Function updates an article's status (draft, published, archived)
// @Summary Update article status
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param request body request.UpdateArticleStatus true "Update article status data"
// @Success 200 {object} response.Article
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Unauthorized
// @Failure 404 {object} response.Error
// @Security ApiKeyAuth
// @Router /admin/articles/{id}/status [put]
func (h *UpdateArticleStatusApi) Handle(c *core.Ctx) error {
	updateArticleStatusDto := c.GetData(constants.Data).(dto.UpdateArticleStatus)

	article, err := services.UpdateArticleStatus(updateArticleStatusDto)
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
