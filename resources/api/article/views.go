package article


import (
	"github.com/gin-gonic/gin"
	"niceBackend/transform/response"
)

func GetArticles(c *gin.Context){
	response.OkWithMessage("ok", c)
	return
}

