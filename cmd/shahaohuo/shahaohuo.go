package main

import (
	"os"
	"shahaohuo.com/shahaohuo/pkg/bucket"
	"shahaohuo.com/shahaohuo/pkg/config"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/router"
	"shahaohuo.com/shahaohuo/pkg/server/storage"
)

const VERSION = "v0.0.6"

func main() {
	configPath := configPath("/opt/shahaohuo/configs/app.yaml")
	resourcesPath := resourcesPath("/opt/shahaohuo/resources")
	staticPath := resourcesPath + "/static"
	templatePath := resourcesPath + "/template/*.html"

	mysqlConfig := config.NewMysqlConfig(configPath, "yaml")
	s3Config := config.NewS3Config(configPath, "yaml")

	// init storage
	initStorage(s3Config)

	// orm
	orm.Init(mysqlConfig)
	// task
	// task.InitTasks()
	// websocket
	// ws.InitHub()

	/*
		request server
	*/
	server := router.Server(templatePath, staticPath, VERSION)
	if e := server.Run(":8080"); e != nil {
		panic(e)
	}
}

func initStorage(config config.S3Config) {
	if bkt, e := bucket.NewBucket("shahaohuo", config); e == nil {
		storage.InitStorage(bkt)
	} else {
		panic(e)
	}
}

func resourcesPath(defau string) string {
	path := os.Getenv("RESOURCES_PATH")
	if len(path) > 0 {
		return path
	}
	return defau
}

func configPath(path string) string {
	p := os.Getenv("CONFIG_PATH")
	if len(p) > 0 {
		return p
	}
	return path
}
