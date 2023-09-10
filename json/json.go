package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name  string  `json:"name"` // Name not name it send param to Marshal
	Level int     `json:"level"`
	Price float64 `json:"price"`
}

func main() {
	// Marshal
	a := Course{
		Name:  "Pom",
		Level: 5,
		Price: 10,
	}
	aa, _ := json.Marshal(a)
	fmt.Printf("aa Full: %#v\n", aa)
	fmt.Printf("aa Type a: %T\n", aa)
	fmt.Printf("aa byte: %v\n", aa)
	fmt.Printf("aa string: %s\n", aa)

	//###############################################
	fmt.Println("###############################################")

	//unMarchal
	data := []byte(`{
		"name": "Me",
		"level": 500,
		"price": 1000
	}`)

	b := Course{}
	err := json.Unmarshal(data, &b)
	if err != nil {
		fmt.Printf("Error unmarshalling: %v\n", err)
		return
	}
	fmt.Printf("bb Full: %#v\n", b)
	fmt.Printf("bb Type a: %T\n", b)
	fmt.Printf("bb byte: %v\n", b)

}
