package admin

import (
	"github.com/gin-gonic/gin"
	"myWeb/app/admin/article"
	"myWeb/app/admin/article_ip"
	"myWeb/app/admin/category"
	"myWeb/app/admin/log"
	"myWeb/app/admin/login"
	"myWeb/app/admin/tags"
	"myWeb/middleware/midd_auth"
	"myWeb/middleware/midd_session"
)

func Routers(e *gin.Engine) {
	// 模板存储路径
	e.Use(midd_session.MiddlewareSession())
	e.GET("/login", login.Index)
	e.POST("/login", login.Login)
	e.Use(midd_auth.Init())

	e.GET("/admin", Index)
	e.GET("/logout", login.Logout)

	e.GET("/admin/category/index", category.Index)
	e.POST("/admin/category/categoryJson", category.CategoryJson)
	e.GET("/admin/category/add", category.Add)
	e.GET("/admin/category/edit", category.Edit)
	e.POST("/admin/category/updateData", category.UpdateData)

	e.GET("/admin/tags/index", tags.Index)
	e.POST("/admin/tags/tagsJson", tags.TagsJson)
	e.GET("/admin/tags/add", tags.Add)
	e.GET("/admin/tags/edit", tags.Edit)
	e.POST("/admin/tags/updateData", tags.UpdateData)

	e.GET("/admin/article/index", article.Index)
	e.POST("/admin/article/articleJson", article.ArticleJson)
	e.GET("/admin/article/add", article.Add)
	e.GET("/admin/article/edit", article.Edit)
	e.POST("/admin/article/del", article.Del)
	e.POST("/admin/article/updateData", article.UpdateData)

	e.GET("/admin/log/index", log.Index)
	e.POST("/admin/log/logJson", log.LogJson)
	e.GET("/admin/article_ip/index", article_ip.Index)
	e.POST("/admin/article_ip/acticleIpJson", article_ip.ArticleIpJson)

}
