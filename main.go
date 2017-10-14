package main

import (
	_ "kbyun/routers"
	"kbyun/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@/boke?charset=utf8")
	orm.RegisterModel(
		new(models.Users),
		new(models.Category),
		new(models.Article),
	)
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
