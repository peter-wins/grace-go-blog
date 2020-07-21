package model

import (
	"myWeb/lib/mysql"
	"strconv"
	"time"
)

type Tags struct {
	Id        int    `form:"id"`
	Name      string `form:"name"`
	Status    int    `form:"status"`
	CreatedAt int64
	UpdatedAt int64
}

var TagsStatus map[int]string = map[int]string{1: "是", 2: "否"}

func GetTagsList(tags *Tags, page *OptPage) (interface{}, int) {
	var SearchData []Tags          //查询数据存储
	var returnData [][]interface{} //返回数据
	Db := mysql.MysqlConnect
	if tags.Name != "" {
		Db = Db.Where("name like ?", "%"+tags.Name+"%")
	}
	Db.Offset(page.Start).Limit(page.Length).Find(&SearchData)
	if len(SearchData) > 0 {
		var total int
		Db.Model(&tags).Count(&total)
		for _, v := range SearchData {
			option := "<a href='edit?id=" + strconv.Itoa(v.Id) + "' class='btn btn-primary btn-sm' >编辑</a>"
			option += "<a href='javascript:;' class='btn btn-danger btn-sm' onclick='del(" + strconv.Itoa(v.Id) + ")'>删除</a>"
			created_at := time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05") //开服时间
			updated_at := time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05") //开服时间
			tmp := []interface{}{v.Id, v.Name, v.Status, created_at, updated_at, option}
			returnData = append(returnData, tmp)
		}

		return returnData, total
	} else {
		return SearchData, 0
	}
}

/**
*查看是否存在某字段
 */
func CheckTagsField(field string, data string) (res Tags) {
	mysql.MysqlConnect.Where(field+" = ?", data).First(&res)
	return
}

/*
 * 记录管理员操作日志信息
 */
func UpdateTags(tags *Tags) Optreturn {

	if tags.Id != 0 {
		tags.UpdatedAt = time.Now().Unix()
		if mysql.MysqlConnect.Model(&tags).Update(tags).Error != nil {
			return Optreturn{Code: 0, Msg: "修改失败"}
		}
	} else {
		data := CheckTagsField("name", tags.Name)
		if data.Name != "" {
			return Optreturn{Code: 0, Msg: "名称已存在"}
		}
		tags.CreatedAt = time.Now().Unix()
		if err := mysql.MysqlConnect.Create(&tags).Error; err != nil {
			return Optreturn{Code: 0, Msg: "添加失败"}
		}
	}
	return Optreturn{Code: 1, Msg: "添加成功", Url: "/admin/tags/index"}
}

func GetTagsInfo(id int) (data Tags) {
	mysql.MysqlConnect.Where("id =? ", id).First(&data)
	return
}

func GetTags() map[int]string {
	var tags []Tags
	var tagsLists = make(map[int]string)
	mysql.MysqlConnect.Select("id,name").Find(&tags)
	for _, v := range tags {
		tagsLists[v.Id] = v.Name
	}
	return tagsLists
}

func ApiTags() map[int]string {
	var tags []Tags
	var tagsLists = make(map[int]string)
	Db := mysql.MysqlConnect

	Db.Select("id,name").Find(&tags)
	for _, v := range tags {
		tagsLists[v.Id] = v.Name
	}
	return tagsLists
}
