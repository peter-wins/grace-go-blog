package log

import (
	"github.com/gin-gonic/gin"
	"myWeb/model"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/log/index.html", gin.H{})
}
func LogJson(c *gin.Context) {
	method := c.Request.Method
	if method == "POST" {
		var optPage model.OptPage
		var searchWhere model.Log
		if err := c.ShouldBind(&optPage); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		if err := c.ShouldBind(&searchWhere); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		log, total := model.GetLogList(&searchWhere, &optPage)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "", "data": log, "recordsTotal": total, "recordsFiltered": total})
		return
	} else {
		c.HTML(http.StatusOK, "/admin/log/index.html", gin.H{})
	}
}
