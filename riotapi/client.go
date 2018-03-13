package riotapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// TODO: Make the URL configurable
func getObj(url string, obj interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Riot-Token", RIOT_API_KEY)

	res, getErr := Client.Do(req)
	if getErr != nil {
		return err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}
