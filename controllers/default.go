package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"webgo/library"

	"webgo/models"
	_"webgo/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	id := c.Ctx.Input.Param(":id")
	fmt.Println(id)
	c.TplName = "test.html"
	c.Data["Photos"] = models.Getphoto()
	c.Data["Title"] =models.GetTitle()

}

func (c *MainController)Dologin(){
	fmt.Println("ppapacao")
	c.TplName="test.html"
}
func (c *MainController)Userlist(){
	u:=new(models.Userinfo)
	//conds :=map[string]string{"username":"wangw"}
	conds:=map[string]string{}
	res ,err:=models.Selectuser(u,0,0,conds,"")

	//fmt.Println(res ,err)
	if err!=nil{
		fmt.Println("err ",err)
	}
	var resp interface{}
	resp=library.Respmsg(0,"",res)

	if res==0{
		resp= library.Respmsg(1001,"del fail","")
	}
	c.Data["json"]=resp
	c.ServeJSON()

}