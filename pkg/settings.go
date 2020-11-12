package pkg

import "github.com/spf13/viper"

func InitSettings() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")
	err := viper.ReadInConfig()
	return err
}
