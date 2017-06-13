package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getWeather(lat string, lon string) (w *Weather, e error) {
	key := os.Getenv("DARKSKY_KEY")
	url := "https://api.darksky.net/forecast/" + key + "/" + lat + "," + lon + "?units=auto&exclude=flags,minutely"

	log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// req.Header.Set("Accept-Encoding", "gzip")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var weather *Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	fmt.Println(weather.Currently)
	return weather, nil
}
