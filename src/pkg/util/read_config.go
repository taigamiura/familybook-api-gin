package util

import (
	"fmt"

	"github.com/spf13/viper"
)
func ReadConfig() error {
  ginMode := GetGinMode()
	configFile := fmt.Sprintf("config.%s", ginMode)
	rootDir := ProjectRoot()

	viper.SetConfigName(configFile)
	viper.SetConfigType("yml")
	viper.AddConfigPath(fmt.Sprintf("%s/", rootDir))
	return viper.ReadInConfig()
}