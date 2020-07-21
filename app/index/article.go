package index

import (
	"github.com/gin-gonic/gin"
	"myWeb/model"
	"net/http"
	"strconv"
	"time"
)

func DetailArticle(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	article := model.ApiArticle(id)
	assginData := model.MergeAssignData(c, gin.H{
		"article": article,
	})
	ip := c.ClientIP()
	var article_ip model.ArticleIp
	article_ip.Ip = ip
	article_ip.ArticleId, _ = strconv.Atoi(id)
	article_ip.CreatedAt = time.Now().Unix()
	model.InsertArticleIp(&article_ip)
	c.HTML(http.StatusOK, "index/article/detail.html", assginData)

}
func SearchArticle(c *gin.Context) {
	var optPage model.OptPage
	optPage.Length = 10
	optPage.Start = 0
	var searchWhere model.SearchWhere
	title := c.DefaultQuery("s", "")
	searchWhere.Title = title

	article, total := model.ApiArticleList(&searchWhere, &optPage)
	assginData := model.MergeAssignData(c, gin.H{
		"articles": article,
		"total":    total,
	})
	c.HTML(http.StatusOK, "index/article/search.html", assginData)

}
