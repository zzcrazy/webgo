package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"goodsdetail/models"
	_"goodsdetail/models"
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
	c.Data["data"]="wuxiaoyu"
	c.Data["Photos"] = models.Getphoto()
	c.Data["Title"] =models.GetTitle()

}
func (c *MainController)Post(){
	u :=new(models.Userinfo)
	u.Age=11
	u.Username="wangw"
	res ,err:=models.Insertuser(u)
	fmt.Println(res,err)
	c.TplName = "test.html"
	c.Data["data"]="liuxinran"

}
func (c *MainController)Dologin(){
	fmt.Println("ppapacao")
	c.TplName="test.html"
}
func (c *MainController)Userlist(){
	u:=new(models.Userinfo)
	conds :=map[string]string{"username":"wangw"}

	res ,err:=models.Selectuser(u,0,0,conds,"")

	fmt.Println(res ,err)


	c.Data["json"]=res
	c.ServeJSON()
	//c.TplName = "test.html"
	//c.Data["data"]="liuxinran"

}