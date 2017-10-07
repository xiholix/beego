package controllers

import (
	"fmt"

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

func (c *MainController) GetParam() {
	param1 := c.GetString("param")
	c.Data["param1"] = param1
	c.TplName = "print.tpl"
}

func (c *MainController) PostFile() {
	c.TplName = "uploadFile.tpl"
}

func (c *MainController) SaveFile() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	c.SaveToFile("uploadname", "static/upload/"+h.Filename)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = struct {
		Message string
	}{"upload ok!!!"}
	c.ServeJSON()
}
