package routers

import (
	"webgo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/detail/?:id", &controllers.MainController{})
	//beego.AutoRouter(&controllers.UserController{})
	beego.Router("/pa", &controllers.MainController{}, "Get:Dologin")
	beego.Router("/userlist", &controllers.UserController{}, "Get:Userlist")
	beego.Router("/del?:id", &controllers.UserController{}, "Get:Del")
	beego.Router("/adduser", &controllers.UserController{}, "Post:Adduser")

	beego.Router("/do", &controllers.MainController{})

}
