package category

import (
	"github.com/gin-gonic/gin"
	"myWeb/model"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/category/index.html", gin.H{})
}
func CategoryJson(c *gin.Context) {
	method := c.Request.Method
	if method == "POST" {
		var optPage model.OptPage
		var searchWhere model.Category
		if err := c.ShouldBind(&optPage); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		if err := c.ShouldBind(&searchWhere); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		category, total := model.GetCategoryList(&searchWhere, &optPage)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "", "data": category, "recordsTotal": total, "recordsFiltered": total})
		return
	} else {
		c.HTML(http.StatusOK, "/admin/category/index.html", gin.H{})
	}
}
func Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/category/add.html", gin.H{
		"categoryStatus": model.CategoryStatus,
		"ctypes": model.Ctypes,
	})
}
func Edit(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "请选择服务器"})
		return
	}
	Ids, _ := strconv.Atoi(id)
	data := model.GetCategoryInfo(Ids)
	c.HTML(http.StatusOK, "admin/category/edit.html", gin.H{
		"data": data,
		"categoryStatus": model.CategoryStatus,
		"ctypes": model.Ctypes,
	})
}

func UpdateData(c *gin.Context) {
	var category model.Category
	err := c.ShouldBind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "msg": err.Error()})
	}
	data := model.UpdateCategory(&category)
	c.JSON(http.StatusOK, gin.H{"code": data.Code, "msg": data.Msg, "url": data.Url})
}
