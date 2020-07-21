package model

import (
	"myWeb/lib/mysql"
	"strconv"
	"time"
)

type Article struct {
	Id        int      `form:"id"`
	Username  string   `json:"username"`
	Title     string   `form:"title"`
	Content   string   `form:"content"`
	Status    int      `form:"status"`
	Sorts     int      `form:"sort"`
	Tags      []string `gorm:"-" form:"tags[]"`
	Category  []string `gorm:"-" form:"category[]"`
	CreatedAt int64
	UpdatedAt int64
}

var ArticleStatus map[string]string = map[string]string{"1": "发布", "2": "草稿"}

func GetArticleList(article *Article, page *OptPage) (interface{}, int) {
	var SearchData []Article       //查询数据存储
	var returnData [][]interface{} //返回数据
	Db := mysql.MysqlConnect
	if article.Title != "" {
		Db = Db.Where("name like ?", "%"+article.Title+"%")
	}
	Db.Offset(page.Start).Limit(page.Length).Find(&SearchData)
	if len(SearchData) > 0 {
		var total int
		Db.Model(&article).Count(&total)
		for _, v := range SearchData {
			option := "<a href='edit?id=" + strconv.Itoa(v.Id) + "' class='btn btn-primary btn-sm' >编辑</a>"
			option += "<a href='javascript:;' class='btn btn-danger btn-sm' onclick='del(" + strconv.Itoa(v.Id) + ")'>删除</a>"
			created_at := time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05") //开服时间
			updated_at := time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05") //开服时间
			tmp := []interface{}{v.Id, v.Title, v.Status, created_at, updated_at, option}
			returnData = append(returnData, tmp)
		}

		return returnData, total
	} else {
		return SearchData, 0
	}
}

/*
 * 记录管理员操作日志信息
 */
func UpdateArticle(article *Article) Optreturn {

	tags_arr := article.Tags
	category_arr := article.Category
	if article.Id != 0 {
		article.UpdatedAt = time.Now().Unix()
		if mysql.MysqlConnect.Model(&article).Update(article).Error != nil {
			return Optreturn{Code: 0, Msg: "修改失败"}
		} else {
			DelArticleTags(article.Id)
			for i := 0; i < len(tags_arr); i++ {
				tags_id, _ := strconv.Atoi(tags_arr[i])
				article_tags := ArticleTags{
					UpdatedAt: time.Now().Unix(),
					CreatedAt: time.Now().Unix(),
					ArticleId: article.Id,
					TagId:     tags_id,
				}
				mysql.MysqlConnect.Create(&article_tags)
			}
			DelArticleCategory(article.Id)
			for i := 0; i < len(category_arr); i++ {
				caregory_id, _ := strconv.Atoi(category_arr[i])
				article_category := ArticleCategory{
					UpdatedAt: time.Now().Unix(),
					CreatedAt: time.Now().Unix(),
					ArticleId:  article.Id,
					CategoryId: caregory_id,
				}
				mysql.MysqlConnect.Create(&article_category)
			}
			return Optreturn{Code: 1, Msg: "操作成功", Url: "/admin/article/index"}
		}
	} else {
		article.CreatedAt = time.Now().Unix()
		result := mysql.MysqlConnect.Create(&article)
		if err := result.Error; err != nil {
			return Optreturn{Code: 0, LastID: 0, Msg: "添加失败"}
		}
		for i := 0; i < len(tags_arr); i++ {
			tags_id, _ := strconv.Atoi(tags_arr[i])
			article_tags := ArticleTags{
				UpdatedAt: time.Now().Unix(),
				CreatedAt: time.Now().Unix(),
				ArticleId: article.Id,
				TagId:     tags_id,
			}
			mysql.MysqlConnect.Create(&article_tags)
		}
		for i := 0; i < len(category_arr); i++ {
			caregory_id, _ := strconv.Atoi(category_arr[i])
			article_category := ArticleCategory{
				UpdatedAt: time.Now().Unix(),
				CreatedAt: time.Now().Unix(),
				ArticleId:  article.Id,
				CategoryId: caregory_id,
			}
			mysql.MysqlConnect.Create(&article_category)
		}
		return Optreturn{Code: 1, Msg: "操作成功", Url: "/admin/article/index"}

	}

}

