package article

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myWeb/model"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/article/index.html", gin.H{})
}
func ArticleJson(c *gin.Context) {
	method := c.Request.Method
	if method == "POST" {
		var optPage model.OptPage
		var searchWhere model.Article
		if err := c.ShouldBind(&optPage); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		if err := c.ShouldBind(&searchWhere); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		}
		article, total := model.GetArticleList(&searchWhere, &optPage)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "", "data": article, "recordsTotal": total, "recordsFiltered": total})
		return
	} else {
		c.HTML(http.StatusOK, "/admin/category/index.html", gin.H{})
	}
}
func Add(c *gin.Context) {
	categorys := model.GetCategorys()
	tags := model.GetTags()
	c.HTML(http.StatusOK, "admin/article/add.html", gin.H{
		"categorys":     categorys,
		"tags":          tags,
		"articleStatus": model.ArticleStatus,
	})
}
func Edit(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "请选择服务器"})
		return
	}
	Ids, _ := strconv.Atoi(id)
	articleData, categorys, tags := model.GetArticleInfo(Ids)

	c.HTML(http.StatusOK, "admin/article/edit.html", gin.H{
		"data":          articleData,
		"categorys":     categorys,
		"tags":          tags,
		"articleStatus": model.ArticleStatus,
	})
}
func Del(c *gin.Context) {
	id := c.PostForm("id")
	article_id, _ := strconv.Atoi(id)
	data := model.DelArticle(article_id)
	c.JSON(http.StatusOK, gin.H{"code": data.Code, "msg": data.Msg, "url": data.Url})
}

func UpdateData(c *gin.Context) {
	var article model.Article
	err := c.ShouldBind(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "msg": err.Error()})
	}
	session := sessions.Default(c)
	article.Username = session.Get("username").(string)
	data := model.UpdateArticle(&article)
	c.JSON(http.StatusOK, gin.H{"code": data.Code, "msg": data.Msg, "url": data.Url})
}
