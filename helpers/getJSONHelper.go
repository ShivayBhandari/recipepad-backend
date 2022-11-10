package helpers

import (
	"encoding/json"
	"net/http"
)

var client *http.Client

func GetJSON(url string, target interface{}) error {
	resp, err := client.Get(url)
	if(err != nil) {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}