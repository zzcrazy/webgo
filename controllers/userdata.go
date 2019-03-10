package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"webgo/library"
	"webgo/models"
)

type UserController struct {
	beego.Controller
}
func (c *UserController) Del(){
	//idStr := c.Ctx.Input.Param(":id")
	id,err:=c.GetInt("id")
	fmt.Println("get",id)
	u:=models.Userinfo{}
	u.Id =int64(id)
	ret ,err:=models.Del(&u)
	var resp interface{}
	resp=library.Respmsg(0,"","")
	if ret==0{
		resp= library.Respmsg(1001,"del fail","")
		fmt.Println(err,ret)
	}
	c.Data["json"]=resp
	c.ServeJSON()
}
func (c *UserController)Adduser(){
	u :=new(models.Userinfo)
	u.Username=c.GetString("username")
	age,err:=c.GetInt("age")
	if err!=nil{
		fmt.Println(err)
		return
	}
	u.Age=age
	aged :=c.Input().Get("age")
	fmt.Println(aged)
	res ,err:=models.Insertuser(u)

	var resp interface{}
	resp=library.Respmsg(0,"","")
	if res==0{
		resp= library.Respmsg(1001," add fil","")
	}
	c.Data["json"]=resp
	c.ServeJSON()

}