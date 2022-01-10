package service

import (
	"errors"
	"go.uber.org/zap"
	"niceBackend/common/global"
	"niceBackend/internal/model"
)

func GetAnnouncements() (err error, list interface{}) {
	var annList []model.Announcement
	if err = global.NiceDb.Find(&annList).Error; err != nil {
		return errors.New("查询通告失败"), &annList
	} else {
		return nil, &annList
	}
}

func PostAnnouncement(an model.Announcement) (err error) {
	err = global.NiceDb.Create(&an).Error
	if err != nil {
		global.NiceLog.Error("插入通告失败!", zap.Any("err", err))
	}
	return err

}
