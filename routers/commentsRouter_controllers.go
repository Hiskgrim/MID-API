package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["contratacion_mid_api/controllers:CalculoSalarioController"] = append(beego.GlobalControllerRouter["contratacion_mid_api/controllers:CalculoSalarioController"],
		beego.ControllerComments{
			Method: "CalcularSalario",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
