package service

import (
	"errors"
	"go.uber.org/zap"
	"niceBackend/common/global"
	"niceBackend/internal/repository/db_repo/announcement_repo"
)

func GetAnnouncements() (err error, list interface{}) {
	var annList []announcement_repo.Announcement
	if err = global.NiceDb.Find(&annList).Error; err != nil {
		return errors.New("查询通告失败"), &annList
	} else {
		return nil, &annList
	}
}

func PostAnnouncement(an announcement_repo.Announcement) (err error) {
	err = global.NiceDb.Create(&an).Error
	if err != nil {
		global.NiceLog.Error("插入通告失败!", zap.Any("err", err))
	}
	return err

}
