package controllers

import (
	"github.com/astaxie/beego"
)

type AutoController struct {
	beego.Controller
}

func (this *AutoController) Simple() {
	//http://localhost:8080/auto/simple/2018/02.json
	//param1 =2018 param2=02.json ext=""
	//http://localhost:8080/auto/simple.json
	//param1="" param2="" ext=json
	param1 := this.Ctx.Input.Param("0")
	param2 := this.Ctx.Input.Param("1")
	ext := this.Ctx.Input.Param(":ext")

	this.Data["param1"] = param1
	this.Data["param2"] = param2
	this.Data["ext"] = ext

	this.TplName = "autocontroller.tpl"
}
