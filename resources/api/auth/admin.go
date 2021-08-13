package auth

import (
	"github.com/gin-gonic/gin"
	"niceBackend/transform/response"
)

func Admin(c *gin.Context){
	response.FailWithCode(201, c)
	return
}
