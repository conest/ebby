package cfgloader

import (
	"fmt"

	"github.com/spf13/viper"
)

// Init viper config
func Init() *(viper.Viper) {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath(".")

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("[Viper config] Fatal error config file: %s", err))
	}

	return v
}