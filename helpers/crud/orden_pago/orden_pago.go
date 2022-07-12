package ordenpago

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	helpers "central_cuentas_mid/helpers"
	m "central_cuentas_mid/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_contables_mid/models"
	e "github.com/udistrital/utils_oas/errorctrl"
	r "github.com/udistrital/utils_oas/request"
)

func PostMovimientosOrdenPagoId(id string) (outputError map[string]interface{}) {
	const funcion string = "GetOrdenPagoId"
	var fullResponse map[string]interface{}
	var orden m.OrdenPago
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

	helpers.LimpiezaRespuestaRefactorBody(fullResponse, &orden)
	consecutivoId, _ := strconv.Atoi(orden.Consecutivo)

	var movimientos []models.MovimientoResumido

	for _, v := range orden.MovimientoContable {
		var mov models.MovimientoResumido = models.MovimientoResumido{
			CuentaId:         v.ID,
			NombreCuenta:     v.Nombre,
			TipoMovimientoId: 45,
			Valor:            float64(v.Valor),
			Descripcion:      orden.Detalle,
			Activo:           false,
		}
		movimientos = append(movimientos, mov)
	}

	var transaccion models.TransaccionMovimientos = models.TransaccionMovimientos{
		ConsecutivoId:    consecutivoId,
		Etiquetas:        "",
		Descripcion:      orden.Detalle,
		FechaTransaccion: time.Time{},
		Activo:           true,
		Movimientos:      movimientos,
	}

	var response map[string]interface{}

	if err := r.SendJson(beego.AppConfig.String("MovimientosContablesMidService")+"/transaccion_movimientos", "POST", &response, transaccion); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/RegistroTransaccionMovimientos9", "err": err.Error(), "status": "502"}
		return outputError
	}

	return
}
