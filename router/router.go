package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		v1.GET("", func(context *gin.Context) {

		})
	}

	v2 := api.Group("/v2")
	{
		v2.GET("", func(context *gin.Context) {

		})
	}

	return r
}
