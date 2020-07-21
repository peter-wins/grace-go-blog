package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myWeb/model"
	"net/http"
	"time"
)

type PostParams struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/index.html", gin.H{})
}
func Login(c *gin.Context) {
	var postParams PostParams
	status, msg := 1, "非法参数"
	//获取页面参数
	err := c.ShouldBind(&postParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 0,
			"msg":    err.Error(),
		})
	}
	//账号与数据表里对比是否存在
	data := model.CheckUser(postParams.Username)
	if data.Username == "" {
		status = 0
		msg = "账号不存在"
	} else {
		//密码与数据表里对比
		if data.Password != model.ConvertPassword(postParams.Password, data.Salt) {
			status = 0
			msg = "密码错误"
		}
	}
	if status == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": status,
			"msg":    msg,
		})
	} else {
		//查询对应组
		session := sessions.Default(c)
		session.Set("user_id", data.Username)
		session.Set("username", data.Username)
		session.Save()

		//更新管理员状态
		var updateUser model.UpdateUser
		updateUser.Id = data.Id
		updateUser.Logintime = int(time.Now().Unix())
		updateUser.Loginip = model.GetLocalIp()

		new(model.UpdateUser).UpdateUser(updateUser)
		var log model.Log
		log.Username = session.Get("username").(string)
		log.Content = "用户登录"
		new(model.Log).InsertLog(log)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "登录成功",
			"url":    "/admin",
		})
	}
}

/*
 * 登出
 */
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}
