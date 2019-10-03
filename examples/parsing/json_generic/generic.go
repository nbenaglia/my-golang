package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	bytes := []byte(`{
		"id" : 1,
		"content" : "Hello World!",
		"author" : {
			"id" : 2,
			"name" : "Nicola"
		},
		"comments" : [
			{ 
				"id" : 3, 
				"content" : "Have a great day!", 
				"author" : "Nico"
			},
			{
				"id" : 4, 
				"content" : "How are you today?", 
				"author" : "Stella"
			}
		]
	}`)
	var p2 interface{}
	json.Unmarshal(bytes, &p2)
	fmt.Println(p2)
}
