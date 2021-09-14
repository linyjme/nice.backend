package article


import (
	"github.com/gin-gonic/gin"
	"niceBackend/common/transform/response"
)

func GetArticles(c *gin.Context){
	response.OkWithMessage("ok", c)
	return
}

