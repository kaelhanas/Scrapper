package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

type ViperConfig struct {
	Viper *viper.Viper
}

var once sync.Once
var viperInstance *ViperConfig

func GetViper() (*viper.Viper){
	once.Do(func (){

		v := viper.New()
		v.SetConfigFile("config.yml")
		v.AddConfigPath(".")
		v.AddConfigPath(`D:\Documents\Projects\Scrapper_Golang\Scrapper_POC\config.yml`)

		if err := v.ReadInConfig() ; err != nil {
			log.Fatalf("failed to ReadInConfig :: %v", err)
		}
		viperInstance = &ViperConfig{v}

	})

	return viperInstance.Viper
}
