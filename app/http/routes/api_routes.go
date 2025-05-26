package routes

import (
	"fmt"
	"gfly/app/domain/models/types"
	"gfly/app/http/controllers/api"
	adminArticle "gfly/app/http/controllers/api/admin/article"
	"gfly/app/http/controllers/api/article"
	"gfly/app/http/controllers/api/user"
	"gfly/app/http/middleware"
	authRoute "gfly/app/modules/auth/routes"

	"github.com/gflydev/core"
	"github.com/gflydev/core/utils"
)

// ApiRoutes func for describe a group of API routes.
func ApiRoutes(r core.IFly) {
	prefixAPI := fmt.Sprintf(
		"/%s/%s",
		utils.Getenv("API_PREFIX", "api"),
		utils.Getenv("API_VERSION", "v1"),
	)

	// API Routers
	r.Group(prefixAPI, func(apiRouter *core.Group) {
		// curl -v -X GET http://localhost:7789/api/v1/info | jq
		apiRouter.GET("/info", api.NewDefaultApi())

		// Public routes (no authentication required)
		apiRouter.Group("articles", func(publicRouter *core.Group) {
			publicRouter.GET("", article.NewListArticlesApi())
			publicRouter.GET("/{slug:[a-z0-9-]+}", article.NewGetArticleBySlugApi())
		})

		/* ============================ Auth Group ============================ */
		authRoute.RegisterApi(apiRouter)

		// Protected routes (require authentication)

		apiRouter.Group("/admin", func(adminRouter *core.Group) {
			adminRouter.Use(middleware.CheckRolesMiddleware(
				[]types.Role{types.RoleAdmin},
				prefixAPI+"/admin",
			))

			adminRouter.Group("/users", func(userRouter *core.Group) {
				// Allow admin permission to access `/users/*` API
				userRouter.Use(middleware.CheckRolesMiddleware(
					[]types.Role{types.RoleAdmin},
					prefixAPI+"/users/profile",
				))

				userRouter.GET("", user.NewListUsersApi())
				userRouter.POST("", user.NewCreateUserApi())
				userRouter.PUT("/{id}/status", r.Apply(middleware.PreventUpdateYourSelf)(user.NewUpdateUserStatusApi()))
				userRouter.PUT("/{id}", r.Apply(middleware.PreventUpdateYourSelf)(user.NewUpdateUserApi()))
				userRouter.DELETE("/{id}", r.Apply(middleware.PreventUpdateYourSelf)(user.NewDeleteUserApi()))
				userRouter.GET("/{id}", user.NewGetUserByIdApi())
				userRouter.GET("/profile", user.NewGetUserProfileApi())
			})

			/* ============================ Admin Article Group ============================ */
			adminRouter.Group("/articles", func(articleRouter *core.Group) {
				// Allow admin permission to access `/articles/*` API
				articleRouter.Use(middleware.CheckRolesMiddleware(
					[]types.Role{types.RoleAdmin},
					prefixAPI+"/articles",
				))

				articleRouter.POST("", adminArticle.NewCreateArticleApi())
				articleRouter.GET("", adminArticle.NewListArticlesApi())
				articleRouter.GET("/{id}", adminArticle.NewGetArticleByIdApi())
				articleRouter.PUT("/{id}", adminArticle.NewUpdateArticleApi())
				articleRouter.PUT("/{id}/status", adminArticle.NewUpdateArticleStatusApi())
				articleRouter.DELETE("/{id}", adminArticle.NewDeleteArticleApi())
			})
		})
	})
}
