package moysklad_api

import (
	"net/url"
	"strconv"
)
import (
	"encoding/json"
)

type Filter struct {
	movementFrom string
	movementTo   string
}

func init() {

	GetReportByEmployee(Filter{movementTo: "2023-05-03", movementFrom: "2016-04-15 15:48:46"}, 0, 20)
}

func GetReportByEmployee(filter Filter, offset int, limit int) (Employee, error) {
	var url1 = "offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)

	url1 += "&filter=" + url.QueryEscape("momentFrom="+filter.movementFrom)

	var data, err = SendHttpGetRequest("https://online.moysklad.ru/api/remap/1.2/report/profit/byemployee", url1)

	var red Employee

	json.Unmarshal([]byte(data), &red)

	return red, err
}
