package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Cost struct {
	Value int    `json:"value"`
	ETD   string `json:"etd"`
	Note  string `json:"note"`
}

type CostDetail struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []Cost `json:"cost"`
}

type Courier struct {
	Code  string       `json:"code"`
	Name  string       `json:"name"`
	Costs []CostDetail `json:"costs"`
}

type RajaOngkir struct {
	Results []Courier `json:"results"`
}

type Result struct {
	RajaOngkir RajaOngkir `json:"rajaongkir"`
}

func ResponseData(destination, weight string) (string, int) {
	url := "https://api.rajaongkir.com/starter/cost"
	origin := "160"
	payload := strings.NewReader("origin=" + origin + "&destination=" + destination + "&weight=" + weight + "&courier=jne")

	req, _ := http.NewRequest("POST", url, payload)

	godotenv.Load()
	keyRajaOngkir := os.Getenv("KEY_RAJAONGKIR")
	req.Header.Add("key", keyRajaOngkir)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		defer res.Body.Close()

	}
	body, _ := ioutil.ReadAll(res.Body)

	var result Result
	check := json.Unmarshal(body, &result)
	if check != nil {
		return "", 0
	}

	cost := result.RajaOngkir.Results[0].Costs[0].Cost[0].Value
	ekspedisi := result.RajaOngkir.Results[0].Name
	return ekspedisi, cost
}
