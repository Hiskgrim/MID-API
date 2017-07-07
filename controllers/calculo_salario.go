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
	c.Mapping("CalcularSalarioPrecontratacion", c.CalcularSalarioPrecontratacion)
}

func (c *CalculoSalarioController) CalculoSalarioContratacion() {
	idVinculacionStr := c.Ctx.Input.Param(":idVincuacion")
	vinculacionDocente := CargarVinculacionDocente(idVinculacionStr)
	escalafon := CargarEscalafon(strconv.Itoa(vinculacionDocente.IdPersona))
	predicados := `categoria(`+strconv.Itoa(vinculacionDocente.IdPersona)+`,`+escalafon+`, 2016).`+ "\n"
	predicados = predicados+`vinculacion(`+strconv.Itoa(vinculacionDocente.IdPersona)+`,`+vinculacionDocente.IdDedicacion.NombreDedicacion+`,2016).`+ "\n"
	predicados = predicados+`horas(`+strconv.Itoa(vinculacionDocente.IdPersona)+`,`+strconv.Itoa(vinculacionDocente.NumeroHorasSemanales*vinculacionDocente.NumeroSemanas)+`,2016).`+ "\n"
	reglasbase := CargarReglasBase()
	reglasbase = reglasbase+predicados
	fmt.Println(reglasbase)
	m := NewMachine().Consult(reglasbase)
	var a string
	contratos := m.ProveAll(`valor_contrato(`+vinculacionDocente.IdResolucion.NivelAcademico+`,`+strconv.Itoa(vinculacionDocente.IdPersona)+`,2016,X).`)
	for _, solution := range contratos {
		a = fmt.Sprintf("%s", solution.ByName_("X"))
	}
	f, _ := strconv.ParseFloat(a, 64)
	salario := int(f)
	c.Data["json"] = salario
	c.ServeJSON()

}

func (c *CalculoSalarioController) CalcularSalarioPrecontratacion() {
	nivelAcademico := c.Ctx.Input.Param(":nivelAcademico")
	idPersonaStr := c.Ctx.Input.Param(":idProfesor")
	numHorasStr := c.Ctx.Input.Param(":numHoras")
	numHoras, _ := strconv.Atoi(numHorasStr)
	numSemanasStr := c.Ctx.Input.Param(":numSemanas")
	numSemanas, _ := strconv.Atoi(numSemanasStr)
	categoria := c.Ctx.Input.Param(":categoria")
	vinculacion := c.Ctx.Input.Param(":dedicacion")
	predicados := `categoria(`+idPersonaStr+`,`+categoria+`, 2016).`+ "\n"
	predicados = predicados+`vinculacion(`+idPersonaStr+`,`+vinculacion+`,2016).`+ "\n"
	predicados = predicados+`horas(`+idPersonaStr+`,`+strconv.Itoa(numHoras*numSemanas)+`,2016).`+ "\n"
	reglasbase := CargarReglasBase()
	reglasbase = reglasbase+predicados
	fmt.Println(reglasbase)
	m := NewMachine().Consult(reglasbase)
	var a string
	contratos := m.ProveAll(`valor_contrato(`+nivelAcademico+`,`+idPersonaStr+`,2016,X).`)
	for _, solution := range contratos {
		a = fmt.Sprintf("%s", solution.ByName_("X"))
	}
	f, _ := strconv.ParseFloat(a, 64)
	salario := int(f)
	c.Data["json"] = salario
	c.ServeJSON()

}

func CargarEscalafon(idPersona string) (e string) {
	escalafon := ""
	var v models.EscalafonPersona

	if err := getJson("http://10.20.0.254/cdve_api_crud/v1/escalafon_persona/?query=id_persona_natural:"+idPersona, &v); err == nil {
		escalafon=v.IdEscalafon.NombreEscalafon
	}
	return escalafon
}

func CargarVinculacionDocente(idVinculacion string) (a models.VinculacionDocente) {
	var vinculacionDocente models.VinculacionDocente

	if err := getJson("http://10.20.0.254/cdve_api_crud/v1/vinculacion_docente/"+idVinculacion, &vinculacionDocente); err == nil {
	} else {
	}
	return vinculacionDocente
}

func CargarReglasBase() (reglas string) {
	var reglasbase string = ``
	var v []models.Predicado

	if err := getJson("http://10.20.0.254/ruler/v1/predicado/?query=Dominio.Id:8&limit=0", &v); err == nil {
		for _, regla := range v {
			reglasbase = reglasbase + regla.Nombre + "\n"
		}
	} else {

	}
	return reglasbase
}
