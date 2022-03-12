package initialize

import (
	"fmt"

	"niceBackend/common/global"
	"niceBackend/pkg"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	configIni := pkg.GetConfigIniPath()
	v := viper.New()
	v.SetConfigFile(configIni)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.NiceConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.NiceConfig); err != nil {
		fmt.Println(err)
	}
	return v
}
