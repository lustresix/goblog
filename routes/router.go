package routes

import (
	"github.com/gin-gonic/gin"
	v1 "goblog/api/v1"
	"goblog/config"
	"goblog/middleware"
)

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JWT())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCate)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArtInfo)
		auth.GET("admin/article", v1.GetArt)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		// 上传文件
		auth.POST("upload", v1.UpLoad)
	}
	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 文章分类信息模块
		router.GET("category", v1.GetCate)

		// 文章模块
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)

		// 登录控制模块
		router.POST("login", v1.UserLogin)
	}
	r.Run(config.HttpPort)
}
