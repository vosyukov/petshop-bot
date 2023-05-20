package moysklad_api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var client *http.Client

var API_KEY = os.Getenv("MOYSKLAD_API_KEY")

func init() {
	client = http.DefaultClient
}

func SendHttpGetRequest(url1 string, params string) ([]byte, error) {

	fmt.Println(url1 + "?" + params)

	req, err := http.NewRequest("GET", url1+"?"+params, nil)

	req.Header.Set("Authorization", "Bearer "+API_KEY)

	resp, err := client.Do(req)

	defer resp.Body.Close()
	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

	body, err := io.ReadAll(resp.Body)

	return body, err
}
