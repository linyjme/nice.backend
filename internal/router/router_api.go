package router

import (
	"niceBackend/internal/api/controller/article_handler"
	"niceBackend/internal/api/controller/auth_handler"
	"niceBackend/internal/api/controller/category_handler"
	"niceBackend/internal/api/controller/expansion_handler"
	"niceBackend/internal/api/controller/tag_handler"
)

func setApiRouter(r *resource) {

	BaseRouter := r.mux.Group("")
	tagHandler := tag_handler.New()
	categoryHandler := category_handler.New()
	articleHandler := article_handler.New()
	authHandler := auth_handler.New()
	expansionHandler := expansion_handler.New()
	{
		BaseRouter.GET("/tag", tagHandler.List())           // 创建Api
		BaseRouter.GET("/category", categoryHandler.List()) // 创建Api
		BaseRouter.GET("/article", articleHandler.List())   // 创建Api
		BaseRouter.POST("/auth/login", authHandler.Login())  // 创建Api
		BaseRouter.GET("/auth/admin", authHandler.List())   // 创建Api
		BaseRouter.GET("/expansion/statistic", expansionHandler.List())   // 创建Api
		BaseRouter.POST("/expansion/statistic", expansionHandler.List())   // 创建Api
	}

}
