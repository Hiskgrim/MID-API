package controllers

import (
	"fmt"
	"contratacion_mid_api/models"
	"github.com/astaxie/beego"
	. "github.com/mndrix/golog"
	"strconv"
)

// PreliquidacionController operations for Preliquidacion
type ClasificacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ClasificacionController) URLMapping() {
	c.Mapping("Clasificar", c.Clasificar)
}

// Post ...
// @Title Create
// @Description create Preliquidacion
// @Param	body		body 	models.Preliquidacion	true		"body for Preliquidacion content"
// @Success 201 {object} models.Preliquidacion
// @Failure 403 body is empty
// @router / [post]
func (c *ClasificacionController) Clasificar() {

		idPersonaStr := c.Ctx.Input.Param(":idPersona")
		numSemanasStr := c.Ctx.Input.Param(":numSemanas")
		fmt.Println(numSemanasStr)
		numSemanas, err := strconv.Atoi(numSemanasStr)
		if err != nil {
   		}
		numHorasSemanalesStr := c.Ctx.Input.Param(":numHorasSemanales")
		fmt.Println(numHorasSemanalesStr)
		numHorasSemanales,err := strconv.Atoi(numHorasSemanalesStr)
		if err != nil {
   		}
		categoriaStr := c.Ctx.Input.Param(":categoria")
		dedicacionStr := c.Ctx.Input.Param(":dedicacion")


		reglasbase := `categoria(`+idPersonaStr+`,`+categoriaStr+`,2016).`+"\n"
		reglasbase = reglasbase+`vinculacion(`+idPersonaStr+`,`+dedicacionStr+`,2016).`+"\n"
		reglasbase = reglasbase+`horas(`+idPersonaStr+`,`+strconv.Itoa(numSemanas*numHorasSemanales)+`,2016).`+"\n"
		reglasbase = reglasbase+`valor(`+ strconv.FormatFloat(CargarValorPunto(), 'f', 6, 64)+`,2016).`+"\n"

		reglasbase = reglasbase + CargarReglasBase()
		fmt.Println(reglasbase)

		m := NewMachine().Consult(reglasbase)



		/*experiencia := CargarExperienciaLaboral()
		fmt.Println(experiencia)

		titulosPregrado, titulosPosgrado:= CargarFormacionAcademica()
		fmt.Println(titulosPregrado)
		fmt.Println(titulosPosgrado)

		investigaciones := CargarTrabajosInvestigacion()
		fmt.Println(investigaciones)*/		

		var a string

		contratos := m.ProveAll(`valor_contrato(`+idPersonaStr+`,2016,X).`)
		for _, solution := range contratos {
		  a = fmt.Sprintf("%s", solution.ByName_("X"))
		}

		fmt.Printf(a);

		c.Data["json"] = a

		c.ServeJSON()

}

func CargarReglasBase() (reglas string) {
	//carga de reglas desde el ruler
	var reglasbase string = ``
	var v []models.Predicado

	if err := getJson("http://"+beego.AppConfig.String("Urlruler")+predicado/?limit=0", &v); err == nil {
		for _, regla := range v {
			reglasbase = reglasbase + regla.Nombre + "\n"
		}
	} else {

	}
	return reglasbase
}

func CargarValorPunto() (valor float64) {
	var valorPunto float64 = 0
	var v []models.PuntoSalarial
	if err := getJson("http://"+beego.AppConfig.String("cdveService")+punto_salarial/?limit=0", &v); err == nil {
		valorPunto=v[0].ValorPunto
	} else {

	}
	return valorPunto
}

func CargarExperienciaLaboral() (experienciaLaboral int) {
	//carga de reglas desde el ruler
	var experiencias int = 0
	var v []models.ExperienciaDocente

	if err := getJson("http://"+beego.AppConfig.String("hojasdevidaService")+experiencia_docente/?limit=0", &v); err == nil {
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

	if err := getJson("http://"+beego.AppConfig.String("hojasdevidaService")+formacion_academica/?limit=0", &v); err == nil {
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

	if err := getJson("http://"+beego.AppConfig.String("hojasdevidaService")+investigacion/?limit=0", &v); err == nil {
		trabajos=len(v)
	} else {

	}
	return trabajos
}
