package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Country  string `json:"country"`
	City     string `json:"city"`
	Provider string `json:"isp"`
	IP       string `json:"query"`
}

func main() {
	resp, err := http.Get("http://ip-api.com/json/")

	if err != nil {
		fmt.Println("Возникла ошибка запроса: ", err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Возникла ошибка чтения ответа: ", err)
		return
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Ошибка парсинга", err)
		return
	}

	fmt.Println("Country: ", response.Country)
	fmt.Println("City: ", response.City)
	fmt.Println("Provider: ", response.Provider)
	fmt.Println("IP: ", response.IP)
}
