package models

//	"github.com/astaxie/beego/orm"

type Users struct {
	Uid      int    `orm:"pk;auto"`
	Username string `orm:"unique"`
	Password string `orm:"size(60)"`
	Nickname string `orm:"size(32)"`
	Role     int    `orm:"size(1)"`
}
