package controllers

import (
	"fmt"
	"contratacion_mid_api/models"
	"github.com/astaxie/beego"
	. "github.com/mndrix/golog"
	"strconv"
)

// PreliquidacionController operations for Preliquidacion
type CalculoSalarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *CalculoSalarioController) URLMapping() {
	c.Mapping("CalcularSalario", c.CalcularSalario)
}

// Post ...
// @Title Create
// @Description create Preliquidacion
// @Param	body		body 	models.Preliquidacion	true		"body for Preliquidacion content"
// @Success 201 {object} models.Preliquidacion
// @Failure 403 body is empty
// @router / [post]
func (c *CalculoSalarioController) CalcularSalario() {

		nivelAcademico := c.Ctx.Input.Param(":nivelAcademico")

		idProfesorStr := c.Ctx.Input.Param(":idProfesor")

		numHorasStr := c.Ctx.Input.Param(":numHoras")
		numHoras, _ := strconv.Atoi(numHorasStr)

		numSemanasStr := c.Ctx.Input.Param(":numSemanas")
		numSemanas, _ := strconv.Atoi(numSemanasStr)

		categoria := c.Ctx.Input.Param(":categoria")

		vinculacion := c.Ctx.Input.Param(":dedicacion")

		predicados := `categoria(`+idProfesorStr+`,`+categoria+`, 2016).`+ "\n"
		predicados = predicados+`vinculacion(`+idProfesorStr+`,`+vinculacion+`,2016).`+ "\n"
		predicados = predicados+`horas(`+idProfesorStr+`,`+strconv.Itoa(numHoras*numSemanas)+`,2016).`+ "\n"

		reglasbase := CargarReglasBase()
		reglasbase = reglasbase+predicados
		fmt.Println(reglasbase)

		m := NewMachine().Consult(reglasbase)

		/*experiencia := CargarExperienciaLaboral()
		fmt.Println(experiencia)

		titulosPregrado, titulosPosgrado:= CargarFormacionAcademica()
		fmt.Println(titulosPregrado)
		fmt.Println(titulosPosgrado)

		investigaciones := CargarTrabajosInvestigacion()
		fmt.Println(investigaciones)		

		var a string
		validar := m.ProveAll(`categoria(`+strconv.Itoa(experiencia)+`,`+strconv.Itoa(investigaciones)+`,`+strconv.Itoa(titulosPregrado)+`,`+strconv.Itoa(titulosPosgrado)+`,X).`)
		for _, solution := range validar {
		  a = fmt.Sprintf("%s", solution.ByName_("X"))
		}

		fmt.Printf(a);*/

		//clasificacion := CargarClasificacion()

		var a string

			contratos := m.ProveAll(`valor_contrato(`+nivelAcademico+`,`+idProfesorStr+`,2016,X).`)
			for _, solution := range contratos {
			    a = fmt.Sprintf("%s", solution.ByName_("X"))
			}


		f, _ := strconv.ParseFloat(a, 64)
	    salario := int(f)

		c.Data["json"] = salario

		c.ServeJSON()

}

func CargarClasificacion() (reglas string) {
	//carga de reglas desde el ruler
	var categoria string = ""
	var v []models.NivelEscalafon

	if err := getJson("http://localhost:8081/v1/nivel_escalafon/?limit=0", &v); err == nil {
		for _, escalafon := range v {
			categoria = escalafon.NombreEscalafon
		}
	} else {

	}
	return categoria
}

func CargarReglasBase() (reglas string) {
	//carga de reglas desde el ruler
	var reglasbase string = ``
	var v []models.Predicado

	if err := getJson("http://localhost:8086/v1/predicado/?limit=0", &v); err == nil {
		for _, regla := range v {
			reglasbase = reglasbase + regla.Nombre + "\n"
		}
	} else {

	}
	return reglasbase
}

func CargarExperienciaLaboral() (experienciaLaboral int) {
	//carga de reglas desde el ruler
	var experiencias int = 0
	var v []models.ExperienciaDocente

	if err := getJson("http://localhost:8081/v1/experiencia_docente/?limit=0", &v); err == nil {
		experiencias=len(v)
	} else {

	}
	return experiencias
}

func CargarFormacionAcademica() (titulospregrado int, titulosposgrado int) {
	//carga de reglas desde el ruler
	var titulosPregrado int = 0
	var titulosPosgrado int = 0
	var v []models.FormacionAcademica

	if err := getJson("http://localhost:8081/v1/formacion_academica/?limit=0", &v); err == nil {
		titulosPregrado=len(v)
		titulosPosgrado=len(v)
	} else {

	}
	return titulosPregrado, titulosPosgrado
}

func CargarTrabajosInvestigacion() (trabajosInvestigacion int) {
	//carga de reglas desde el ruler
	var trabajos int = 0
	var v []models.Investigacion

	if err := getJson("http://localhost:8081/v1/investigacion/?limit=0", &v); err == nil {
		trabajos=len(v)
	} else {

	}
	return trabajos
}