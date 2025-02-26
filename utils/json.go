package utils

import (
	"encoding/json"
	"fmt"
)

func PrintAsJSON(v any) {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j))
}
