package controllers

import (
	"net/http"

	"github.com/astaxie/beego"

	e "github.com/udistrital/utils_oas/errorctrl"
)

// ContabilizarOrdenPagoController operations for contabilizar_orden_pago
type ContabilizarOrdenPagoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContabilizarOrdenPagoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create contabilizar_orden_pago
// @Param	body		body 	{}	true		"body for contabilizar_orden_pago content"
// @Success 201 {object} {}
// @Failure 403 body is empty
// @router / [post]
func (c *ContabilizarOrdenPagoController) Post() {
	defer e.ErrorControlController(c.Controller, "ContabilizarOrdenPagoController")

	idStr := "e"
	if idStr != "" {
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": "OK"}
	} else {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": "500", "Message": "Fail", "Data": "fail"}
	}
	c.ServeJSON()
}
