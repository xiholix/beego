package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (main *MainController) HelloSitePoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "123872842@qq.com"
	main.TplName = "hello-sitepoint.tpl"
}
