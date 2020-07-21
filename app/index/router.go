package index

import (
	"github.com/gin-gonic/gin"
	"myWeb/middleware/midd_index"
)

func Routers(e *gin.Engine) {
	e.Use(midd_index.TagsNav)
	e.Use(midd_index.HeaderAndRightNav)
	e.GET("/", Index)
	e.GET("/index", Index)
	e.GET("/index/category", CategoryIndex)
	e.GET("/index/tags", TagsIndex)
	e.GET("/index/article/detail", DetailArticle)
	e.GET("/index/article/search", SearchArticle)
}
