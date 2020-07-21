package index

import (
	"github.com/gin-gonic/gin"
	"myWeb/model"
	"net/http"
)

func Index(c *gin.Context) {
	var optPage model.OptPage
	optPage.Length = 10
	optPage.Start = 0
	var searchWhere model.SearchWhere
	if err := c.ShouldBind(&searchWhere); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
	}
	article, total := model.ApiArticleList(&searchWhere, &optPage)
	assginData := model.MergeAssignData(c, gin.H{
		"articles": article,
		"total":    total,
	})
	c.HTML(http.StatusOK, "index/index/index.html", assginData)
}
