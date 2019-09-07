package controllers

import (
	"beego/web/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}
func (c *RegisterController) Get() {
	/*
	//1.有ORM对象
	//查询的对象
	//3.指定查询对象字段值
	//4.查询
	//-----------
	//更新步骤
	//1.要有ORM对象
	o := orm.NewOrm()
	//2.需要更新的结构体
	user := models.User{}
	//3.查到需要更新的数据
	user.Id =1
	err := o.Read(&user)
	//4.给数据重新赋值
	if err == nil {
		user.Name = "xxx"
		user.Pwd = "333"
		//5.更新
		_,err = o.Update(&user)
		if err != nil {
			beego.Info("更新失败",err)
			return
		}
	}
	// 删除操作
	// 1.有ORM对象
	o = orm.NewOrm()
	//2.删除的对象
	user =  models.User{}
	//3.指定删除那一条数据
	user.Id = 1
	//4.删除
	_, err =o.Delete(&user)
	if err != nil {
		beego.Info("删除失败", err)
		
	}

	 */
	c.Data["title"] ="欢迎注册 aicken-wang blog!"
	c.TplName = "register.html"
}
func (c *RegisterController) Post() {
	//1、有ORM 对象
	//2.有一个要插入数据的结构体对象
	//3.对结构体赋值
	//4.插入
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("注册" ,username ,password)
	//对数据进行校验
	if username != "aicken" && password !="123" {
		beego.Info("密码校验错误")
		//c.Ctx.WriteString("户名或密码错误！")
		//指定试图文件，同时可以给这个视图传递一些数据
		c.Data["errMsg"] = "户名或密码格式错误！"
		c.TplName = "register.html"
		return
	}
	//插入数据
	o := orm.NewOrm()
	user := models.User{}
	user.Name = username
	user.Pwd = password
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("注册失败")
		c.Redirect("/register",302)
		c.Ctx.WriteString("注册失败！")
		return
	}
	//返回登陆界面
	c.Ctx.WriteString("注册成功！")
	c.Data["title"] ="注册成功 aicken-wang blog!"
	//重定向 跳转 ，不能传递数据
	c.Redirect("/",302)
}