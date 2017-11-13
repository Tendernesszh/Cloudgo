package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//If no arguments are passed in, output Hello World!
func (c *MainController) Get() {
    c.Ctx.WriteString("Hello World!\n")
}

//Establish UserController
type UserController struct {
    beego.Controller
}

//Handle incoming address with name arguments to output username
func (c *UserController) Get() {
    username := c.Ctx.Input.Param(":name")
		c.Ctx.WriteString("Hello " + username + "!\n")
}
