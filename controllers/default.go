package controllers

import (
	"beego_damo02/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//获取请求数据
	user:=c.Ctx.Input.Query("user")
	pwd:=c.Ctx.Input.Query("pwd")
	//使用固定数据进行数据校验
	if user!="hxl"||pwd!="123456" {
		c.Ctx.ResponseWriter.Write([]byte("对不起,数据错误！"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据校验成功"))
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
//func(c *MainController)Post(){
//	//接收post请求的参数
//	name:=c.Ctx.Request.FormValue("name")
//	age:=c.Ctx.Request.FormValue("age")
//	sex:=c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,age,sex)
//	if name!="hxl"&&age!="18" {
//		c.Ctx.WriteString("数据校验失败")
//		return
//	}
//	c.Ctx.WriteString("成功")
//}
func(c *MainController)Delete(){


}
func (c *MainController)Post()  {
	var person models.Person
	dateBytes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err!=nil {
		c.Ctx.WriteString("数据接收错误！")
		return
	}
	err=json.Unmarshal(dateBytes,&person)
	if err!=nil{
		c.Ctx.WriteString("数据解析错误")
		return
	}
	fmt.Println("姓名:",person.Name)
	fmt.Println("年龄:",person.Age)
	fmt.Println("性别：",person.Sex)
	c.Ctx.WriteString("数据解析成功")
}