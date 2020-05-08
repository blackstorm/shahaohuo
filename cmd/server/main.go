package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"shahaohuo.com/shahaohuo/pkg/bucket"
	"shahaohuo.com/shahaohuo/pkg/config"
	"shahaohuo.com/shahaohuo/pkg/health"
	"shahaohuo.com/shahaohuo/pkg/seo"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/storage"
	"shahaohuo.com/shahaohuo/pkg/server/task"
	"shahaohuo.com/shahaohuo/pkg/server/web/api"
	"shahaohuo.com/shahaohuo/pkg/server/web/html"
	"shahaohuo.com/shahaohuo/pkg/server/web/middleware"
	"shahaohuo.com/shahaohuo/pkg/server/web/ws"
)

func main() {
	configPath := configPath("/opt/shahaohuo/configs/app.yaml")
	resourcesPath := resourcesPath("/opt/shahaohuo/resources")
	staticPath := resourcesPath + "/static"
	templatePath := resourcesPath + "/template/*.html"

	mysqlConfig := config.NewMysqlConfig(configPath, "yaml")
	s3Config := config.NewS3Config(configPath, "yaml")

	bkt, e := bucket.NewBucket("shahaohuo", s3Config, true)
	if e != nil {
		panic(e)
	}
	storage.InitStorage(bkt)
	// orm
	orm.Init(mysqlConfig)
	// task
	task.InitTasks()
	// websocket
	ws.InitHub()

	/*
		gin server
	*/
	server := gin.Default()
	// max size 4m TODO is not working ?
	server.MaxMultipartMemory = 4 << 20
	//middleware
	//xssMdlwr := &xss.XssMw{
	//	FieldsToSkip: []string{"password"},
	//	BmPolicy:     "UGCPolicy",
	//}
	//// server.Use(xssMdlwr.RemoveXss())
	// statics
	server.SetFuncMap(template.FuncMap{
		"BaseKeyWorlds": seo.BaseKeyWorlds,
	})
	server.LoadHTMLGlob(templatePath)
	server.Static("/static", staticPath)

	router := server.Group("/", middleware.ContentSet())

	// open htmls
	router.GET("/", html.Index)
	router.GET("/status", health.Status)
	router.GET("/login", html.Login)
	router.GET("/register", html.Login)
	router.GET("/rule", html.Rule)
	router.GET("/logout", html.Logout)
	router.GET("/users/:id", html.Users)
	router.GET("/haohuo/:id", html.Haohuo)

	// websocket
	router.GET("/ws", ws.ServeWs)

	// auth page
	cookieAuth := router.Group("/", middleware.CookieAuth())
	cookieAuth.GET("/share", html.Share)
	cookieAuth.GET("/settings", html.Settings)

	// api open
	router.POST("/open/api/v1/login", api.Login)
	router.POST("/open/api/v1/register", api.Register)

	// auth api
	authorized := router.Group("/api/v1", middleware.ApiAuth())
	authorized.PUT("/haohuos/:id", api.CreateOrUpdateHaohuo)
	authorized.PUT("/haohuos/:id/favorite", api.FavoriteHaohuo)
	authorized.PUT("/haohuos/:id/comment", api.CommentHaohuo)
	authorized.POST("/images/upload", api.UploadImage)

	if e := server.Run(":8080"); e != nil {
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
