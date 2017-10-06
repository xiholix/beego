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

func (c *MainController) QuestRegex() {
	s := c.Ctx.Input.Param(":id")
	ext := c.Ctx.Input.Param(":ext")
	c.Data["name"] = s
	c.Data["ext"] = ext
	c.TplName = "question.tpl"
}

func (c *MainController) Download() {
	path := c.Ctx.Input.Param(":path")
	ext := c.Ctx.Input.Param(":ext")
	c.Data["path"] = path
	c.Data["ext"] = ext
	c.TplName = "download.tpl"
}

func (c *MainController) Splat() {
	splat := c.Ctx.Input.Param(":splat")
	c.Data["splat"] = splat
	c.TplName = "splat.tpl"
}
