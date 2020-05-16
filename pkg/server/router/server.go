package router

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"shahaohuo.com/shahaohuo/pkg/health"
	"shahaohuo.com/shahaohuo/pkg/seo"
	"shahaohuo.com/shahaohuo/pkg/server/version"
	"shahaohuo.com/shahaohuo/pkg/server/web/api"
	"shahaohuo.com/shahaohuo/pkg/server/web/html"
	"shahaohuo.com/shahaohuo/pkg/server/web/middleware"
	tpl "shahaohuo.com/shahaohuo/pkg/template"
)

func Server(templatePath, staticPath, appVersion string) *gin.Engine {
	version.SetVersion(appVersion)
	/*
		request server
	*/
	server := gin.New()
	server.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/status"},
	}))
	server.Use(gin.Recovery())

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
		"AppVersion":    version.Version.GetVersion,
		"StaticVersion": version.Version.StaticVersion,
		"IntAdd":        tpl.IntAdd,
		"IntReduce":     tpl.IntReduce,
	})
	server.LoadHTMLGlob(templatePath)
	server.Static("/static", staticPath)

	router := server.Group("/", middleware.ContentSet())

	// open htmls
	router.GET("/", html.Index)
	router.GET("/status", health.Status)
	router.GET("/login", html.Login)
	router.GET("/register", html.Login)
	router.GET("/rule", html.PageHandle("rule"))
	router.GET("/help", html.PageHandle("help"))
	router.GET("/feedback", html.PageHandle("help"))
	router.GET("/logout", html.Logout)
	router.GET("/users/:id", html.Users)
	router.GET("/users/:id/haohuos", html.UserHaohuosPage)
	router.GET("/haohuo/:id", html.Haohuo)
	router.GET("/tags/:id", html.TagsPage)

	// websocket
	// router.GET("/ws", ws.ServeWs)

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
	authorized.PATCH("/users", api.PatchUser)
	authorized.POST("/images/upload", api.UploadImage)

	return server
}
