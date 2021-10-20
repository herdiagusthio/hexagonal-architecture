package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppPort        int    `mapstructure:"app_port"`
	AppEnvironment string `mapstructure:"app_environment"`
	DbDriver       string `mapstructure:"db_driver"`
	DbHost         string `mapstructure:"db_host"`
	DbPort         int    `mapstructure:"db_port"`
	DbUsername     string `mapstructure:"db_username"`
	DbPassword     string `mapstructure:"db_password"`
	DbName         string `mapstructure:"db_name"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

//GetConfig Init config in singleton way
func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	//re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	defaultConfig.AppPort = 8000
	defaultConfig.AppEnvironment = ""
	defaultConfig.DbDriver = "mongodb"
	defaultConfig.DbHost = "localhost"
	defaultConfig.DbPort = 27017
	defaultConfig.DbUsername = ""
	defaultConfig.DbPassword = ""
	defaultConfig.DbName = "alta_store"

	viper.AutomaticEnv()
	viper.SetEnvPrefix("alta_store")
	viper.BindEnv("app_port")
	viper.BindEnv("app_environment")
	viper.BindEnv("db__driver")
	viper.BindEnv("db_host")
	viper.BindEnv("db_port")
	viper.BindEnv("db_username")
	viper.BindEnv("db_password")
	viper.BindEnv("db_name")

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
		return &defaultConfig
	}

	return &finalConfig
}
