package article_ip

import (
	"github.com/gin-gonic/gin"
	"myWeb/model"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/article_ip/index.html", gin.H{})
}
func ArticleIpJson(c *gin.Context) {
	method := c.Request.Method
	if method == "POST" {
		var optPage model.OptPage
		var searchWhere model.ArticleIp
		if err := c.ShouldBind(&optPage); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		if err := c.ShouldBind(&searchWhere); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		ipList, total := model.GetArticleIpList(&searchWhere, &optPage)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "", "data": ipList, "recordsTotal": total, "recordsFiltered": total})
		return
	} else {
		c.HTML(http.StatusOK, "/admin/article_ip/index.html", gin.H{})
	}
}
