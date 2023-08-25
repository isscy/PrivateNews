package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"serverGo/controller"
	"serverGo/foundation/logger"
	"serverGo/foundation/middlewares"
	"serverGo/foundation/resp"
	"time"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode { // 设置成发布模式
		gin.SetMode(gin.ReleaseMode)
	}
	//初始化 gin Engine  新建一个没有任何默认中间件的路由
	r := gin.New()
	//设置中间件
	r.Use(logger.GinLogger(),
		Cors(),
		logger.GinRecovery(true), // Recovery 中间件会 recover掉项目可能出现的panic，并使用zap记录相关日志
		middlewares.RateLimitMiddleware(2*time.Second, 40), // 每两秒钟添加十个令牌  全局限流
	)
	r.LoadHTMLFiles("webpage/templates/index.html") // 加载html
	r.Static("/static", "webpage/static")           // 加载静态文件
	r.GET("/templates/email", func(context *gin.Context) {
		context.HTML(http.StatusOK, "email.html", nil)
	})
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	// 登录注册相关模块
	auth := r.Group("/api/auth")
	auth.POST("/login", controller.LoginController)         // 登陆业务
	auth.GET("/test", controller.TestController)            // 登陆业务
	auth.POST("/refresh_token", controller.LoginController) // 刷新accessToken
	// 业务相关模块
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		// 用户相关
		v1.GET("/user/getUserInfo", controller.UserInfoController) // 获取用户信息

		// Tag相关
		v1.GET("/tag/getBuiltInTag", controller.BuiltInTagController) // 获取内置tagList
		v1.POST("/tag", controller.LoginController)                   // 创建一组tag

		// 内容相关
		v1.POST("/publish", controller.PublishOneArticleController) // 创建一篇新闻笔记
		v1.POST("/newsPage", controller.GetArticlePageController)   //  分页获取新闻

	}
	pprof.Register(r) // 注册pprof相关路由
	r.NoRoute(func(c *gin.Context) {
		resp.Fail(c, resp.InvalidPath)
	})
	return r
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}

}
