package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["contratacion_mid_api/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["contratacion_mid_api/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "Clasificar",
			Router: `/:idPersona/:numSemanas/:numHorasSemanales/:categoria/:dedicacion`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
