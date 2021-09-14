package tags

import (
	"github.com/gin-gonic/gin"
	"niceBackend/common/transform/response"
)

func GetTags(c *gin.Context){
	response.OkWithMessage("ok", c)
	return
}