package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(data interface{}) {
	json, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println("JSON:", string(json))
}
