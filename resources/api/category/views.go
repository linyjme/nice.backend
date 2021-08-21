package category

import (
	"github.com/gin-gonic/gin"
	"niceBackend/transform/response"
)

func GetCategory(c *gin.Context){
	response.OkWithMessage("ok", c)
	return
}

