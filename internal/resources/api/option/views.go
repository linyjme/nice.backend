package option

import (
	"github.com/gin-gonic/gin"
	"niceBackend/common/transform/response"
)

func GetOption(c *gin.Context){
	var result map[string]interface{}
	result = make(map[string]interface{})
	response.OkWithData(result, c)
	return
}

