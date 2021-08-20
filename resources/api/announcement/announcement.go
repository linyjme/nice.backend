package announcement

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"niceBackend/common/global"
	"niceBackend/model"
	"niceBackend/service"
	"niceBackend/transform/request"
	"niceBackend/transform/response"
	"niceBackend/utils"
	"time"
)

func PostAnnouncement(c *gin.Context) {
	var r request.Announcements
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.AnnouncementVerify); err != nil {
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
	var result []model.Announcement
	database := utils.GetDatabase()
	filter := bson.D{{"_id", "aaaa"}}
	//filter :bson.D{{}}
	err := database.Collection("announcements").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		global.NICE_LOG.Error("查询失败", zap.Any("err", err))
	}
	fmt.Println(result)
	response.OkWithMessage("查询成功", c)
}
