package initialize

import (
	"niceBackend/common/global"
	"niceBackend/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	configIni := utils.GetConfigIniPath()
	v := viper.New()
	v.SetConfigFile(configIni)
	v.SetConfigType("ini")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.RAY_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.RAY_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
