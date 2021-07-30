package model

type User struct {
	BASE_MODEL
	ID                         uint   `gorm:"primarykey;AUTO_INCREMENT"`                   // 主键ID
	Account                    string `json:"account" gorm:"type:varchar(128);not null"`   // 用户登录名
	Password                   string `json:"-"  gorm:"type:varchar(256)"`                 //
	Name                       string `json:"name" gorm:"type:varchar(128)"`               //
	Email                      string `json:"email" gorm:"type:varchar(256)"`              //
	Phone                      string `json:"phone" gorm:"type:varchar(32)"`               //
	Company                    string `json:"company" gorm:"type:varchar(256)"`            //
	Address                    string `json:"address" gorm:"type:varchar(256)"`            //
	Remark                     string `json:"remark" gorm:"type:varchar(512)"`             //
	Status                     uint8  `json:"status" gorm:"default:0"`                     //
	AccountType                uint8  `json:"accountType" gorm:"default:2;not null"`       //
	UploadSpeed                uint32 `json:"uploadSpeed" gorm:"default:0"`                //
	DownloadSpeed              uint32 `json:"downloadSpeed" gorm:"default:0"`              //
	UserLockedFlag             bool   `json:"userLockedFlag"`                              //
	LoginFailTimes             uint8  `json:"loginFailTimes" gorm:"default:0"`             //
	UserLockedStartTime        uint8  `json:"userLockedStartTime" gorm:"default:0"`        //
	UserLockedFinishTime       uint8  `json:"userLockedFinishTime" gorm:"default:0"`       //
	NeedUpdatePwd              bool   `json:"needUpdatePwd" gorm:"default:false"`          //
	UserHomeSize               uint32 `json:"userHomeSize" gorm:"default:0"`               //
	UserHomeTotal              uint32 `json:"userHomeTotal" gorm:"default:0"`              //
	File_count                 uint16 `json:"file_count" gorm:"default:0"`                 //
	Folder_count               uint16 `json:"folder_count" gorm:"default:0"`               //
	UserHomeFlag               bool   `json:"userHomeFlag"`                                //
	EmailSenderType            uint8  `json:"emailSenderType" gorm:"default:0"`            //
	TransNotificationSwitch    bool   `json:"transNotificationSwitch"`                     //
	ShareEmailNotificationType uint8  `json:"shareEmailNotificationType" gorm:"default:0"` //
	EmailAccount               string `json:"emailAccount"`                                //
	EmailPassword              string `json:"emailPassword"`                               //
	EmailSmtpHost              string `json:"emailSmtpHost"`                               //
	EmailFromName              string `json:"emailFromName"`                               //
	EmailSmtpPort              string `json:"emailSmtpPort"`                               //
	EmailEncryptType           string `json:"emailEncryptType"`                            //
	EmailType                  string `json:"emailType"`                                   //
	StorageId                  uint16 `json:"storageId" gorm:"default:0"`                  //
	Extend_A                   string `json:"extend_a"`                                    //
}
