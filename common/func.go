package common

import (
	"crypto/md5"
	"fmt"
	"goadmin/models"
	"strconv"
)

type Record struct {
	Rows interface{} `json:"rows"`
	Total int `json:"total"`
}

var UserMenu []*models.Menu
func Check(Url string,Uid int,Mode int) bool{
	for _,v := range UserMenu {
		fmt.Println("====》",v)
		if v.Url == Url {
			return true
		}
	}
	return false
}
//MD5
func Md5(str string) string{
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
//字符串转整型
func StringToInt(str string) int  {
	strint,err :=  strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return strint
}
//整型转字符串
func IntToString(i int) string {
	str := strconv.Itoa(i)
	return str
}