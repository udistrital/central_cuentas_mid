package routers

import (
	"central_cuentas_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/contabilizar_orden_pago",
			beego.NSInclude(
				&controllers.ContabilizarOrdenPagoController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
