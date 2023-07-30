package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Port     string
	UserName string
	Addr     string
	Password string
	Charset  string
	DbName   string
}

var (
	HttpPort string
	DbConfig DBConfig
)

func Init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("%+v", err)
		return
	}

	InitConfig()
}

func InitConfig() {
	HttpPort = viper.GetString("settings.application.port")
	initMysqlSetting()
}

func initMysqlSetting() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("%+v", err)
		return
	}
	DbConfig.Charset = viper.GetString("settings.database.charset")
	DbConfig.UserName = viper.GetString("settings.database.userName")
	DbConfig.Password = viper.GetString("settings.database.password")
	DbConfig.Addr = viper.GetString("settings.database.addr")
	DbConfig.DbName = viper.GetString("settings.database.dbName")
	DbConfig.Port = viper.GetString("settings.database.port")
}
