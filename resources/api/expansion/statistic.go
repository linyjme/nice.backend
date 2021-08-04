package expansion

import (
	"github.com/gin-gonic/gin"
	"niceBackend/transform/response"
)

func Statistic(c *gin.Context) {
	response.OkWithMessage("注册成功", c)
}
