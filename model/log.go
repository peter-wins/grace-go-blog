package model

import (
	"myWeb/lib/mysql"
	"time"
)

type Log struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"create_time"`
}

func GetLogList(log *Log, page *OptPage) (interface{}, int) {
	var SearchData []Log           //查询数据存储
	var returnData [][]interface{} //返回数据
	Db := mysql.MysqlConnect
	if log.Username != "" {
		Db = Db.Where("username like ?", "%"+log.Username+"%")
	}
	Db.Offset(page.Start).Limit(page.Length).Find(&SearchData)
	if len(SearchData) > 0 {
		var total int
		Db.Model(&log).Count(&total)
		for _, v := range SearchData {
			created_at := time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05") //开服时间
			tmp := []interface{}{v.Id, v.Username, v.Content, created_at}
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
func (l *Log) InsertLog(log Log) {
	created_at := time.Now().Unix()
	log = Log{Username: log.Username, Content: log.Content, CreatedAt: created_at}
	mysql.MysqlConnect.Create(&log)
}
