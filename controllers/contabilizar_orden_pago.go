package controllers

import (
	"central_cuentas_mid/models"
	"encoding/json"
	"net/http"

	orden_pago "central_cuentas_mid/helpers/crud/orden_pago"

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

	var v models.ContabilizacionOrdenPago = models.ContabilizacionOrdenPago{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		if err := orden_pago.PostMovimientosOrdenPagoId(v.OrdenPagoId); err == nil {
			//c.Data["json"] = "OK"

			c.Ctx.Output.SetStatus(http.StatusCreated)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": "OK"}
		} else {
			panic(err)
		}
	} else {
		panic(map[string]interface{}{"funcion": "PostTransaccionMovimientos", "err": err.Error(), "status": "400"})
	}

	c.ServeJSON()

}
