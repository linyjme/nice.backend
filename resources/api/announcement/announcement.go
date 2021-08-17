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
	var ash model.Announcement
	ash.Content = r.Content
	ash.UUID = uuid.NewV4()
	ash.CreatedAt = time.Now()
	ash.UpdatedAt = time.Now()
	ash.State = r.State
	ash.V = 0
	database := utils.GetDatabase()
	insertResult, err := database.Collection("announcements").InsertOne(context.TODO(), ash)
	if err != nil {
		global.NICE_LOG.Error("插入通告失败!", zap.Any("err", err))
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	response.OkWithMessage("写入成功", c)
}

func GetAnnouncement(c *gin.Context) {
	var result []model.Announcement
	database := utils.GetDatabase()
	queryResult := database.Collection("announcements").FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	fmt.Println(queryResult)
	response.OkWithMessage("查询成功", c)
}
