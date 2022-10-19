package pkg

import (
  "github.com/spf13/viper"
)

func NewViper() *viper.Viper {
  return viper.New()
}
