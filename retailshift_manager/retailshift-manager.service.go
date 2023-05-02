package retailshift_manager

import (
	"awesomeProject1/moysklad_api"
	"awesomeProject1/telegram"
)

func OnHandleUpdateRetailshift(retailshiftId string) {

	var retailshift, _ = moysklad_api.GetRetailShift(retailshiftId)

	if retailshift.CloseDate != "" {

		var text = "Закрыта смена: " + retailshift.Name + "\n"
		text += "Время закрытия: " + retailshift.Created + "\n"

		telegram.SendMessage(263537201, text)
	}

}

func OnHandleCreateRetailshift(retailshiftId string) {

	var retailshift, _ = moysklad_api.GetRetailShift(retailshiftId)

	var text = "Открыта смена: " + retailshift.Name + "\n"
	text += "Время открытия: " + retailshift.Created + "\n"

	telegram.SendMessage(263537201, text)
}
