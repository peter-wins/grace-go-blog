package model

import (
	"myWeb/lib/mysql"
	"time"
)

type ArticleIp struct {
	Id        int
	ArticleId int
	Ip        string
	CreatedAt int64
	UpdatedAt int64
}

func GetArticleIpList(article_ip *ArticleIp, page *OptPage) (interface{}, int) {
	var SearchData []ArticleIp           //查询数据存储
	var returnData [][]interface{} //返回数据
	Db := mysql.MysqlConnect
	if article_ip.Ip != "" {
		Db = Db.Where("ip like ?", "%"+article_ip.Ip+"%")
	}
	Db.Offset(page.Start).Limit(page.Length).Find(&SearchData)
	if len(SearchData) > 0 {
		var total int
		Db.Model(&article_ip).Count(&total)
		for _, v := range SearchData {
			created_at := time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05") //开服时间
			tmp := []interface{}{v.Id, v.Ip, v.ArticleId, created_at}
			returnData = append(returnData, tmp)
		}

		return returnData, total
	} else {
		return SearchData, 0
	}
}


func InsertArticleIp(article_ip *ArticleIp) {
	mysql.MysqlConnect.Create(&article_ip)
	return
}
