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
)

// ====================================================================
// ======================== Controller Creation =======================
// ====================================================================

type CreateArticleApi struct {
	core.Api
}

func NewCreateArticleApi() *CreateArticleApi {
	return &CreateArticleApi{}
}

// ====================================================================
// ======================== Request Validation ========================
// ====================================================================

func (h *CreateArticleApi) Validate(c *core.Ctx) error {
	return http.ProcessRequest[request.CreateArticle, dto.CreateArticle](c)
}

// ====================================================================
// ========================= Request Handling =========================
// ====================================================================

// Handle function allows users to create a new article
// @Description Function allows users to create a new article
// @Summary Create a new article
// @Tags Articles
// @Accept json
// @Produce json
// @Param data body request.CreateArticle true "CreateArticle payload"
// @Success 201 {object} response.Article
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Unauthorized
// @Security ApiKeyAuth
// @Router /articles [post]
func (h *CreateArticleApi) Handle(c *core.Ctx) error {
	createArticleDto := c.GetData(constants.Data).(dto.CreateArticle)

	article, err := services.CreateArticle(createArticleDto)
	if err != nil {
		return c.Error(response.Error{
			Code:    core.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Transform to response data
	articleResponse := transformers.ToArticleResponse(*article)

	return c.
		Status(core.StatusCreated).
		JSON(articleResponse)
}
