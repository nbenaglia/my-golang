package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	gResult struct {
		Ip         string  `json:"ip"`
		IpDecimal  int     `json:"ip_decimal"`
		Country    string  `json:"country"`
		CountryEu  bool    `json:"country_eu"`
		CountryIso string  `json:"country_iso"`
		City       string  `json:"title"`
		Hostname   string  `json:"hostname"`
		Latitude   float32 `json:"latitude"`
		Longitude  float32 `json:"longitude"`
		Asn        string  `json:"asn"`
		AsnOrg     string  `json:"asn_org"`
	}
)

func main() {
	uri := "https://ifconfig.co/json"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into our struct type.
	var result gResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(result)

	// Marshal the struct type into a pretty print version of the JSON document.
	pretty, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(pretty))
}
