package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["testRouter/controllers:AnnotionController"] = append(beego.GlobalControllerRouter["testRouter/controllers:AnnotionController"],
		beego.ControllerComments{
			Method: "Mymethod",
			Router: `/task/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
