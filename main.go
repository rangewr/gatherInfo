package main

import (
	_ "gatherInfo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}


//初始化orm,并设置数据库连接
func init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	connectionString := "root:123456@tcp(localhost:3306)/checkinfo?charset=utf8"
	orm.RegisterDataBase("default","mysql",connectionString)
	//orm.RegisterDataBase("default","mysql","admin:Hecha1809@tcp(222.74.254.17:1366/viewonline?charset=utf8")
	orm.Debug=true
	//orm.SetMaxIdleConns("default",100)
	//orm.SetMaxOpenConns("default",100)
}