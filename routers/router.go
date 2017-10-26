package routers

import (
	"kbyun/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ueditor", &controllers.UeditorController{}, "*:ControllerUE")
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/page", &controllers.PageController{}, "get:Index")
	beego.Router("/login", &controllers.PageController{}, "post:Login")
	beego.Router("/update", &controllers.UpdateController{}, "get:Update")
	beego.Router("/doUpdate", &controllers.UpdateController{}, "post:DoUpdate")
	//beego.Router("/index/del", &controllers.IndexController{}, "post:del")
}
