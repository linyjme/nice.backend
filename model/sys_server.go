package model

type ServerConfig struct {
	BASE_MODEL
	ID                    uint   `gorm:"primarykey"`                       // 主键ID
	Name                  string `json:"name" gorm:"type:varchar(128)"`    //
	Version               string `json:"version" gorm:"type:varchar(128)"` //
	Permission            uint16 `json:"permission" `                      //
	Email                 string `json:"email" gorm:"type:varchar(128)"`   //
	Device                string `json:"device" gorm:"type:varchar(128)"`  //
	Host                  string `json:"host" gorm:"type:varchar(128)"`    //
	Os                    string `json:"os" gorm:"type:varchar(128)"`      //
	Status                uint8    `json:"status"`                           //
	Loglevel              uint8    `json:"loglevel"`                         //
	UploadSpeed           uint8    `json:"uploadSpeed"`                      //
	DownloadSpeed         uint8    `json:"downloadSpeed"`                    //
	UserLockedFlag        string `json:"userLockedFlag"`                   //
	LoginFailTimes        uint8    `json:"loginFailTimes"`                   //
	UserLockedStartTime   uint8    `json:"userLockedStartTime"`              //
	UserLockedFinishTime  uint8    `json:"userLockedFinishTime"`             //
	NeedUpdatePwd         bool   `json:"needUpdatePwd"`                    //
	Remark                string `json:"remark"`                           //
	SpeedLimitType        string `json:"speedLimitType"`                   //
	PackageSize           string `json:"packageSize"`                      //
	EnableProxy           string `json:"enableProxy"`                      //
	ProxyHost             string `json:"proxyHost"`                        //
	ProxyPort             string `json:"proxyPort"`                        //
	License               string `json:"license"`                          //
	Compression           string `json:"compression"`                      //
	Outbound              string `json:"outbound"`                         //
	Inbound               string `json:"inbound"`                          //
	LogRootPath           string `json:"logRootPath"`                      //
	ServerHandler         string `json:"serverHandler"`                    //
	ServerMode            string `json:"serverMode"`                       //
	ServerLanguage        string `json:"serverLanguage"`                   //
	EmailSender           string `json:"emailSender"`                      //
	EmailAccount          string `json:"emailAccount"`                     //
	EmailPassword         string `json:"emailPassword"`                    //
	EmailSmtpHost         string `json:"emailSmtpHost"`                    //
	EmailSmtpPort         string `json:"emailSmtpPort"`                    //
	EmailFromName         string `json:"emailFromName"`                    //
	EmailEncryptType      string `json:"emailEncryptType"`                 //
	Cdn                   string `json:"cdn"`                              //
	TcpEnable             string `json:"tcpEnable"`                        //
	TcpDelay              string `json:"tcpDelay"`                         //
	TcpPortRangeLeft      string `json:"tcpPortRangeLeft"`                 //
	TcpPortRangeRight     string `json:"tcpPortRangeRight"`                //
	AdDomainType          string `json:"adDomainType"`                     //
	AdServer              string `json:"adServer"`                         //
	AdDomainName          string `json:"adDomainName"`                     //
	AdAdminAccount        string `json:"adAdminAccount"`                   //
	AdAdminPassword       string `json:"adAdminPassword"`                  //
	AdOuConfig            string `json:"adOuConfig"`                       //
	AdGroupConfig         string `json:"adGroupConfig"`                    //
	AdSyncGroupToRole     string `json:"adSyncGroupToRole"`                //
	AdSyncConfig          string `json:"adSyncConfig"`                     //
	AdSyncPath            string `json:"adSyncPath"`                       //
	AdSyncInterval        string `json:"adSyncInterval"`                   //
	Ssl_port              string `json:"ssl_port"`                         //
	No_ssl_port           string `json:"no_ssl_port"`                      //
	Blocking_on_close     string `json:"blocking_on_close"`                //
	Multi_ip              string `json:"multi_ip"`                         //
	AccessKeyId           string `json:"AccessKeyId"`                      //
	AccessKeySecret       string `json:"AccessKeySecret"`                  //
	BucketName            string `json:"BucketName"`                       //
	OssBuff               string `json:"OssBuff"`                          //
	EauthSmtpHost         string `json:"eauthSmtpHost"`                    //
	EauthSmtpPort         string `json:"eauthSmtpPort"`                    //
	EauthEncryptType      string `json:"eauthEncryptType"`                 //
	AuthWay               string `json:"authWay"`                          //
	UploadSuffixLimit     string `json:"uploadSuffixLimit"`                //
	SuffixList            string `json:"suffixList"`                       //
	SuffixBlackList       string `json:"suffixBlackList"`                  //
	MaxLockTimes          string `json:"maxLockTimes"`                     //
	LockPeriod            string `json:"lockPeriod"`                       //
	CertSelect            string `json:"certSelect"`                       //
	OptlogPeriod          string `json:"optlogPeriod"`                     //
	TranslogPeriod        string `json:"translogPeriod"`                   //
	UserHomeCountPeriod   string `json:"userHomeCountPeriod"`              //
	UserHomeSize          string `json:"userHomeSize"`                     //
	UserHomeTotal         string `json:"userHomeTotal"`                    //
	File_count            string `json:"file_count"`                       //
	Folder_count          string `json:"folder_count"`                     //
	UserHomeFlag          string `json:"userHomeFlag"`                     //
	SyncSwitch            string `json:"syncSwitch"`                       //
	Proxy_process_enable  string `json:"proxy_process_enable"`             //
	Proxy_process_num     string `json:"proxy_process_num"`                //
	UserPwdSwitch         string `json:"userPwdSwitch"`                    //
	EauthType             string `json:"eauthType"`                        //
	CHashVerifySwitch     string `json:"cHashVerifySwitch"`                //
	CTransEncryptSwitch   string `json:"cTransEncryptSwitch"`              //
	TCPFileSizeLimit      string `json:"TCPFileSizeLimit"`                 //
	AutoUpdateSwitch      string `json:"autoUpdateSwitch"`                 //
	Enable_report_event   string `json:"enable_report_event"`              //
	Enable_report_traffic string `json:"enable_report_traffic"`            //
	Enable_report_space   string `json:"enable_report_space"`              //
	TransUnit             string `json:"transUnit"`                        //
	BindDeviceSwitch      string `json:"bindDeviceSwitch"`                 //
	WaterMarkSwitch       string `json:"waterMarkSwitch"`                  //
	WmPosition            string `json:"wmPosition"`                       //
	WeakPwdList           string `json:"weakPwdList"`                      //
	WeakSwitch            string `json:"weakSwitch"`                       //
	Extend_A              string `json:"extend_a"`                         //
}

func (s ServerConfig) TableName() string {
	return "t_server_config"
}
