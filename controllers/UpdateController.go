package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"kbyun/models"
	"strconv"
	"kbyun/component"
)

type UpdateController struct {
	beego.Controller
}

func (c *UpdateController) Update() {
	info := c.GetSession("master")
	if info == "" {
		panic("身份识别失败")
	}
	id, _ := strconv.Atoi(c.Ctx.Input.Query("id"))
	c.Data["article"] = models.GetArticleByAid(id)
	c.Data["category"] = models.GetCategory()
	c.Layout = "layout.html"
	c.TplName = "updatecontroller/update.tpl"
}

func (c *UpdateController) doUpdate() {
	aid := c.Ctx.Input.Query("id")
	title := c.Ctx.Input.Query("title")
	content := c.Ctx.Input.Query("content")
	categoryId, _ := strconv.Atoi(c.Ctx.Input.Query("categoryId"))
	res := component.Error{Code: 200, Content: "操作成功"}
	var num int64
	if aid != "" {
		num = models.InsertArticle(categoryId, content, title)
	} else {
		id, _ := strconv.Atoi(c.Ctx.Input.Query(aid))
		num = models.UpdateArticle(id, categoryId, content, title)
	}
	if num == 0{
		res.Code = 201
		res.Content = "操作失败"
	}
	c.Data["json"] = res
	c.ServeJSON()
}
