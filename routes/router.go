package routes

import (
	"github.com/gin-gonic/gin"
	v1 "goblog/api/v1"
	"goblog/middleware"
	"goblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// 上传文件
		auth.POST("upload",v1.UpLoad)
	}

	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.GET("category", v1.GetCategory)
		routerV1.GET("article", v1.GetArticle)
		routerV1.GET("category/articles/:id", v1.GetCategoryArticle)
		routerV1.GET("article/info/:id", v1.GetArticleInfo)
		routerV1.POST("login", v1.Login)
	}
	_ = r.Run(utils.HttpPort)
}
