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

}

func GetReportByEmployee(filter Filter, offset int, limit int) (Employee, error) {
	var url1 = "offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)

	url1 += "&filter=" + url.QueryEscape("momentFrom="+filter.movementFrom)

	var data, err = SendHttpGetRequest("https://online.moysklad.ru/api/remap/1.2/report/profit/byemployee", url1)

	var red Employee

	json.Unmarshal([]byte(data), &red)

	return red, err
}
