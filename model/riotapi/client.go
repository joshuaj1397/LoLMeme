package riotapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	RIOT_API_KEY string

	// Client would instantiate a new http.Client with a 5 sec timeout on requests
	client = http.Client{
		Timeout: time.Second * 5,
	}
)

// Get the RIOT_API_KEY from env vars
func init() {
	RIOT_API_KEY = os.Getenv("RIOT_API_KEY")
	if RIOT_API_KEY == "" {
		panic("RIOT_API_KEY NOT FOUND IN ENV")
	}
}

// TODO: Make the URL configurable
func GetObj(url string, obj interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Riot-Token", os.Getenv("RIOT_API_KEY"))
	fmt.Println(req.Header)

	res, getErr := client.Do(req)
	if getErr != nil {
		return err
	}
	fmt.Println(res)

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}
