package model

import (
	"myWeb/lib/mysql"
	"strconv"
)

type ArticleTags struct {
	Id        int
	TagId     int
	ArticleId int
	CreatedAt int64
	UpdatedAt int64
}

func TagsAtricle(tag_id string) []string {
	var tags []ArticleTags
	var article_ids []string
	mysql.MysqlConnect.Select("article_id").Where("tag_id = ?", tag_id).Find(&tags)
	for _, v := range tags {
		Ids := strconv.Itoa(v.ArticleId)
		article_ids = append(article_ids, Ids)
	}
	return article_ids
}

func GetArticleTagsIds(article_id int) map[int]string {
	var tags []ArticleTags
	var tags_ids = make(map[int]string)
	mysql.MysqlConnect.Select("article_id,tag_id").Where("article_id = ?", article_id).Find(&tags)
	for _, v := range tags {
		tags_ids[v.TagId] = strconv.Itoa(v.ArticleId)
	}
	return tags_ids
}
