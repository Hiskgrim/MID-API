package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["contratacion_mid_api/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["contratacion_mid_api/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "Clasificar",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
