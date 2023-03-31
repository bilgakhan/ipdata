package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPInfo struct {
	Country string `json:"country"`
	City    string `json:"regionName"`
	TZ      string `json:"timeZone"`
	ISP     string `json:"isp"`
}

func main() {
	fmt.Print("Enter IP address: ")
	var ip string
	fmt.Scanln(&ip)

	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var info IPInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		panic(err)
	}
	fmt.Printf("IP %s", ip)
	fmt.Printf("Country: %s\n", info.Country)
	fmt.Printf("City: %s\n", info.City)
	fmt.Printf("ISP: %s\n", info.ISP)
	fmt.Printf("Time Zone: %s\n", info.TZ)
}
