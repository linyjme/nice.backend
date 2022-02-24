package article_repo

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"niceBackend/internal/repository/db_repo/categroy_repo"
	"niceBackend/internal/repository/db_repo/tag_repo"
)

type Article struct {
	gorm.Model
	UUID    uuid.UUID `json:"_id "gorm:"comment:announcementUUID;column:_id" bson:"_id"`
	Password    string                 `json:"password" bson:"password" gorm:"type:varchar(64);column:password"` //
	Keywords    string                 `json:"keywords" bson:"keywords" gorm:"type:varchar(128);not null"`       //
	State       uint8                  `json:"state" bson:"state" gorm:"default:1;not null"`                     //
	Public      uint8                  `json:"public" bson:"public" gorm:"default:1;not null"`                   //
	Origin      uint8                  `json:"origin" bson:"origin" gorm:"default:0;not null"`                   //
	Tag         tag_repo.Tags          `json:"tag" gorm:"foreignKey:TagsId;references:TagsId;comment:文字标签"`
	Category    categroy_repo.Category `json:"category" bson:"category" gorm:"foreignKey:CategoryId;references:CategoryId;comment:文章分类"` //
	Description string                 `json:"description"  bson:"description" gorm:"type:varchar(256)"`                                    //
	Title       string                 `json:"title" bson:"title" gorm:"type:varchar(64)"`                                                  //
	Content     string                 `json:"content" bson:"content" gorm:"type:varchar(1024)"`                                            //
	Thumb       string                 `json:"thumb" bson:"thumb" gorm:"type:varchar(128)"`                                                 //
	Meta        string                 `json:"meta" bson:"meta" gorm:"type:varchar(256)"`                                                   //
	Extends     string                 `json:"extends" bson:"extends" gorm:"type:varchar(256)"`                                             //
}

func (A Article) TableName() string {
	return "tb_articles"
}
