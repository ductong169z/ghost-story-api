package api

import (
	"gfly/app/http/response"
	"gfly/app/modules/auth"
	"gfly/app/modules/auth/services"
	"github.com/gflydev/core"
)

// ====================================================================
// ======================== Controller Creation =======================
// ====================================================================

// NewSignOutApi As a constructor to create new API.
func NewSignOutApi(authType auth.Type) *SignOutApi {
	return &SignOutApi{
		Type: authType,
	}
}

type SignOutApi struct {
	Type auth.Type
	core.Api
}

// ====================================================================
// ========================= Request Handling =========================
// ====================================================================

// Handle method to invalidate users access token by adding them to a blacklist in Redis
// and delete refresh token from the Redis
// @Description De-authorize user and delete refresh token from Redis.
// @Summary de-authorize user and delete refresh token from Redis
// @Tags Auth
// @Accept json
// @Produce json
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Unauthorized
// @Success 204
// @Security ApiKeyAuth
// @Router /auth/signout [delete]
func (h *SignOutApi) Handle(c *core.Ctx) error {
	if h.Type == auth.TypeAPI {
		jwtToken := services.ExtractToken(c)

		if err := services.SignOut(jwtToken); err != nil {
			return c.Error(response.Error{
				Code:    core.StatusBadRequest,
				Message: err.Error(),
			})
		}
	} else {
		c.SetSession(auth.SessionUsername, "")
	}

	return c.NoContent()
}
