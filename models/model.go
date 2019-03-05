package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id   int
	Name string `orm:"unique"`
	Pwd  string
}

//文章结构体
type Article struct {
	Id int `orm:"pk;auto"`
	Aname string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string
	Aimg string
}

func init() {
	//设置数据库基本信息
	_ = orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/firstgo?charset=utf8")
	//映射model数据
	orm.RegisterModel(new(User),new(Article))
	//生成表
	_ = orm.RunSyncdb("default", false, true)
}
