package models

type ContabilizacionOrdenPago struct {
	OrdenPagoId string `json:"ordenPagoId"`
	Aprobacion  bool   `json:"aprobacion"`
}
