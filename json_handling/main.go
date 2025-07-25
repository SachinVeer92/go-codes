package main

import (
	"encoding/json"
	"fmt"
)

// type User struct {
// 	Name    string   `json:"name"`
// 	Age     int      `json:"age"`
// 	Hobbies []string `json:"hobbies"`
// }

func main() {
	data := []byte(`{"data" : {"name" : "Sachin", "age" : 32, "hobbies" : ["football", "cricket"]} }`)

	var u map[string]map[string]interface{}

	err := json.Unmarshal(data, &u)

	if err != nil {
		fmt.Println(err)
		return
	}

	d := u["data"]

	fmt.Println(d["name"])
	fmt.Println(d["age"])
	fmt.Println(d["hobbies"])
}
