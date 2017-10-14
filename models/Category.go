package models

import(
	"github.com/astaxie/beego/orm"	
)

type Category struct {
	Id       int    `orm:"pk;auto"`
	TypeName string `orm"size(60)"`
	Ctime    string
}

func GetCategory() []*Category {
	o := orm.NewOrm()
	var category Category
	var categorys []*Category
	o.QueryTable(category).All(&categorys)
	return categorys
}
