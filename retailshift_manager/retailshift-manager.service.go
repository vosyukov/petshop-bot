package retailshift_manager

import (
	"awesomeProject1/moysklad_api"
	"awesomeProject1/telegram"
)

func OnHandleUpdateRetailshift(retailshiftId string) {

	telegram.SendMessage(263537201, retailshiftId)
}

func OnHandleCreateRetailshift(retailshiftId string) {

	var retailshift, _ = moysklad_api.GetRetailShift(retailshiftId)

	var text = "Открыта смена: " + retailshift.Name

	telegram.SendMessage(263537201, text)
}
