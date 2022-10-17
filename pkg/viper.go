package pkg

import (
	"fmt"
	"gin_template/internal/app/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Viper(filePath string) *viper.Viper {
	rootPath, _ := os.Getwd()
	v := viper.New()
	v.SetConfigFile(rootPath + filePath)
	v.SetConfigType("yaml")
	fmt.Println()
	err := v.ReadInConfig()
	if err != nil {
		panic(any(fmt.Errorf("Fatal error config file: %s \n", err)))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
