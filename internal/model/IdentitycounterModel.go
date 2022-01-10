package model

type IdentityCounters struct {
	Id    string `json:"_id "gorm:"comment:IdentityCountersUUID" bson:"_id"`
	Model string `json:"model" bson:"model" gorm:"type:varchar(128)"` //
	Field string `json:"id" bson:"id" gorm:"type:varchar(128)"`       //
	Count uint64 `json:"count" bson:"count"`
}
