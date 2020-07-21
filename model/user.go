package model

import (
	"myWeb/lib/mysql"
)

type User struct {
	Id        int
	Username  string
	Password  string
	Salt      string
	Logintime int
	Loginip   string
}
type UpdateUser struct {
	Id        int
	Logintime int
	Loginip   string
}

func CheckUser(username string) (User User) {
	mysql.MysqlConnect.Table("zcr_user").
		Where("username = ?", username).
		First(&User)
	return
}

func (user *UpdateUser) UpdateUser(updateUser UpdateUser) {
	mysql.MysqlConnect.Table("zcr_user").
		Model(user).
		Where("id = ?", updateUser.Id).
		Updates(map[string]interface{}{"loginip": updateUser.Loginip, "logintime": updateUser.Logintime})
}
