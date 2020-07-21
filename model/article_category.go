package model

import (
	"myWeb/lib/mysql"
	"strconv"
)

type ArticleCategory struct {
	Id         int
	CategoryId int
	ArticleId  int
	CreatedAt  int64
	UpdatedAt  int64
}

func CategoryArticle(category_id string) []string {
	var categorys []ArticleCategory
	var article_ids []string
	mysql.MysqlConnect.Select("article_id").Where("category_id = ?", category_id).Find(&categorys)
	for _, v := range categorys {
		Ids := strconv.Itoa(v.ArticleId)
		article_ids = append(article_ids, Ids)
	}
	return article_ids
}

func GetArticleCategoryIds(article_id int) map[int]string  {
	var categorys []ArticleCategory
	var category_ids  =  make(map[int]string)
	mysql.MysqlConnect.Select("article_id,category_id").Where("article_id = ?", article_id).Find(&categorys)
	for _, v := range categorys {
		category_ids[v.CategoryId] = strconv.Itoa(v.ArticleId)
	}
	return category_ids
}
