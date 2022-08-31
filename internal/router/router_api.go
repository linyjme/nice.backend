package router

import (
	"niceBackend/internal/api/controller/announcement_handler"
	"niceBackend/internal/api/controller/article_handler"
	"niceBackend/internal/api/controller/auth_handler"
	"niceBackend/internal/api/controller/category_handler"
	"niceBackend/internal/api/controller/comment_handler"
	"niceBackend/internal/api/controller/expansion_handler"
	"niceBackend/internal/api/controller/option_handler"
	"niceBackend/internal/api/controller/tag_handler"
)

func setApiRouter(r *resource) {

	BaseRouter := r.mux.Group("")
	tagHandler := tag_handler.New()
	categoryHandler := category_handler.New()
	announcementHandler := announcement_handler.New()
	commentHandler := comment_handler.New()
	articleHandler := article_handler.New()
	authHandler := auth_handler.New()
	expansionHandler := expansion_handler.New()
	optionHandler := option_handler.New()
	{
		BaseRouter.GET("/tag", tagHandler.List())                        //
		BaseRouter.GET("/category", categoryHandler.List())              //
		BaseRouter.GET("/announcement", announcementHandler.List())      //
		BaseRouter.GET("/comment", commentHandler.List())                //
		BaseRouter.GET("/option", optionHandler.List())                  //
		BaseRouter.GET("/article", articleHandler.List())                //
		BaseRouter.POST("/auth/login", authHandler.Login())              //
		BaseRouter.POST("/auth/check", authHandler.TokenCheck())         //
		BaseRouter.GET("/auth/admin", authHandler.List())                //
		BaseRouter.GET("/expansion/statistic", expansionHandler.List())  //
		BaseRouter.POST("/expansion/statistic", expansionHandler.List()) //
	}

}
