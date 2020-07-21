package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net"
	"strconv"
	"time"
)

/*
*@code 返回码
*@msg  返回信息
*@data 返回跳转链接
 */
type Optreturn struct {
	Code   int
	Msg    string
	LastID int
	Url    string
}

/**分页*/
type OptPage struct {
	Start  int `form:"start"`
	Length int `form:"length"`
}

type SearchWhere struct {
	Cid   string
	Title string
	TagId string
}

/*
 * 处理明文密码
 * @param password 待处理的明文密码
 * @param salt 随机密匙
 * @return string 加密后的密码
 */
func ConvertPassword(password string, salt string) string {
	var pwd = Md5Func(salt + password)
	return pwd
}

/*
 * 返回MD5
 * @param str 待处理的字符串
 * @return string md5后的字符串
 */
func Md5Func(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
* 创建随机密匙
* @return string
 */
func CreateSalt() string {
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(99999999)
	return Md5Func(strconv.Itoa(rand))
}

/*
 * 生成用户密码
 */
func GetRandChar(length int) string {
	rand.Seed(time.Now().UnixNano())
	strPol := "0123456789abcdefghijklmnopqrstuvwxyz"
	max := len(strPol) - 1
	var passStr = ""
	for i := 0; i < length; i++ {
		str := string(strPol[rand.Intn(max)])
		passStr = passStr + str
	}
	return passStr
}

//获取本地ip
func GetLocalIp() (IpAddr string) {
	IpAddr = "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("Get local IP addr failed!!!")
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				IpAddr = ipnet.IP.String()
			}
		}
	}
	return
}
func MergeAssignData(c *gin.Context, h gin.H) gin.H {
	if c.Keys == nil {
		return h
	}
	if h == nil || len(h) == 0 {
		return c.Keys
	}
	mh := make(gin.H)
	for key, val := range c.Keys {
		mh[key] = val
	}
	for key, val := range h {
		mh[key] = val
	}
	return mh
}
