package controllers

import (
	"beego_damo02/db_mysql"
	"beego_damo02/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	bodyBytes,err:=ioutil.ReadAll(r.Ctx.Request.Body)
	if err!=nil {
		r.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	var user models.User
	err=json.Unmarshal(bodyBytes,&user)
	if err !=nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析错误")
		return
	}
	id,err:=db_mysql.InserUser(user)
	if err!=nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存错误")
		return
	}
	fmt.Println(id)
	result:=models.ResponseResult{
		Code:    0,
		Message: "保存成功",
		Data:    nil,
	}
	r.Data["json"]=&result
	r.ServeJSON()
}
