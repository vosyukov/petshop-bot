package retailshift_manager

import "awesomeProject1/telegram"

func OnHandleUpdateRetailshift(retailshiftId string) {
	telegram.SendMessage(263537201, retailshiftId)
}

func OnHandleCreateRetailshift(retailshiftId string) {
	telegram.SendMessage(263537201, retailshiftId)
}
