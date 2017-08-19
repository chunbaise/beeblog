package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
	}
	c.Data["IsCategory"] = true

	c.TplName = "category.html"
}
