package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}
func (c *UserController) List(){
	//idStr := c.Ctx.Input.Param(":id")
	fmt.Println("12323213")

}