package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lynnau/fuelfor.cheap-api/config"
)

func fetch() (*ProjectZeroThreeAPIResponse, error) {
	uri := config.Get("priceapi.uri").(string)
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.0 Safari/537.36 Edg/80.0.360.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pzt := &ProjectZeroThreeAPIResponse{}
	err = json.Unmarshal(bytes, pzt)
	if err != nil {
		fmt.Println(string(bytes))
		return nil, err
	}

	return pzt, nil
}
