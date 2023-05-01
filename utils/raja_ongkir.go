package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetDistrict() {
	url := "https://api.rajaongkir.com/starter/province"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "a82f240de01102755058ed742668c223")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
