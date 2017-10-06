package routers

import (
	"testRouter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/?:id", &controllers.MainController{}, "get:QuestRegex")
	beego.Router("/api2/:id", &controllers.MainController{}, "get:QuestRegex")
	//无法匹配:id为空的情况
	beego.Router("/download/*.*", &controllers.MainController{}, "get:Download")
	beego.Router("/testsplat/*", &controllers.MainController{}, "get:Splat")
	beego.AutoRouter(&controllers.AutoController{})
	beego.Include(&controllers.AnnotionController{})
}
