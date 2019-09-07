package main

import (
	_ "beego/web/routers"
	_ "beego/web/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

