package models

import (
	"github.com/astaxie/beego/orm"
	"kbyun/component"

	"strconv"
)

type Article struct {
	Id         int       `orm"pk;auto"`
	Category   *Category `orm:"rel(fk)"`
	Title	   string`orm:"size(255)"`
	Ctime      string
	BrowseTime int    `orm:"size(32);default(0)"`
	Content    string `orm:"type(text)"`
}

func GetArticleByCategory(category *Category) []*Article {
	o := orm.NewOrm()
	var article Article
	var articles []*Article
	o.QueryTable(article).RelatedSel().Filter("Category", category).All(&articles)
	return articles
}

func GetArticle(size, p int, articles *[]*Article) component.Pages  {
	o := orm.NewOrm()
	var article Article
	source := o.QueryTable(article)
	count,_ := source.Limit(-1).Count()
	source.OrderBy("Ctime").Limit(size).Offset((p-1)*size).All(articles)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return component.Page(c, p, size)
}

func GetArticleByAid(aid int) Article {
	o := orm.NewOrm()
	article := Article{Id: aid}
	var data Article
	o.QueryTable(article).One(&data)
	return data
}
