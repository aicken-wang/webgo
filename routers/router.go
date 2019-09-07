package routers

import (
	"beego/web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{})
    beego.Router("/register",&controllers.RegisterController{})
    //注意若有自定义方法，请求将不再访问默认的方法 多个自定义方法使用分号分割 ;
	beego.Router("/user/:username([\\w]+)", &controllers.UserController{},"*:GetURI")
}
