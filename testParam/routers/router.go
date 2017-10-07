package routers

import (
	"testParam/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/get", &controllers.MainController{}, "get:GetParam")
	// beego.Router("/post2", &controllers.MainController{}, "get:PostFile;post:SaveFile") //;中间不能由空格
	beego.Router("/post2", &controllers.MainController{}, "get:PostFile")
	beego.Router("/anothersave", &controllers.MainController{}, "post:SaveFile")
}
