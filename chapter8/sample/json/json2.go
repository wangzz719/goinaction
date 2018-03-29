package main

import (
	"encoding/json"
	"log"
	"fmt"
)

func main() {
	c := map[string]interface{} {
		"name": "Gopher",
		"title": "programer",
		"contact": map[string]interface{} {
			"home": "415.333.3333",
			"ceil": "415.555.3333",
		},
	}
	// 编码为不带缩进的 json 文档
	data, err := json.Marshal(c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(string(data))

	// 编码为带缩紧的 json 文档
	data, err = json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(string(data))
}
