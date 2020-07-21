package tags

import (
	"github.com/gin-gonic/gin"
	"myWeb/model"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/tags/index.html", gin.H{})
}
func TagsJson(c *gin.Context) {
	method := c.Request.Method
	if method == "POST" {
		var optPage model.OptPage
		var searchWhere model.Tags
		if err := c.ShouldBind(&optPage); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		if err := c.ShouldBind(&searchWhere); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		tags, total := model.GetTagsList(&searchWhere, &optPage)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "", "data": tags, "recordsTotal": total, "recordsFiltered": total})
		return
	} else {
		c.HTML(http.StatusOK, "/admin/tags/index.html", gin.H{})
	}
}
func Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/tags/add.html", gin.H{"tagsStatus": model.TagsStatus})
}
func Edit(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "非法参数"})
		return
	}
	Ids, _ := strconv.Atoi(id)
	data := model.GetTagsInfo(Ids)
	c.HTML(http.StatusOK, "admin/tags/edit.html", gin.H{"data": data, "tagsStatus": model.TagsStatus})
}

func UpdateData(c *gin.Context) {
	var tags model.Tags
	err := c.ShouldBind(&tags)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "msg": err.Error()})
	}
	data := model.UpdateTags(&tags)
	c.JSON(http.StatusOK, gin.H{"code": data.Code, "msg": data.Msg, "url": data.Url})
}
