package moysklad_api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

var client *http.Client

const API_KEY = "300fe2a14e70b4d9db9fca7d4c2fa688378d293b"

func init() {
	client = http.DefaultClient
}

func SendHttpGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

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
