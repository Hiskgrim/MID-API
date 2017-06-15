package controllers

import(
	"fmt"
	//"contratacion_mid_api/models"
	"github.com/astaxie/beego"
	. "github.com/mndrix/golog"
	"strconv"
)

type ValidarContratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ValidarContratoController) URLMapping() {
	c.Mapping("ValidarContrato", c.ValidarContrato)
}

func (c *ValidarContratoController) ValidarContrato (){

	idProfesorStr := c.Ctx.Input.Param(":idProfesor")

	numHorasStr := c.Ctx.Input.Param(":numHoras")
	//numHoras, _ := strconv.Atoi(numHorasStr)
	vinculacion := c.Ctx.Input.Param(":dedicacion")

	predicados := `horas_semanales(`+idProfesorStr+`,`+vinculacion+`,`+numHorasStr+`).`+ "\n"

	reglasbase := CargarReglasBase()
	reglasbase = reglasbase+predicados
	fmt.Println(reglasbase)

	m := NewMachine().Consult(reglasbase)

	var a string
	contratos := m.ProveAll(`cumple_tiempo(`+idProfesorStr+`,X).`)
	for _, solution := range contratos {
	    a = fmt.Sprintf("%s", solution.ByName_("X"))
	}
	fmt.Println(a)
	fmt.Println(idProfesorStr)
	fmt.Println(numHorasStr)
	fmt.Println(vinculacion)
	validez, _ := strconv.Atoi(a)

	c.Data["json"] = validez

	c.ServeJSON()
}