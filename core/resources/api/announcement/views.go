package announcement

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"niceBackend/common/transform/request"
	"niceBackend/common/transform/response"
	"niceBackend/core/api/service"
	"niceBackend/core/model"
	"niceBackend/pkg"
	"time"
)

func PostAnnouncement(c *gin.Context) {
	var r request.Announcements
	_ = c.ShouldBindJSON(&r)
	if err := pkg.Verify(r, pkg.AnnouncementVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var an model.Announcement
	an.Content = r.Content
	an.CreatedAt = time.Now()
	an.UpdatedAt = time.Now()
	an.State = r.State
	an.UUID = uuid.NewV4()
	err := service.PostAnnouncement(an)
	if err != nil {
		response.OkWithMessage("通告写入失败", c)
	} else {
		response.OkWithMessage("通告写入成功", c)
	}
}

func GetAnnouncement(c *gin.Context) {
	err, anList := service.GetAnnouncements()
	if err != nil {
		response.FailWithCode(501, c)
	} else {
		response.OkWithData(anList, c)
	}
}

func GetAnnouncementById(c *gin.Context) {
	response.OkWithMessage("查询成功", c)
}
