package main

import (
	"fmt"
	"myWeb/app/admin"
	"myWeb/app/index"
	"myWeb/lib/mysql"
	"myWeb/routers"
)

func main() {
	//链接数据库
	err := mysql.InitMysql()
	if err != nil {
		panic(err)
	}
	defer mysql.MysqlConnect.Close()
	// 加载多个APP的路由配置
	routers.Include(index.Routers, admin.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(":8800", ); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

}
