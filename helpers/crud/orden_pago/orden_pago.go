package ordenpago

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	helpers "central_cuentas_mid/helpers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_contables_mid/models"
	e "github.com/udistrital/utils_oas/errorctrl"
	r "github.com/udistrital/utils_oas/request"
)

func GetOrdenPagoId(id string, orden_pago interface{}) (outputError map[string]interface{}) {
	const funcion string = "GetOrdenPagoId"
	var fullResponse map[string]interface{}
	defer e.ErrorControlFunction(funcion+" - Unhandled Error!", strconv.Itoa(http.StatusInternalServerError))

	url := beego.AppConfig.String("CentralCuentasCrudService") + "/orden-pago/" + id
	if resp, err := r.GetJsonTest(url, &fullResponse); err != nil || resp.StatusCode != http.StatusOK {
		status := http.StatusBadGateway
		if err == nil { // resp.StatusCode != http.StatusOK
			err = fmt.Errorf("undesired status code - %s", http.StatusText(resp.StatusCode))
			status = resp.StatusCode
		}
		logs.Error(err)
		outputError = e.Error(funcion+" - r.GetJsonTest(url, &orden-pago)", err, strconv.Itoa(status))
	}

	var transaccion models.TransaccionMovimientos = models.TransaccionMovimientos{
		ConsecutivoId:    fullResponse["Consecutivo"].(int),
		Etiquetas:        "",
		Descripcion:      "",
		FechaTransaccion: time.Time{},
		Activo:           true,
		Movimientos:      []models.MovimientoResumido{},
	}

	var response map[string]interface{}

	if err := r.SendJson(beego.AppConfig.String("MovimientosContablesMidService")+"//transaccion_movimientos", "POST", &response, transaccion); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/RegistroTransaccionMovimientos9", "err": err.Error(), "status": "502"}
		return outputError
	}

	helpers.LimpiezaRespuestaRefactorBody(fullResponse, &orden_pago)
	return
}
