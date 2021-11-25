package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"niceBackend/common/constants"
)

type Response struct {
	Code int         `json:"code"`
	Result interface{} `json:"result"`
	Message  string      `json:"message"`
	Status  string      `json:"status"`
}

const (
	ERROR   = 7
	SUCCESS = 200
)

func Result(code int, result interface{}, message string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		result,
		message,
		message,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(result interface{}, c *gin.Context) {
	Result(SUCCESS, result, "success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "failed", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithCode(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, constants.ResultCode[code], c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
