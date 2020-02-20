package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	//json
	p1 := Person{"Jack", 12}

	// stringify
	jsonStr, _ := json.Marshal(p1)

	fmt.Printf("%s", string(jsonStr))
}