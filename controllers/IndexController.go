package controllers

import (
	"fmt"
	"kbyun/models"
	"regexp"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	id := c.Ctx.Input.Query("cid")
	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	c.Data["S"] = p
	if p == 0 {
		p = 1
	}
	var articles []*models.Article
	if id != "" {
		cid, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		category := models.Category{Id: cid}
		c.Data["page"] = models.GetArticleByCategory(&category, &articles)
	} else {
		c.Data["page"] = models.GetArticle(1, p, &articles)
	}
	fmt.Println(c.Data["page"])
	c.Data["category"] = models.GetCategory()
	for index, data := range articles {
		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		data.Content = re.ReplaceAllString(data.Content, "")
		data.Content = data.Content[0:300]
		ctime, _ := strconv.ParseInt(data.Ctime, 10, 64)
		data.Ctime = time.Unix(ctime, 0).Format("2006-01-02 15:04:05")
		articles[index] = data
	}
	c.Data["article"] = articles
	if c.GetSession("master") == nil || c.GetSession("master") != "master" {
		c.Data["isMaster"] = false
	} else {
		c.Data["isMaster"] = true
	}
	c.Layout = "layout.html"
	c.TplName = "indexcontroller/index.tpl"
}
