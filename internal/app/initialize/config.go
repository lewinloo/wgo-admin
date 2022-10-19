package initialize

import (
	"fmt"
	"gin_template/internal/app/config"
	"gin_template/pkg"
	"github.com/fsnotify/fsnotify"
	"os"
)

func Config(filePath string, config *config.Application) {
	v := pkg.NewViper()
	rootPath, _ := os.Getwd()
	v.SetConfigFile(rootPath + filePath)
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(any(fmt.Errorf("Fatal error config file: %s \n", err)))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(config); err != nil {
		fmt.Println(err)
	}
}
