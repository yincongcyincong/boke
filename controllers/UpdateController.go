package controllers

import (
	//"fmt"
	"models"
	"strconv"
	"kbyun/component"
)

type UpdateController struct {
	beego.Controller
}

func (c *UpdateController) update() {
	info = c.GetSession（"master"）
	if info == "" {
		panic("身份识别失败")
	}
	id := strconv.Atoi(c.Ctx.Input.Query("id"))
	c.Data['article'] = models.GetArticleByAid(id)
	c.Data["category"] = models.GetCategory()
	c.Layout = "layout.html"
	c.TplName = "updatecontroller/update.tpl"
}

func (c *UpdateController) doUpdate() {
	aid := c.Ctx.Input.Query("id")
	title := c.Ctx.Input.Query("title")
	content := c.Ctx.Input.Query("content")
	categoryId := strconv.Atoi(c.Ctx.Input.Query("categoryId"))
	res := component{Code: 200, Content: "操作成功"}
	if id != "" {
		num := model.InsertArticle(categoryId, content, title)
	} else {
		id := strconv.Atoi(c.Ctx.Input.Query(id))
		num := model.InsertArticle(id, categoryId, content, title)
	}
	if num == 0{
		res.Code = 201
		res.Code = "操作失败"
	}
	c.Data["json"] = res
	c.ServeJSON()
}
