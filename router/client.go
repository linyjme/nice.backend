package router

import (
	"github.com/gin-gonic/gin"
	"niceBackend/resources/api/article"
	"niceBackend/resources/api/category"
	"niceBackend/resources/api/option"
	"niceBackend/resources/api/tags"
)

func BaseClientRouter(Router *gin.RouterGroup) () {

	BaseRouter := Router.Group("api")
	{
		BaseRouter.GET("tag", tags.GetTags) // 创建Api
		BaseRouter.GET("category", category.GetCategory) // 创建Api
		BaseRouter.GET("article", article.GetArticles) // 创建Api
		BaseRouter.GET("option", option.GetOption) // 创建Api
	}
}
