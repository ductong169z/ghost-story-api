package article

import (
	"gfly/app/constants"
	"gfly/app/http"
	"gfly/app/http/response"
	"gfly/app/services"

	"github.com/gflydev/core"
	"github.com/gflydev/core/log"
)

// ====================================================================
// ======================== Controller Creation =======================
// ====================================================================

type DeleteArticleApi struct {
	core.Api
}

func NewDeleteArticleApi() *DeleteArticleApi {
	return &DeleteArticleApi{}
}

// ====================================================================
// ======================== Request Validation ========================
// ====================================================================

func (h *DeleteArticleApi) Validate(c *core.Ctx) error {
	return http.ProcessPathID(c)
}

// ====================================================================
// ========================= Request Handling =========================
// ====================================================================

// Handle function allows users to delete an article
// @Description Function allows users to delete an article
// @Summary Delete an article
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 204
// @Failure 401 {object} response.Unauthorized
// @Failure 404 {object} response.Error
// @Security ApiKeyAuth
// @Router /admin/articles/{id} [delete]
func (h *DeleteArticleApi) Handle(c *core.Ctx) error {
	articleID := c.GetData(constants.Data).(int)

	err := services.DeleteArticleByID(articleID)
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
			Code:    core.StatusInternalServerError,
			Message: "Error occurred while deleting the article",
		}, core.StatusInternalServerError)
	}

	return c.NoContent()
}