func DelArticleTags(id int) {
	var article_tags ArticleTags
	mysql.MysqlConnect.Where("article_id =?", id).Delete(article_tags)
	return
}
func DelArticleCategory(id int) {
	var article_category ArticleCategory
	mysql.MysqlConnect.Where("article_id =?", id).Delete(article_category)
	return
}

func GetArticleInfo(id int) (map[string]string, interface{}, interface{}) {

	var SearchData Article //查询数据存储
	returnArticleData := make(map[string]string)

	var returnTagsData []interface{}
	var returnCategoryData []interface{}

	tagsData := GetTags()
	tempTags := GetArticleTagsIds(id)

	for k, v := range tagsData {
		tempTagsData := make(map[string]string)
		_, ok := tempTags[k]
		if ok {
			tempTagsData["checked"] = "checked"
		} else {
			tempTagsData["checked"] = ""
		}
		tempTagsData["id"] = strconv.Itoa(k)
		tempTagsData["name"] = v
		returnTagsData = append(returnTagsData, tempTagsData)
	}

	categoryData := GetCategorys()
	tempCategorys := GetArticleCategoryIds(id)
	for k, v := range categoryData {
		tempCategoryData := make(map[string]string)
		_, ok := tempCategorys[k]
		if ok {
			tempCategoryData["checked"] = "checked"
		} else {
			tempCategoryData["checked"] = ""
		}
		tempCategoryData["id"] = strconv.Itoa(k)
		tempCategoryData["name"] = v
		returnCategoryData = append(returnCategoryData, tempCategoryData)
	}

	mysql.MysqlConnect.Where("id =? ", id).First(&SearchData)
	returnArticleData["Title"] = SearchData.Title
	returnArticleData["Content"] = SearchData.Content
	returnArticleData["Id"] = strconv.Itoa(SearchData.Id)
	returnArticleData["Sorts"] = strconv.Itoa(SearchData.Sorts)
	returnArticleData["Status"] = strconv.Itoa(SearchData.Status)
	returnArticleData["Username"] = SearchData.Username
	return returnArticleData, returnCategoryData,returnTagsData
}

func DelArticle(id int) Optreturn {
	if mysql.MysqlConnect.Delete(&Article{}, "id=?", id).Error != nil {
		return Optreturn{Code: 1, Msg: "删除失败"}
	}
	DelArticleTags(id)
	DelArticleCategory(id)
	return Optreturn{Code: 0, Msg: "删除成功", Url: "/admin/article/index"}
}

func ApiArticleList(searchWhere *SearchWhere, page *OptPage) (interface{}, int) {
	var SearchData []Article     //查询数据存储
	var returnData []interface{} //返回数据
	Db := mysql.MysqlConnect
	if searchWhere.Title != "" {
		Db = Db.Where("title like ?", "%"+searchWhere.Title+"%")
	}

	if searchWhere.Cid != "" {
		article_ids := CategoryArticle(searchWhere.Cid)
		if len(article_ids) != 0 {
			Db = Db.Where("id in (?)", article_ids)
		}
	}

	if searchWhere.TagId != "" {
		article_ids := TagsAtricle(searchWhere.TagId)
		if len(article_ids) != 0 {
			Db = Db.Where("id in (?)", article_ids)
		}
	}
	var article Article
	Db.Offset(page.Start).Limit(page.Length).Find(&SearchData)

	if len(SearchData) > 0 {
		var total int
		Db.Model(&article).Count(&total)
		for _, v := range SearchData {
			tmpMap := make(map[string]string)
			created_at := time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05") //开服时间
			tmpMap["Id"] = strconv.Itoa(v.Id)
			tmpMap["Title"] = v.Title
			tmpMap["Username"] = v.Username
			tmpMap["CreatedAt"] = created_at
			returnData = append(returnData, tmpMap)
		}
		return returnData, total
	} else {
		return SearchData, 0
	}
}

func ApiArticle(Aid string) map[string]string {
	var SearchData Article //查询数据存储
	returnData := make(map[string]string)
	mysql.MysqlConnect.Where("id =? ", Aid).First(&SearchData)
	returnData["Title"] = SearchData.Title
	returnData["Content"] = SearchData.Content
	return returnData
}
