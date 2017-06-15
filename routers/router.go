// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"contratacion_mid_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/calculo_salario",
			beego.NSInclude(
				&controllers.CalculoSalarioController{},
			),
		),

		beego.NSNamespace("/validar_contrato",
			beego.NSInclude(
				&controllers.ValidarContratoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
