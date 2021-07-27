package model

type User struct {
	BASE_MODEL
	Account              string `json:"account" gorm:"comment:用户账号"` // 用户登录名
	Password             string `json:"-"  gorm:"default:用户登录密码"`    //
	Permission           int    `json:"permission" `                 //
	Name                 string `json:"name"`                        //
	Email                string `json:"email"`                       //
	Phone                string `json:"phone"`                       //
	Company              string `json:"company"`                     //
	Address              string `json:"address"`                     //
	Remark               string `json:"remark"`                      //
	Status               int    `json:"status"`                      //
	AccountType          int    `json:"accountType"`                 //
	UploadSpeed          int    `json:"uploadSpeed"`                 //
	DownloadSpeed        int    `json:"downloadSpeed"`               //
	UserLockedFlag       string `json:"userLockedFlag"`              //
	LoginFailTimes       int    `json:"loginFailTimes"`              //
	UserLockedStartTime  int    `json:"userLockedStartTime"`         //
	UserLockedFinishTime int    `json:"userLockedFinishTime"`        //
	NeedUpdatePwd        bool   `json:"needUpdatePwd"`               //
	Extend_A             string `json:"extend_a"`                    //
}
