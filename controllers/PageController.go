package controllers

import (
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
	c.Layout = "layout.html"
	c.TplName = "pagecontroller/page.tpl"
}
