package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}
type IndexController struct {
	beego.Controller
}
func (c *IndexController) Get() {
	//c.Data["Website"] = "wangshunqing.work"
	//c.Data["Email"] = "aicken@gmail.com"
	//c.TplName = "index.tpl"
	c.Ctx.WriteString("Hello World!")
}