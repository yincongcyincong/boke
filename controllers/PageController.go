package controllers

import(
        "kbyun/models"
        "github.com/astaxie/beego"
	"strconv"
)

type PageController struct {
        beego.Controller
}

func (c *PageController) Index(){
        aid,_ := strconv.Atoi(c.Ctx.Input.Query("aid"))
        c.Data["Article"] = models.GetArticleByAid(aid)
        c.Layout = "layout.html"
        c.TplName = "pagecontroller/page.tpl"
}

