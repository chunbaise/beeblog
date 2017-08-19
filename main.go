package main

import (
	"beeblog/controllers"
	"beeblog/models"
	//_ "beeblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	// 开启ORM 调试模式
	orm.Debug = true
	// 自动创建表
	orm.RunSyncdb("default", false, true)

	// 注册路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/login", &controllers.LoginController{})

	// 启动beego
	beego.Run()
}
