// @APIVersion 1.12.1
// @Title beego Central Cuentas API
// @Description beego has a very cool tools to autogenerate documents for your API
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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
