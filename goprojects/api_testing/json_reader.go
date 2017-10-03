package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"
	"io/ioutil"
)

type people struct {
	People []string `json: "people"`
	Number int `json: "number"`
	Message string `json: "message"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"
	client := http.Client{
			Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "json_client-testing")

	res, getErr := client
  .Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Printf("HTTP: %s\n", res.Status)

	/*
	text := `{

			"people": [
				{"craft": "ISS", "name": "Sergey Rizhikov"},
				{"craft": "ISS", "name": "Andrey Borisenko"},
				{"craft": "ISS", "name": "Shane Kimbrough"},
				{"craft": "ISS", "name": "Oleg Novitskiy"},
				{"craft": "ISS", "name": "Thomas Pesquet"},
				{"craft": "ISS", "name": "Peggy Whitson"}],
			"message": "success",
			"number": 6}`
	textBytes := []byte(text)

	*/

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if err != nil {
		log.Fatal(jsonErr)
		return
	}

	fmt.Println(people1.People[3])
	fmt.Println(people1.Number)
	fmt.Println(people1.Message)
}


