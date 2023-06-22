package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  byte
	pin  int
}

func main() {
	p1 := Person{
		Name: "Anurag",
		Age:  23,
		pin:  1234,
	}
	fmt.Printf("Person Object: %+v\n", p1) //{Name:Anurag Age:23 pin:1234}
	if bs, err := json.Marshal(p1); err != nil {
		panic(err)
	} else {
		// byte slice
		fmt.Printf("JSON Byte Slice: %v\n", bs)

		// json string
		fmt.Printf("JSON String: %v\n", string(bs)) //pin will be not be parsed because its in lowercase

		// unmarshaling it
		var p2 Person
		if err := json.Unmarshal(bs, &p2); err != nil {
			panic(err)
		}
		fmt.Printf("Unmarshaled Person Object: %+v\n", p2) // {Name:Anurag Age:23 pin:0}

	}

}
