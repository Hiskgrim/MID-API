package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["contratacion_mid_api/controllers:CalculoSalarioController"] = append(beego.GlobalControllerRouter["contratacion_mid_api/controllers:CalculoSalarioController"],
		beego.ControllerComments{
			Method: "CalcularSalario",
			Router: `/:nivelAcademico/:idProfesor/:numHoras/:numSemanas/:categoria/:dedicacion`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["contratacion_mid_api/controllers:ValidarContratoController"] = append(beego.GlobalControllerRouter["contratacion_mid_api/controllers:ValidarContratoController"],
		beego.ControllerComments{
			Method: "ValidarContrato",
			Router: `/:idProfesor/:numHoras/:dedicacion`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
