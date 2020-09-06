package routers

import (
	"github.com/gin-gonic/gin"

	"blog/packages/settings"
	"blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tag", v1.GetTags)
		apiV1.POST("/tag", v1.AddTag)
		apiV1.PUT("/tag/:id", v1.EditTag)
		apiV1.DELETE("/tag/:id", v1.DeleteTag)
		apiV1.GET("/post", v1.GetPosts)
		apiV1.GET("/post/:id", v1.GetPost)
		apiV1.POST("/post", v1.AddPost)
		apiV1.PUT("/post/:id", v1.EditPost)
		apiV1.DELETE("/post/:id", v1.DeletePost)
	}

	return r
}
