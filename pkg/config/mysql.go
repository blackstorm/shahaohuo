package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

type MySQLConfig struct {
	Host     string
	UserName string
	Password string
	Port     int
	Database string
	Params   string
	ShowSQL  bool
}

func (m MySQLConfig) Addr() string {
	return m.Host + ":" + strconv.Itoa(m.Port)
}

func NewMysqlConfig(configFilePath string, configType string) MySQLConfig {
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	mysql := viper.GetStringMap("mysql")

	return MySQLConfig{
		Host:     mysql["host"].(string),
		UserName: mysql["username"].(string),
		Password: mysql["password"].(string),
		Port:     mysql["port"].(int),
		Database: mysql["database"].(string),
		ShowSQL:  mysql["showsql"].(bool),
	}
}
