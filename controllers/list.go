package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"wind_wallpaper/models"
)

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *ListController) Index()  {
	query := &models.Search{Categories: "010", Sorting: "random", Purity:"100", Order:"desc", Page: 2}
	str := query.SearchQuery()
	response := models.Get(str)
	content, page := models.List(response)
	c.Data["Str"] = content
	c.Data["Page"] = page
	fmt.Println(page)
	c.TplName = "list/list.tpl"
}
