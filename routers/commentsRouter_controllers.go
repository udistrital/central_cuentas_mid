package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/central_cuentas_mid/controllers:ContabilizarOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/central_cuentas_mid/controllers:ContabilizarOrdenPagoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
