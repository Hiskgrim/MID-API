package models

import (
	"time"
)

type PuntoSalarial struct {
	Id            int       `orm:"column(id);pk"`
	Vigencia      int       `orm:"column(vigencia)"`
	FechaRegistro time.Time `orm:"column(fecha_registro);type(date)"`
	ValorPunto    float64   `orm:"column(valor_punto)"`
	Comentarios   string    `orm:"column(comentarios);null"`
}
