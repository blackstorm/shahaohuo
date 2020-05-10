package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type S3Config struct {
	Id               string
	Secret           string
	Endpoint         string
	AccessEndpoint   string
	Region           string
	DisableSSL       bool
	S3ForcePathStyle bool
}

func NewS3Config(configFilePath string, configType string) S3Config {
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	s3 := viper.GetStringMap("s3")

	return S3Config{
		Id:               s3["id"].(string),
		Secret:           s3["secret"].(string),
		Endpoint:         s3["endpoint"].(string),
		AccessEndpoint:   s3["access-endpoint"].(string),
		Region:           s3["region"].(string),
		DisableSSL:       s3["disable-ssl"].(bool),
		S3ForcePathStyle: s3["s3-force-path-style"].(bool),
	}
}
