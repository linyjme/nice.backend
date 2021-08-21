package tags

import (
	"github.com/gin-gonic/gin"
	"niceBackend/transform/response"
)

func GetTags(c *gin.Context){
	response.OkWithMessage("ok", c)
	return
}