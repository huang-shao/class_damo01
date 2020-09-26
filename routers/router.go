package routers

import (
	"beego_damo02/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.RegisterController{})
}
//register:注册