package main

import (
	_ "beego_damo02/db_mysql"
	_ "beego_damo02/routers"
	"github.com/astaxie/beego"
)



func main() {

	//appName := config.String("appname")
	//fmt.Println("应用名称:", appName)
	//port, err := config.Int("httpport")
	//if err != nil {
	//	panic("项目配置文件解析失败，请检查配置文件")
	//	//panic:使系统进入恐慌状态
	//}

	beego.Run()
}
