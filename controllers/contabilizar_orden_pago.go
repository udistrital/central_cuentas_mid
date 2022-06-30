package controllers

import (
	"github.com/astaxie/beego"
)

// ContabilizarOrdenPagoController operations for Contabilizar_orden_pago
type ContabilizarOrdenPagoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContabilizarOrdenPagoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Contabilizar_orden_pago
// @Param	body		body 	models.Contabilizar_orden_pago	true		"body for Contabilizar_orden_pago content"
// @Success 201 {object} models.Contabilizar_orden_pago
// @Failure 403 body is empty
// @router / [post]
func (c *ContabilizarOrdenPagoController) Post() {

}
