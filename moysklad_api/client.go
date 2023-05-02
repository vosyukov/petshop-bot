package moysklad_api

import (
	"io"
	"net/http"
)

var client *http.Client

const API_KEY = "a6f697ec1a1e0fb59e68dc617859287ceab260d8"

func init() {
	client = http.DefaultClient
}

func SendHttpGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", "Bearer "+API_KEY)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return body, err
}
