package models

import (
	"kbyun/component"

	"github.com/astaxie/beego/orm"

	"strconv"
)

type Article struct {
	Id         int       `orm"pk;auto"`
	Category   *Category `orm:"rel(fk)"`
	Title      string    `orm:"size(255)"`
	Ctime      string
	BrowseTime int    `orm:"size(32);default(0)"`
	Content    string `orm:"type(text)"`
}

func GetArticleByCategory(category *Category, articles *[]*Article, size, p int) component.Pages {
	o := orm.NewOrm()
	var article Article
	source := o.QueryTable(article)
	count, _ := source.Limit(-1).Filter("Category", category).Count()
	source.RelatedSel().OrderBy("Ctime").Filter("Category", category).Limit(size).Offset((p - 1) * size).All(articles)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return component.Page(c, p, size)
}

func GetArticle(size, p int, articles *[]*Article, keyword string) component.Pages {
	o := orm.NewOrm()
	var article Article
	source := o.QueryTable(article)
	count, _ := source.Filter("title__icontains", keyword).Limit(-1).Count()
	source.OrderBy("Ctime").Filter("title__icontains", keyword).Limit(size).Offset((p - 1) * size).All(articles)
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

func UpdateArticle(id, categoryId int, title, content string) int64 {
	o := orm.NewOrm()
	res, err := o.Raw("update article set category_id = ?, content = ?, title = ? where id = ?", categoryId, content, title).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num
	} else {
		return 0
	}
}

func InsertArticle(categoryId int, title, content string) int64 {
	o := orm.NewOrm()
	res, err := o.Raw("insert into article(category_id, title, content) value(?, ?, ?)", categoryId, title, content).Exec()
	if err == nil {
                num, _ := res.RowsAffected()
		return num
        } else {
                return 0
        }
}

