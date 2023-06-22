package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// suppose this is some json that you fetched from an API
	jsonString := `{"Name":"Anurag","Address":"Bangalore","Age":23}`

	// convert the json string to a byte slice
	bs := []byte(jsonString)

	var jsonObject interface{}

	if err := json.Unmarshal(bs, &jsonObject); err != nil {
		panic(err)
	} else {
		// you've successfully populated the jsonObject
		// Let's see what we'll get when we log it on the console
		fmt.Println(jsonObject)
		// Also, let's log its type as well
		fmt.Printf("%T\n", jsonObject)

		if jsonMap, ok := jsonObject.(map[string]interface{}); !ok {
			panic("map[string]interface{} not found")
		} else {
			fmt.Printf("{\n")
			for k, v := range jsonMap {
				fmt.Printf("\t%v:%v,\n", k, v)
			}
			fmt.Printf("}\n")
		}
	}
}
