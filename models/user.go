package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
//表的设计  ORM 会 for  range
type User struct {
	Id int
	Name string
	Pwd string
}

func  init()  {
	// 设置数据库基本信息
	err := orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败！",err)
	}
	// 映射model数据,new(Article),new(ArticleType)
	orm.RegisterModel(new(User))
	// 生成表
	orm.RunSyncdb("default", false, true)
}