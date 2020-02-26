package route

import (
	"github.com/gin-gonic/gin"
	"github.com/vegchic/fullbottle/api/handler"
	"github.com/vegchic/fullbottle/api/middleware"
	"net/http"
)

func GinRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.WithContext())
	router.Use(middleware.ApiLogWrapper())

	registerV1Routes(router)

	return router
}

func registerV1Routes(g *gin.Engine) {
	api := g.Group("/v1")
	{
		api.POST("/register", handler.RegisterUser)
		api.POST("/login", handler.UserLogin)

		api.GET("/users/avatar", handler.GetUserAvatar) // no permission asked

		api.Use(middleware.LoginRequired())

		// user
		api.GET("/users/profile", handler.GetUser)
		api.PUT("/users/profile", handler.UpdateUser)
		api.POST("/users/avatar", handler.UploadAvatar)

		// space
		api.Use(middleware.FolderAccessCheck())

		api.GET("/space", handler.GetSpaceMeta)
		api.POST("/space/folders", handler.CreateFolder)
		api.GET("/space/folders/:folder_id", handler.GetFolder)
		api.PUT("/space/folders/:folder_id", handler.UpdateFolder)
		api.DELETE("/space/folders/:folder_id", handler.RemoveFolder)
	}

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid api",
		})
	})

}
