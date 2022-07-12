package models

import "time"

type OrdenPago struct {
	Activo                 bool      `json:"Activo"`
	FechaModificacion      time.Time `json:"Fecha_modificacion"`
	ID                     string    `json:"_id"`
	AreaFuncional          int       `json:"AreaFuncional"`
	Consecutivo            string    `json:"Consecutivo"`
	SolicitudGiro          string    `json:"SolicitudGiro"`
	Vigencia               int       `json:"Vigencia"`
	DocumentoBeneficiario  string    `json:"DocumentoBeneficiario"`
	NombreBeneficiario     string    `json:"NombreBeneficiario"`
	Compromiso             int       `json:"Compromiso"`
	NumeroCompromiso       int       `json:"NumeroCompromiso"`
	Supervisor             string    `json:"Supervisor"`
	Detalle                string    `json:"Detalle"`
	TipoOrdenPago          int       `json:"TipoOrdenPago"`
	ImputacionPresupuestal []struct {
		ID             string `json:"_id"`
		Disponibilidad string `json:"Disponibilidad"`
		Codigo         string `json:"Codigo"`
		Registro       int    `json:"Registro"`
		Nombre         string `json:"Nombre"`
		Valor          int    `json:"Valor"`
	} `json:"ImputacionPresupuestal"`
	Concepto             string `json:"Concepto"`
	ImpuestosRetenciones []struct {
		ID        string `json:"_id"`
		Nombre    string `json:"Nombre"`
		Codigo    string `json:"Codigo"`
		Base      int    `json:"Base"`
		Descuento string `json:"Descuento"`
		Valor     int    `json:"Valor"`
	} `json:"ImpuestosRetenciones"`
	CuentaValorNeto    string `json:"CuentaValorNeto"`
	MovimientoContable []struct {
		ID     string `json:"_id"`
		Nombre string `json:"Nombre"`
		Codigo string `json:"Codigo"`
		Valor  int    `json:"Valor"`
	} `json:"MovimientoContable"`
	FechaCreacion time.Time `json:"Fecha_creacion"`
	V             int       `json:"__v"`
}
