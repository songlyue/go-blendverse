package initialize

import (
	"blendverse/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile("./setting-dev.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal err config file %s", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
