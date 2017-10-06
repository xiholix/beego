package controllers

import (
	"github.com/astaxie/beego"
)

type AnnotionController struct {
	beego.Controller
}

func (this *AnnotionController) URLMapping() {
	this.Mapping("Mymethod", this.Mymethod)
}

//@router /task/:id
func (this *AnnotionController) Mymethod() {
	id := this.Ctx.Input.Param(":id")
	this.Data["param1"] = id
	this.TplName = "autocontroller.tpl"
}
