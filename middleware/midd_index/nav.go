package midd_index

import (
	"github.com/gin-gonic/gin"
	"myWeb/model"
)

func HeaderAndRightNav(c *gin.Context) {
	categoryTops, categoryLefts := model.ApiCategorys("")
	c.Set("categoryTops", categoryTops)
	c.Set("categoryLefts", categoryLefts)
	c.Next()
}

func TagsNav(c *gin.Context) {
	tags := model.ApiTags()
	c.Set("tags", tags)
	c.Next()
}
