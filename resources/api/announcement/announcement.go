package announcement

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	oson "gopkg.in/mgo.v2/bson"
	"log"
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
	//ash.UUID _ := uuid.New()
	ash.CreatedAt = time.Now()
	ash.UpdatedAt = time.Now()
	ash.State = r.State
	uid := oson.NewObjectId().String()
	ash.UUID = uid
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
	var results []model.Announcement
	database := utils.GetDatabase()
	findOptions := options.Find()
	findOptions.SetLimit(2)
	cur, err := database.Collection("announcements").Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		global.NICE_LOG.Error("查询失败", zap.Any("err", err))
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem model.Announcement
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// Close the cursor once finished
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	response.OkWithMessage("查询成功", c)
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
