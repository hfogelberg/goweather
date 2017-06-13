package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getWeather(url string) (w *Weather, e error) {
	log.Println("get weather")
	log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

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
