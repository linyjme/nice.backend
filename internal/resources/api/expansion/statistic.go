package expansion

import (
	"github.com/gin-gonic/gin"
	"niceBackend/common/transform/request"
	"niceBackend/common/transform/response"
	"niceBackend/internal/api/service"
	"niceBackend/pkg"
)

func GetStatistic(c *gin.Context) {
	service.GetStatistic()
	var result map[string]interface{}
	result = make(map[string]interface{})
	var data [1]int
	result["data"] = data
	response.OkWithData(result, c)
}

func PostStatistic(c *gin.Context) {
	var r request.Register
	_ = c.ShouldBindJSON(&r)
	if err := pkg.Verify(r, pkg.AnnouncementVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service.GetStatistic()
	response.OkWithMessage("写入成功", c)
}
