package category

import (
	"github.com/gin-gonic/gin"
	"niceBackend/common/transform/response"
)

func GetCategory(c *gin.Context){
	response.OkWithMessage("ok", c)
	return
}

