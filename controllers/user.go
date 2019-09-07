package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController)GetURI()  {
	splat := c.Ctx.Input.Param(":username")
	c.Ctx.WriteString("URI " + splat )
}