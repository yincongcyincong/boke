package controllers

import (
	"crypto/md5"
	"kbyun/component"
	"kbyun/models"
	"strconv"
	"encoding/hex"
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
	c.Data["Master"] = c.GetSession("master")
	c.Data["Category"] = models.GetCategory()
	c.Layout = "layout.html"
	c.TplName = "pagecontroller/page.tpl"
}

func (c *PageController) Login() {
	md := md5.New()
	username := c.Ctx.Input.Query("username")
	password := c.Ctx.Input.Query("password")
	md.Write([]byte(password))
	passMd5 := md.Sum(nil)
	data := models.GetPasswordByUsername(username)
	result := component.Error{Code: 200, Content: "操作成功"}
	if hex.EncodeToString(passMd5) != data.Password {
		result.Code = 201
		result.Content = "密码错误"
	} else {
		if data.Role == 1 {
			c.SetSession("master", data)
		}
		c.SetSession("member", data)		
	}
	c.Data["json"] = result
	c.ServeJSON()
}
