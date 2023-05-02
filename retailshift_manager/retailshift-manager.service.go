package retailshift_manager

import (
	"awesomeProject1/moysklad_api"
	"awesomeProject1/telegram"
	"net/url"
	"path"
)

func OnHandleUpdateRetailshift(retailshiftId string) {

	var retailshift, _ = moysklad_api.GetRetailShift(retailshiftId)

	if retailshift.CloseDate != "" {

		u, _ := url.Parse(retailshift.Owner.Meta.Href)
		var employeeId = path.Base(u.Path)

		var employee, _ = moysklad_api.GetEmployee(employeeId)

		var text = "Закрыта смена: " + retailshift.Name + "\n"
		text += "Время закрытия: " + retailshift.Created + "\n"
		text += "Сотрудник: " + employee.FullName + "\n"

		telegram.SendMessage(263537201, text)
	}

}

func OnHandleCreateRetailshift(retailshiftId string) {

	var retailshift, _ = moysklad_api.GetRetailShift(retailshiftId)

	u, _ := url.Parse(retailshift.Owner.Meta.Href)
	var employeeId = path.Base(u.Path)

	var employee, _ = moysklad_api.GetEmployee(employeeId)

	var text = "Открыта смена: " + retailshift.Name + "\n"
	text += "Время открытия: " + retailshift.Created + "\n"
	text += "Сотрудник: " + employee.FullName + "\n"

	telegram.SendMessage(263537201, text)
}
