package routers

import (
	"webgo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/detail/?:id", &controllers.MainController{})
	//beego.AutoRouter(&controllers.UserController{})
	beego.Router("/pa", &controllers.MainController{}, "Get:Dologin")
	beego.Router("/userlist", &controllers.MainController{}, "Get:Userlist")

	beego.Router("/do", &controllers.MainController{})

}
