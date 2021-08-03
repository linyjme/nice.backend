package model

type ServerConfig struct {
	BASE_MODEL
	ID       uint   `gorm:"primarykey"`                       // 主键ID
	Name     string `json:"name" gorm:"type:varchar(128)"`    //
	Version  string `json:"version" gorm:"type:varchar(128)"` //
	Extend_A string `json:"extend_a"`                         //
}

func (s ServerConfig) TableName() string {
	return "t_server_config"
}