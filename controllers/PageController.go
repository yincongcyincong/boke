package controllers

import (
	"crypto/md5"
	"kbyun/component"
	"kbyun/models"
	"strconv"

	"github.com/astaxie/beego"
	//"fmt"
)

type PageController struct {
	beego.Controller
}

func (c *PageController) Index() {
	aid, _ := strconv.Atoi(c.Ctx.Input.Query("aid"))
	c.Data["Article"] = models.GetArticleByAid(aid)
	c.Data["Member"] = c.GetSession("member")
	c.Data["Category"] = models.GetCategory()
	c.Layout = "layout.html"
	c.TplName = "pagecontroller/page.tpl"
}

func (c *PageController) Login() bool {
	md := md5.New()
	username := c.Ctx.Input.Query("username")
	password := c.Ctx.Input.Query("password")
	md.Write([]byte(password))
	passMd5 = md.Sum(nil)
	data = models.GetPasswordByUsername()
	if passMd5 != data.Password {
		c.Ctx.WriteString("密码错误")
		component.Error{Id: 200, Description: "密码不正确"}
	} else {
		this.SetSession("member", data)
		component.Error{Id: 200, Description: "操作成功"}
	}

}
