package model

import (
	"myWeb/lib/mysql"
	"strconv"
	"time"
)

type Category struct {
	Id        int    `form:"id"`
	Name      string `form:"name"`
	Remark    string `form:"remark"`
	Status    int    `form:"status"`
	Ctype     int    `form:"c_type"`
	CreatedAt int64
	UpdatedAt int64
}

var CategoryStatus map[int]string = map[int]string{1: "是", 2: "否"}
var Ctypes map[int]string = map[int]string{1: "顶级分类", 2: "系列分类"}

func GetCategoryList(category *Category, page *OptPage) (interface{}, int) {
	var SearchData []Category      //查询数据存储
	var returnData [][]interface{} //返回数据
	Db := mysql.MysqlConnect
	if category.Name != "" {
		Db = Db.Where("name like ?", "%"+category.Name+"%")
	}
	Db.Offset(page.Start).Limit(page.Length).Find(&SearchData)
	if len(SearchData) > 0 {
		var total int
		Db.Model(&category).Count(&total)
		for _, v := range SearchData {
			option := "<a href='edit?id=" + strconv.Itoa(v.Id) + "' class='btn btn-primary btn-sm' >编辑</a>"
			option += "<a href='javascript:;' class='btn btn-danger btn-sm' onclick='del(" + strconv.Itoa(v.Id) + ")'>删除</a>"
			created_at := time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05") //开服时间
			updated_at := time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05") //开服时间
			tmp := []interface{}{v.Id, v.Name, v.Remark, v.Status, v.Ctype, created_at, updated_at, option}
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
func CheckCategoryField(field string, data string) (res Category) {
	mysql.MysqlConnect.Where(field+" = ?", data).First(&res)
	return
}

/*
 * 记录管理员操作日志信息
 */
func UpdateCategory(category *Category) Optreturn {

	if category.Id != 0 {
		category.UpdatedAt = time.Now().Unix()
		if mysql.MysqlConnect.Model(&category).Update(category).Error != nil {
			return Optreturn{Code: 0, Msg: "修改失败"}
		}
	} else {
		data := CheckCategoryField("name", category.Name)
		if data.Name != "" {
			return Optreturn{Code: 0, Msg: "名称已存在"}
		}
		data = CheckCategoryField("remark", category.Remark)
		if data.Remark != "" {
			return Optreturn{Code: 0, Msg: "标识已存在"}
		}
		category.CreatedAt = time.Now().Unix()
		if err := mysql.MysqlConnect.Create(&category).Error; err != nil {
			return Optreturn{Code: 0, Msg: "添加失败"}
		}
	}
	return Optreturn{Code: 1, Msg: "添加成功", Url: "/admin/category/index"}
}

func GetCategoryInfo(id int) (data Category) {
	mysql.MysqlConnect.Where("id =? ", id).First(&data)
	return
}

func GetCategorys() map[int]string {
	var category []Category
	var categoryLists = make(map[int]string)
	mysql.MysqlConnect.Select("id,name").Find(&category)
	for _, v := range category {
		categoryLists[v.Id] = v.Name
	}
	return categoryLists
}
func ApiCategorys(ctype string) (map[int]string, map[int]string) {
	var category []Category
	var categoryTops = make(map[int]string)
	var categoryLefts = make(map[int]string)
	Db := mysql.MysqlConnect
	if ctype != "" {
		Db = Db.Where("ctype =?", ctype)
	}
	Db.Select("id,ctype,name").Find(&category)
	for _, v := range category {
		if v.Ctype == 1 {
			categoryTops[v.Id] = v.Name
		} else {
			categoryLefts[v.Id] = v.Name
		}

	}
	return categoryTops, categoryLefts
}
